package fileService

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/TimDebug/FitByte/src/infrastructure/domain"
	functionCallerInfo "github.com/TimDebug/FitByte/src/logger/helper"
	loggerZap "github.com/TimDebug/FitByte/src/logger/zap"
	"github.com/TimDebug/FitByte/src/model/dtos/response"
	"github.com/samber/do/v2"
)

type fileService struct {
	logger        loggerZap.LoggerInterface
	storageClient domain.StorageClient
}

func NewFileService(
	logger loggerZap.LoggerInterface,
	storageClient domain.StorageClient,
) FileServiceInterface {
	return &fileService{
		logger:        logger,
		storageClient: storageClient,
	}
}

func NewInject(i do.Injector) (FileServiceInterface, error) {
	_logger := do.MustInvoke[loggerZap.LoggerInterface](i)
	_storageClient := do.MustInvoke[domain.StorageClient](i)
	return NewFileService(_logger, _storageClient), nil
}

func (s *fileService) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader) (response.FileUploadRespondPayload, error) {
	// Simpan file ke server lokal
	file, err := header.Open()

	if err != nil {
		s.logger.Warn(err.Error(), functionCallerInfo.FileServiceUpload, header)
		return response.FileUploadRespondPayload{}, err
	}
	defer file.Close()

	// Tentukan lokasi penyimpanan file
	savePath := fmt.Sprintf("./.uploads/%s", header.Filename)

	// Buat direktori jika belum ada
	if err := os.MkdirAll("./.uploads", os.ModePerm); err != nil {
		s.logger.Error(err.Error(), functionCallerInfo.FileServiceUpload)
		return response.FileUploadRespondPayload{}, err
	}

	// Simpan file
	out, err := os.Create(savePath)
	if err != nil {
		s.logger.Error(err.Error(), functionCallerInfo.FileServiceUpload)
		s.logger.Error("Failed to save file", functionCallerInfo.FileServiceUpload)
		return response.FileUploadRespondPayload{}, err
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		s.logger.Error(err.Error(), functionCallerInfo.FileServiceUpload)
		s.logger.Error("Failed to write file", functionCallerInfo.FileServiceUpload)
		return response.FileUploadRespondPayload{}, err
	}

	return response.FileUploadRespondPayload{}, nil
}

func (s *fileService) DeleteByID(fileid string) error {
	return nil
}
