package fileService

import (
	"context"
	"mime/multipart"

	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/response"
)

type FileServiceInterface interface {
	Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader) (response.FileUploadRespondPayload, error)
	DeleteByID(fileid string) error
}
