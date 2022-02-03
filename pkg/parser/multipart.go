package parser

import (
	"go/types"
	"reflect"

	go2tstypes "github.com/go-generalize/go2ts/pkg/types"
	"golang.org/x/xerrors"
)

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

// GetFileFields returns fields for files
func GetFileFields(obj *go2tstypes.Object) (res map[string]go2tstypes.ObjectEntry, err error) {
	res = make(map[string]go2tstypes.ObjectEntry)

	for k, v := range obj.Entries {
		t, err := GetMultipartUploadType(v.Type, v.RawTag)

		if err != nil {
			return nil, xerrors.Errorf("in %s: %w", v.RawName, err)
		}

		if t == UploadNone {
			continue
		}

		res[k] = v
	}

	return
}

// hasMultipartUpload checks t has *multipart.FileHeader or []*multipart.FileHeader
// Multipart data in map are not supported
func hasMultipartUpload(t *go2tstypes.Object) (bool, error) {
	found := false
	for _, v := range t.Entries {
		if obj, ok := v.Type.(*go2tstypes.Object); ok {
			has, err := hasMultipartUpload(obj)

			if err != nil {
				return false, xerrors.Errorf("in %s: %w", v.RawName, err)
			}

			if has {
				return true, nil
			}

			continue
		}

		res, err := GetMultipartUploadType(v.Type, v.RawTag)

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

// GetMultipartUploadType checks t is *multipart.FileHeader or []*multipart.FileHeader
// Multipart data in map are not supported
func GetMultipartUploadType(t go2tstypes.Type, tag string) (MultipartUploadType, error) {
	nullable, ok := t.(*go2tstypes.Nullable)

	if !ok {
		return UploadNone, nil
	}

	switch t := nullable.Inner.(type) {
	case *go2tstypes.Array:
		_, ok := t.Inner.(*go2tstypes.Nullable)

		if !ok {
			return UploadNone, nil
		}

		r, err := GetMultipartUploadType(t.Inner, tag)

		if err != nil {
			return UploadNone, err
		}

		if r == UploadSingleFile {
			return UploadMultipleFiles, nil
		}

		return UploadNone, nil
	case *MultipartHeader:
		if reflect.StructTag(tag).Get("form") == "" {
			return UploadNone, xerrors.Errorf("'form' tag for MultipartHeader is not set")
		}

		return UploadSingleFile, nil
	}

	return UploadNone, nil
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
