package storage

import "github.com/pieash9/basic-golang/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	DeleteAStudent(id int64) (types.Student, error)
}
