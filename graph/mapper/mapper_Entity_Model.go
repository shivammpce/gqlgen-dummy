package mapper

import (
	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/entity"
	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/model"
)

func EmployeeEntityToModelValidation(employee []*entity.EmployeeSql) []*model.Employees {
	var EmployeeLists []*model.Employees
	for _, employee := range employee {
		employeeModel := model.Employees{
			EmpID:     employee.EmpID.String,
			Title:     employee.Title.String,
			Firstname: employee.Firstname.String,
			Lastname:  employee.Lastname.String,
			Address:   employee.Address.String,
		}
		EmployeeLists = append(EmployeeLists, &employeeModel)
	}
	return EmployeeLists
}
