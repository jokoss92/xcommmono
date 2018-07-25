package model

type Employee struct {

	// id int(6) UN AI PK
	// name varchar(50)
	// address varchar(100)
	// city varchar(50)
	// email varchar(100)

	EmpID      int64  `json:"id"`
	EmpName    string `json:"name"`
	EmpAddress string `json:"address"`
	EmpCity    string `json:"city"`
	EmpEmail   string `json:"email"`
}
