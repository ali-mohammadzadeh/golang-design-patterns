package main

const (
	DEVELOPER = "developer"
	TESTER    = "tester"
)

type Developer struct {
	is bool `default:false`
}
type Tester struct {
	is bool `default:false`
}
type Employee struct {
	userName string
	Developer
	Tester
}

var Employees []Employee

func createEmployee(username string, emType string) Employee {
	employee := Employee{userName: username}
	switch emType {
	case DEVELOPER:
		employee.Developer = Developer{
			is: true,
		}
		break
	case TESTER:
		employee.Tester = Tester{
			is: true,
		}
		break
	}
	return employee
}
