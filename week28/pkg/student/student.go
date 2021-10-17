package student

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name string
	Age int
	Grade int
}

func NewStudent(line string) Student {
	args := strings.Fields(line)
	name := args[0]
	age, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Errorf("Error while parsing")
	}
	grade, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Errorf("Error while parsing")
	}
	return Student{name, age, grade}
}

func PrintStudent(student *Student) {
	fmt.Println(student.Name, student.Age, student.Grade)
}