package validato

import "github.com/go-playground/validator/v10"

func GetValidationErrors(err error) map[string]string {
    errors := make(map[string]string)

    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        for _, fieldError := range validationErrors {
            errors[fieldError.Field()] = GetErrorMessage(fieldError)
        }
    }

    return errors
}

// 定义错误消息
func GetErrorMessage(fe validator.FieldError) string {
    switch fe.Tag() {
    case "required":
        return "This field is required"
    case "min":
        return "Value is too short"
    case "max":
        return "Value is too long"
    case "email":
        return "Invalid email format"
    case "gte":
        return "Value must be greater than or equal to " + fe.Param()
    case "lte":
        return "Value must be less than or equal to " + fe.Param()
    default:
        return "Invalid value"
    }
}