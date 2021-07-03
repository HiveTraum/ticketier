package domain

import (
	"context"
	"net/url"
)

type FileRepository interface {
	Upload(ctx context.Context, files []*UploadFileDTO) error
	UploadAttachments(ctx context.Context, attachments []*TicketAttachment) error
	URL(path string) (*url.URL, error)
}

type UploadFileDTO struct {
	Payload []byte
	Path    string
}
