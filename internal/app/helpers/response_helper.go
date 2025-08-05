package app_helpers

import app_models "go-clean-achitech/internal/app/models"

func BuildSuccessResponse(message string, data any) app_models.BaseResponseModel {
	return app_models.BaseResponseModel{
		Success: true,
		Message: &message,
		Data:    data,
	}

}

func BuildFailResponse(message string, data any) app_models.BaseResponseModel {
	return app_models.BaseResponseModel{
		Success: false,
		Message: &message,
		Data:    data,
	}

}

func BuildErrorResponse(errorMessage string, data any) app_models.BaseResponseModel {
	return app_models.BaseResponseModel{
		Success:      false,
		ErrorMessage: &errorMessage,
		Data:         data,
	}

}
