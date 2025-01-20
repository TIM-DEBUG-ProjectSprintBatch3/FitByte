package fileController

import (
	"io"
	"path/filepath"

	"github.com/TimDebug/FitByte/src/infrastructure/domain"
	functionCallerInfo "github.com/TimDebug/FitByte/src/logger/helper"
	loggerZap "github.com/TimDebug/FitByte/src/logger/zap"
	fileService "github.com/TimDebug/FitByte/src/services/file"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type FileController interface {
	Upload(ctx *fiber.Ctx) error
}

type fileController struct {
	service       fileService.FileServiceInterface
	logger        loggerZap.LoggerInterface
	storageClient domain.StorageClient
}

func NewHandler(service fileService.FileServiceInterface, logger loggerZap.LoggerInterface, storageClient domain.StorageClient) FileController {
	return &fileController{service: service, logger: logger, storageClient: storageClient}
}

func NewInject(i do.Injector) (FileController, error) {
	_service := do.MustInvoke[fileService.FileServiceInterface](i)
	_logger := do.MustInvoke[loggerZap.LoggerInterface](i)
	_storageClient := do.MustInvoke[domain.StorageClient](i)
	return NewHandler(_service, _logger, _storageClient), nil
}

// Upload godoc
// @Tags file
// @Summary Upload an file
// @Description Upload an file
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token"
// @Param File formData file true "Body with file zip"
// @Success 200 {object} helper.Response{data=dto.FileUploadRespondPayload} "File uploaded successfully"
// @Success 201 {object} helper.Response{data=dto.FileUploadRespondPayload} "File created successfully"
// @Failure 400 {object} helper.Response{errors=helper.ErrorResponse} "Bad Request"
// @Failure 415 {object} helper.Response{errors=helper.ErrorResponse} "Unsupported Media Type"
// @Failure 413 {object} helper.Response{errors=helper.ErrorResponse} "Payload Too Large"
// @Failure 401 {object} helper.Response{errors=helper.ErrorResponse} "Unauthorized - Missing or invalid token"
// @Router /v1/file [POST]
func (h fileController) Upload(ctx *fiber.Ctx) error {
	// Validasi token
	if ctx.Get("Authorization") == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization token is required"})
	}

	// Ambil file dari request
	header, err := ctx.FormFile("file")
	if err != nil {
		h.logger.Warn(err.Error(), functionCallerInfo.FileControllerUpload, header)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to retrieve file"})
	}

	// Validasi tipe file
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		".jpeg":      true,
		".jpg":       true,
		".png":       true,
	}

	if !validTypes[header.Header.Get("Content-Type")] {
		h.logger.Warn("File ContentType failed:", functionCallerInfo.FileControllerUpload, header.Header.Get("Content-Type"))
		if !validTypes[filepath.Ext(header.Filename)] {
			h.logger.Warn("Invalid file type. Only jpeg, jpg, or png are allowed.", functionCallerInfo.FileControllerUpload, filepath.Ext(header.Filename))
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Only JPEG, JPG, or PNG files are allowed"})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Only JPEG, JPG, or PNG files are allowed"})
	}

	// Validasi ukuran file (max 100 KiB)
	if header.Size > 1024*100 {
		h.logger.Warn("File size exceeds 100KiB", functionCallerInfo.FileControllerUpload, header.Size)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File size exceeds 100KiB"})
	}

	// Upload file ke storage
	file, _ := header.Open()
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		h.logger.Warn(err.Error(), functionCallerInfo.FileControllerUpload, file)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read file"})
	}

	// Upload file ke storage
	uploadedURL, err := h.storageClient.PutFile(ctx.Context(), header.Filename, header.Header.Get("Content-Type"), fileContent, true)
	if err != nil {
		h.logger.Warn(err.Error(), functionCallerInfo.FileControllerUpload, header.Filename, header.Header.Get("Content-Type"), fileContent)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload file"})
	}

	// Respon sukses
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"uri": uploadedURL})
}
