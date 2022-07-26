package custom_errors

import "github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"

const INPUT_ERROR_CODE = "INPUT_ERROR"

type InputErrorModel struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ToGraphqlInputErrors(inputErrors dto.InputErrors) []InputErrorModel {
	errs := []InputErrorModel{}
	for k, v := range inputErrors {
		errs = append(errs, InputErrorModel{Field: k, Message: v.Error()})
	}
	return errs
}
