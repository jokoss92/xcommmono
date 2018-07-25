package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"xcomm/config"
	"xcomm/model"
)

func EmployeeIndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list", 301)
}

func EmployeeListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}

	db, err := config.Connect()
	if err != nil {
		config.CheckError(err)
		return
	}

	fmt.Println("Database connect successfully.")
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_m_employee")
	config.CheckInternalServerError(err, w)
	var funcMap = template.FuncMap{
		"addOne": func(n int) int {
			return n + 1
		},
	}

	var employees []model.Employee
	var employee model.Employee
	for rows.Next() {
		err = rows.Scan(&employee.EmpID, &employee.EmpName,
			&employee.EmpAddress, &employee.EmpCity, &employee.EmpEmail)
		config.CheckInternalServerError(err, w)
		employees = append(employees, employee)
	}
	t, err := template.New("index.html").Funcs(funcMap).ParseFiles("view/employee/index.html")
	config.CheckInternalServerError(err, w)
	err = t.Execute(w, employees)
	config.CheckInternalServerError(err, w)
}

func EmployeeCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}

	db, err := config.Connect()
	if err != nil {
		config.CheckError(err)
		return
	}

	fmt.Println("Database connect successfully.")
	defer db.Close()

	var employee model.Employee

	employee.EmpName = r.FormValue("EmpName")
	employee.EmpAddress = r.FormValue("EmpAddress")
	employee.EmpCity = r.FormValue("EmpCity")
	employee.EmpEmail = r.FormValue("EmpEmail")
	fmt.Println(employee)

	// Save to database
	stmt, err := db.Prepare(`INSERT INTO tbl_m_employee(name, address, city, email)
		VALUES(?, ?, ?, ?)`)
	if err != nil {
		fmt.Println("Prepare query error")
		panic(err)
	}
	_, err = stmt.Exec(employee.EmpName, employee.EmpAddress,
		employee.EmpCity, employee.EmpEmail)
	if err != nil {
		fmt.Println("Execute query error")
		panic(err)
	}
	http.Redirect(w, r, "/", 301)
}

func EmployeeUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}

	db, err := config.Connect()
	if err != nil {
		config.CheckError(err)
		return
	}

	fmt.Println("Database connect successfully.")
	defer db.Close()

	var employee model.Employee

	employee.EmpID, _ = strconv.ParseInt(r.FormValue("EmpID"), 10, 64)
	employee.EmpName = r.FormValue("EmpName")
	employee.EmpAddress = r.FormValue("EmpAddress")
	employee.EmpCity = r.FormValue("EmpCity")
	employee.EmpEmail = r.FormValue("EmpEmail")
	fmt.Println(employee)

	stmt, err := db.Prepare(`
		UPDATE tbl_m_employee SET name=?, address=?, city=?, email=?
		WHERE id=?
	`)
	config.CheckInternalServerError(err, w)
	res, err := stmt.Exec(employee.EmpName, employee.EmpAddress,
		employee.EmpCity, employee.EmpEmail, employee.EmpID)
	config.CheckInternalServerError(err, w)
	_, err = res.RowsAffected()
	config.CheckInternalServerError(err, w)
	http.Redirect(w, r, "/", 301)
}

func EmployeeDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}

	db, err := config.Connect()
	if err != nil {
		config.CheckError(err)
		return
	}

	fmt.Println("Database connect successfully.")
	defer db.Close()

	var EmpID, _ = strconv.ParseInt(r.FormValue("EmpID"), 10, 64)
	stmt, err := db.Prepare("DELETE FROM tbl_m_employee WHERE id=?")
	config.CheckInternalServerError(err, w)
	res, err := stmt.Exec(EmpID)
	config.CheckInternalServerError(err, w)
	_, err = res.RowsAffected()
	config.CheckInternalServerError(err, w)
	http.Redirect(w, r, "/", 301)
}
