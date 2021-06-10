package domain

import "net/url"

type FileRepository interface {
	Upload(path string, bytes []byte) (url.URL, error)
	Get(path string) (url.URL, error)
}
