package entity

import "database/sql"

type EmployeeInput struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
	Address   string `json:"address"`
}

type EmployeesRequest struct {
	EmpID     string `json:"empId"`
	Title     string `json:"title"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
}

type EmployeeSql struct {
	EmpID     sql.NullString `json:"id"`
	Title     sql.NullString `json:"title"`
	Firstname sql.NullString `json:"firstname"`
	Lastname  sql.NullString `json:"lastname"`
	Address   sql.NullString `json:"address"`
}
