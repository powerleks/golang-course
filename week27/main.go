package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name string
	age int
	grade int
}

func printStudent(students map[string] *Student) {
	fmt.Println("Студенты из хранилища:")
	for studentName, student := range students {
		fmt.Println(studentName, student.age, student.grade)
	}
}

func newStudent(line string) Student {
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

func readLines() {
	var studentStore = make(map[string] *Student)
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			printStudent(studentStore)
			return
		}
		student := newStudent(line)
		studentStore[student.name] = &student
	}
}

func main() {
	readLines()
}