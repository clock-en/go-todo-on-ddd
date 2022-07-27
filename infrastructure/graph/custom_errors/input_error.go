package custom_errors

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const INPUT_ERROR_CODE = "INPUT_ERROR"

type InputErrorModel struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func toGraphqlInputErrors(inputErrors dto.InputErrors) []InputErrorModel {
	errs := []InputErrorModel{}
	for k, v := range inputErrors {
		errs = append(errs, InputErrorModel{Field: k, Message: v.Error()})
	}
	return errs
}

func GenerateGraphqlInputError(ctx context.Context, inputErrors dto.InputErrors) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "User Input Error",
		Extensions: map[string]interface{}{
			"code":    INPUT_ERROR_CODE,
			"details": toGraphqlInputErrors(inputErrors),
		},
	}

}
