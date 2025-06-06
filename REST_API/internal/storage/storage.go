package storage

import (
	"github.com/atomhudson/go_lang/student-api/internal/types"
)

type Storage interface {
	CreateStudent(name string, email string, phone string, age int) (int, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudentsList() ([]types.Student, error)
	DeleteStudent(id int64) error
	UpdateStudent(id int64, name string, email string, phone string, age int) error
}
