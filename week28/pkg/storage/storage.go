package storage

import (
	"example/pkg/student"
	"fmt"
)

type StudentStorage map[string] *student.Student

func NewStudentStorage() StudentStorage {
	return make(map[string] *student.Student)
}

func (studentStorage StudentStorage) Put(s *student.Student) {
	studentStorage[s.Name] = s
}

func (studentStorage StudentStorage) PrintStudent() {
	fmt.Println("Студенты из хранилища:")
	for _, s := range studentStorage {
		student.PrintStudent(s)
	}
}