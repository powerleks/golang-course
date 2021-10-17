package main

import (
	"bufio"
	"io"
	"os"
	"example/pkg/storage"
	"example/pkg/student"
)


func readLines() {
	var studentStore = storage.NewStudentStorage();
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			studentStore.PrintStudent()
			return
		}
		student := student.NewStudent(line)
		studentStore.Put(&student)
	}
}

func main() {
	readLines()
}