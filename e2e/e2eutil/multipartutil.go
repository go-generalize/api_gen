package e2eutil

import (
	"io"
	"mime/multipart"
)

// ReadMultipartFile reads payload in multipart file.
func ReadMultipartFile(fh *multipart.FileHeader) (string, error) {
	fp, err := fh.Open()

	if err != nil {
		return "", err
	}
	defer fp.Close()

	b, err := io.ReadAll(fp)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
