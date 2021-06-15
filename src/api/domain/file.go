package domain

import "net/url"

type FileRepository interface {
	Upload(files []*UploadFileDTO) error
	URL(path string) (url.URL, error)
}

type UploadFileDTO struct {
	Payload []byte
	Path    string
}
