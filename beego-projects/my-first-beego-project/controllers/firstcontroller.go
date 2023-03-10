package controllers

import (
	"github.com/astaxie/beego"
)

type FirstController struct {
	beego.Controller
}

// defining the struct
type Employee struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees []Employee

func init() {

	employees = Employees{{Id: 1, FirstName: "Foo", LastName: "Bar"},
		{Id: 2, FirstName: "Baz", LastName: "Car"},
		{Id: 3, FirstName: "Xyz", LastName: "App"}}
}

func (this *FirstController) GetEmployees() {
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = employees
	this.ServeJSON()
}

func (this *FirstController) GetEmployeeID() {

	var id int
	this.Ctx.Input.Bind(&id, "id")
	var isEmployeeExist bool
	var emps []Employee

	for _, employee := range employees {
		if employee.Id == id {
			emps = append(emps, Employee{Id: employee.Id,
				FirstName: employee.FirstName, LastName: employee.LastName})
			isEmployeeExist = true
			break
		}
	}

	if !isEmployeeExist {
		this.Abort("Generic")
	} else {
		this.Data["employees"] = emps
		this.TplName = "dashboard.tpl"
	}
}

func (this *FirstController) Dashbaord() {
	this.Data["employees"] = employees
	this.TplName = "dashboard.tpl"
}
