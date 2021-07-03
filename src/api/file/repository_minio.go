package file

import (
	"bytes"
	"context"
	"github.com/minio/minio-go"
	"net/url"
	"src/domain"
	"time"
)

type fileMinioRepository struct {
	client     *minio.Client
	bucketName string
}

func NewFileMinioRepository(client *minio.Client, bucketName string) (domain.FileRepository, error) {

	bucketExist, err := client.BucketExists(bucketName)
	if err != nil {
		return nil, err
	}

	if !bucketExist {
		err = client.MakeBucket(bucketName, "")
		if err != nil {
			return nil, err
		}
	}

	return &fileMinioRepository{client: client, bucketName: bucketName}, nil
}

func (repository *fileMinioRepository) Upload(ctx context.Context, files []*domain.UploadFileDTO) error {
	for _, file := range files {
		reader := bytes.NewReader(file.Payload)
		_, err := repository.client.PutObject(repository.bucketName, file.Path, reader, int64(len(file.Payload)), minio.PutObjectOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository *fileMinioRepository) UploadAttachments(ctx context.Context, attachments []*domain.TicketAttachment) error {
	files := make([]*domain.UploadFileDTO, len(attachments))
	for i, attachment := range attachments {
		files[i] = &domain.UploadFileDTO{
			Payload: attachment.Payload,
			Path:    attachment.Path,
		}
	}

	return repository.Upload(ctx, files)
}

func (repository *fileMinioRepository) URL(path string) (*url.URL, error) {
	return repository.client.PresignedGetObject(repository.bucketName, path, time.Hour*24, url.Values{})
}
