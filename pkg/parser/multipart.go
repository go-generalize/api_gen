package parser

import (
	"go/types"
	"reflect"
	"sort"

	go2tstypes "github.com/go-generalize/go2ts/pkg/types"
	"golang.org/x/xerrors"
)

// FormTagKey is a tag key for fields uploaded in multipart
const FormTagKey = "form"

//MultipartHeader is mime/multipart.FileHeader
type MultipartHeader struct {
	go2tstypes.Common
}

var _ go2tstypes.Type = &MultipartHeader{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (uf *MultipartHeader) UsedAsMapKey() bool {
	return false
}

// String returns this type in string representation
func (uf *MultipartHeader) String() string {
	return "mime/multipart.FileHeader"
}

// FileField is a result type for GetFileFields
type FileField struct {
	Key     string
	Value   go2tstypes.ObjectEntry
	FormTag string
	Type    MultipartUploadType
}

// GetFileFields returns fields for files
func GetFileFields(obj *go2tstypes.Object) (res []FileField, err error) {
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
func hasMultipartUpload(t *go2tstypes.Object) (bool, error) {
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
	UploadNone MultipartUploadType = iota
	UploadSingleFile
	UploadMultipleFiles
)

// ValidateMultipartUploadType checks t is *multipart.FileHeader or []*multipart.FileHeader
// And validate form tag
func ValidateMultipartUploadType(t go2tstypes.Type, tag string) (MultipartUploadType, error) {
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
func GetMultipartUploadType(t go2tstypes.Type) MultipartUploadType {
	nullable, ok := t.(*go2tstypes.Nullable)

	if !ok {
		return UploadNone
	}

	switch t := nullable.Inner.(type) {
	case *go2tstypes.Array:
		_, ok := t.Inner.(*go2tstypes.Nullable)

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

func replacer(t types.Type) go2tstypes.Type {
	named, ok := t.(*types.Named)

	if !ok {
		return nil
	}

	if named.String() == "mime/multipart.FileHeader" {
		return &MultipartHeader{}
	}

	return nil
}
