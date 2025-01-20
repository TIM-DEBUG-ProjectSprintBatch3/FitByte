package validator

import (
	"errors"
	"net/url"

	"github.com/TimDebug/FitByte/src/model/dtos/request"
	"github.com/go-playground/validator/v10"
)

var validate = func() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("uri_with_path", IsValidURI)

	return v
}()

func IsValidURI(fl validator.FieldLevel) bool {
	uriStr, ok := fl.Field().Interface().(string)
	if !ok {
		// If the field is not a string (e.g., nil pointer), skip validation.
		return true
	}

	if uriStr == "" {
		// Does not allow empty URIs
		return false
	}

	// Parse the URI using net/url
	parsedURI, err := url.Parse(uriStr)
	if err != nil || parsedURI.Scheme == "" || parsedURI.Host == "" {
		// Invalid URI, must have scheme (e.g., "http") and host (e.g., "example.com")
		return false
	}

	// Ensure there's a path (e.g., "/image.jpg")
	return parsedURI.Path != "" && parsedURI.Path != "/"
}

func ValidateAuthParams(input request.UserRegister) error {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			return fieldError
		}
	}
	return nil
}

func ValidateUpdateProfile(input request.UpdateProfile) error {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			return fieldError
		}
	}

	if *input.ImageUri != "" && !isImageURIExtensionValid(*input.ImageUri) {
		return errors.New("Invalid imageUri")
	}

	return nil
}

func isImageURIExtensionValid(uri string) bool {
	extension := uri[len(uri)-5:]
	if extension != ".jpeg" {
		extension = extension[1:]
		if extension != ".jpg" && extension != ".png" {
			return false
		}
	}

	return true
}
