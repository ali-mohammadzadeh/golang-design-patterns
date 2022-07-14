package main

import "fmt"

func main() {
	developerEmployee := createEmployee("userDeveloper", DEVELOPER)
	fmt.Printf("The employee was successfully created : %+v\n", developerEmployee)

	testerEmployee := createEmployee("userTester", TESTER)
	fmt.Printf("The employee was successfully created : %+v\n", testerEmployee)
}
