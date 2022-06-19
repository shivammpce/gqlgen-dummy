package mapper

import (
	"log"

	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/entity"
	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/model"
)

// This is model-to-entity.
// Mapper takes the input and put it into entity.

// If required (!)field then pointer is not required, else pointer is mandatory.

func ValidateEmployeeInput(input model.EmployeeRequest) (entity.EmployeeInput, model.Response) {
	var inputEntity entity.EmployeeInput
	var employeeValidatn model.Response
	if input.Firstname == "" {
		employeeValidatn.Error = true
		employeeValidatn.Message = "First Name cannot be blank"
		return inputEntity, employeeValidatn
	} else {
		inputEntity.Firstname = input.Firstname
	}

	if input.Lastname == "" {
		employeeValidatn.Error = true
		employeeValidatn.Message = "Last Name cannot be blank"
		return inputEntity, employeeValidatn
	} else {
		inputEntity.Lastname = input.Lastname
	}

	if input.Title == "" {
		employeeValidatn.Error = true
		employeeValidatn.Message = "Title cannot be blank"
		return inputEntity, employeeValidatn
	} else {
		inputEntity.Title = input.Title
	}

	if input.Address == "" {
		log.Println(input.Address)
		employeeValidatn.Error = true
		employeeValidatn.Message = "Address cannot be blank"
		return inputEntity, employeeValidatn
	} else {
		log.Println(input.Address)
		inputEntity.Address = input.Address
	}

	return inputEntity, employeeValidatn
}
