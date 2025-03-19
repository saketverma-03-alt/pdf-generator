package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pdfmaker/internal/payslip"
)

func main() {
	var emp payslip.EmployeePayslip
	jsonData, err := os.ReadFile("./temp/employee_payslip.json")
	if err != nil {
		panic(err)
	}
	var employeeData payslip.EmployeePayrollDetail
	err = json.Unmarshal(jsonData, &employeeData)
	if err != nil {
		fmt.Println("Error marsaling to json")
		panic(err)
	}
	// fmt.Printf("Person: %+v\n", employeeData)
	emp.GenerateEmployeePayslip(&employeeData)
}
