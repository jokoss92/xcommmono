package route

import (
	"net/http"
	"xcomm/controller"
)

func Routing() {
	http.HandleFunc("/", controller.EmployeeIndexHandler)
	http.HandleFunc("/list", controller.EmployeeListHandler)
	http.HandleFunc("/create", controller.EmployeeCreateHandler)
	http.HandleFunc("/update", controller.EmployeeUpdateHandler)
	http.HandleFunc("/delete", controller.EmployeeDeleteHandler)
	http.Handle("/content/",
		http.StripPrefix("/content/", http.FileServer(http.Dir("./content"))),
	)
}
