package postgress

import (
	"context"
	"log"

	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/entity"
	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/model"
	"github.com/jmoiron/sqlx"
)

func UpsertEmployee(input entity.EmployeeInput) (output model.Response) {
	if pool == nil { // System Variable, Creates the database instances.
		pool = GetPool()
	}
	log.Println(input.Address)
	queryString := `INSERT INTO employees (first_name, last_name, title, address  ) VALUES ($1, $2, $3, $4) `                    // ? = Dyanamic parameters
	commandTag, err := pool.Exec(context.Background(), queryString, input.Firstname, input.Lastname, input.Title, input.Address) // This is the syntex.
	if err != nil {
		output.Error = true
		output.Message = err.Error()
		return output
	} else if commandTag.RowsAffected() < 1 { // if no row is formed.
		output.Error = true
		output.Message = "No row inserted."
		return output
	} else {
		output.Error = false
		output.Message = "Employee Inserted Successful"
		return output
	}
}

func GetEmployeeSelection(input *model.EmployeeGetterRequest) ([]*entity.EmployeeSql, error) {
	if pool == nil {
		pool = GetPool() // Starting a database instance.
	}

	var entities []*entity.EmployeeSql

	var inputArgs []interface{}

	query := `SELECT employee_id, title, first_name, last_name, address  FROM employees where 1 = 1 `

	if input.Title != nil {
		if *input.Title != "" {
			query = query + ` AND title = ?`
			inputArgs = append(inputArgs, *input.Title)
		}

	}

	if input.EmpID != nil {
		if *input.EmpID != "" {
			query = query + ` AND employee_id = ?`
			inputArgs = append(inputArgs, *input.EmpID)
		}

	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)
	log.Println(query)
	log.Println(inputArgs...)
	rows, err := pool.Query(context.Background(), query, inputArgs...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		employeeEntity := entity.EmployeeSql{}
		err := rows.Scan(
			&employeeEntity.EmpID,
			&employeeEntity.Title,
			&employeeEntity.Firstname,
			&employeeEntity.Lastname,
			&employeeEntity.Address,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(employeeEntity)
		entities = append(entities, &employeeEntity)
	}

	return entities, nil
}