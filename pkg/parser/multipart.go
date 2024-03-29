package parser

import (
	"go/types"
	"reflect"
	"sort"

	eptypes "github.com/go-generalize/go-easyparser/types"
	"golang.org/x/xerrors"
)

const fileHeaderType = `mime/multipart.FileHeader`

// FormTagKey is a tag key for fields uploaded in multipart
const FormTagKey = "form"

// MultipartHeader is mime/multipart.FileHeader
type MultipartHeader struct {
	eptypes.Common
}

var _ eptypes.Type = &MultipartHeader{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (uf *MultipartHeader) UsedAsMapKey() bool {
	return false
}

// String returns this type in string representation
func (uf *MultipartHeader) String() string {
	return fileHeaderType
}

// FileField is a result type for GetFileFields
type FileField struct {
	Key     string
	Value   eptypes.ObjectEntry
	FormTag string
	Type    MultipartUploadType
}

// GetFileFields returns fields for files
func GetFileFields(obj *eptypes.Object) (res []FileField, err error) {
	res = make([]FileField, 0)

	for k, v := range obj.Entries {
		t, err := ValidateMultipartUploadType(v.Type, v.RawTag)

		if err != nil {
			return nil, xerrors.Errorf("in %s: %w", v.RawName, err)
		}

		if t == UploadNone {
			continue
		}

		res = append(res, FileField{
			Key:     k,
			Value:   v,
			FormTag: reflect.StructTag(v.RawTag).Get(FormTagKey),
			Type:    t,
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Key < res[j].Key
	})

	return
}

// hasMultipartUpload checks t has *multipart.FileHeader or []*multipart.FileHeader
// Multipart data in map are not supported
func hasMultipartUpload(t *eptypes.Object) (bool, error) {
	found := false
	for _, v := range t.Entries {
		res, err := ValidateMultipartUploadType(v.Type, v.RawTag)

		if err != nil {
			return false, xerrors.Errorf("in %s: %w", v.RawName, err)
		}
		found = found || res != UploadNone
	}

	return found, nil
}

// MultipartUploadType is a type
type MultipartUploadType int

const (
	// UploadNone means no the field is not a file
	UploadNone MultipartUploadType = iota
	// UploadSingleFile means the field is a file
	UploadSingleFile
	// UploadMultipleFiles means the field is a slice of files
	UploadMultipleFiles
)

// ValidateMultipartUploadType checks t is *multipart.FileHeader or []*multipart.FileHeader
// And validate form tag
func ValidateMultipartUploadType(t eptypes.Type, tag string) (MultipartUploadType, error) {
	ut := GetMultipartUploadType(t)

	if ut == UploadNone {
		return UploadNone, nil
	}

	if reflect.StructTag(tag).Get(FormTagKey) == "" {
		return UploadNone, xerrors.Errorf("'%s' tag for MultipartHeader is not set", FormTagKey)
	}

	return ut, nil
}

// GetMultipartUploadType checks t is *multipart.FileHeader or []*multipart.FileHeader
// Multipart data in map are not supported
func GetMultipartUploadType(t eptypes.Type) MultipartUploadType {
	nullable, ok := t.(*eptypes.Nullable)

	if !ok {
		return UploadNone
	}

	switch t := nullable.Inner.(type) {
	case *eptypes.Array:
		_, ok := t.Inner.(*eptypes.Nullable)

		if !ok {
			return UploadNone
		}

		r := GetMultipartUploadType(t.Inner)

		if r == UploadSingleFile {
			return UploadMultipleFiles
		}

		return UploadNone
	case *MultipartHeader:
		return UploadSingleFile
	}

	return UploadNone
}

func replacer(t types.Type) eptypes.Type {
	named, ok := t.(*types.Named)

	if !ok {
		return nil
	}

	if named.String() == fileHeaderType {
		return &MultipartHeader{}
	}

	return nil
}
