package controller

import (
	"context"

	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/mapper"
	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/model"
	postgress "github.com/cbnits/gqlgen-gqlgen-dummy/graph/postgres"
)

func UpsertEmployee(input model.EmployeeRequest) (*model.Response, error) {
	response, err := mapper.ValidateEmployeeInput(input) // Calling the function from mapper package.

	if err.Error {
		return &err, nil // response cannot be used as its response type is not *model.Response.
	} else {
		finalResponse := postgress.UpsertEmployee(response)
		return &finalResponse, nil
	}
}

func GetEmployee(ctx *context.Context, input *model.EmployeeGetterRequest) (*model.EmployeeGetterResponse, error) {
	response := &model.EmployeeGetterResponse{}

	employeeEntities, err := postgress.GetEmployeeSelection(input)

	if err != nil {
		return nil, err
	}
	responseNew := mapper.EmployeeEntityToModelValidation(employeeEntities)

	response.EmployeeGet = responseNew
	return response, err
}
