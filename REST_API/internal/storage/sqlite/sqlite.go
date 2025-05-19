package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/atomhudson/go_lang/student-api/internal/config"
	"github.com/atomhudson/go_lang/student-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		phone TEXT NOT NULL,
		age INTEGER NOT NULL
	)`)
	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil

}

func (s *Sqlite) CreateStudent(name string, email string, phone string, age int) (int, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, phone, age) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name, email, phone, age)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, phone, age FROM students WHERE id = ?")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()
	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student with id %d not found", id)
		}
		return types.Student{}, err
	}
	return student, nil
}

func (s *Sqlite) GetStudentsList() ([]types.Student, error) {
	rows, err := s.Db.Query("SELECT id, name, email, phone, age FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []types.Student
	for rows.Next() {
		var student types.Student
		err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *Sqlite) DeleteStudent(id int64) error {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("student with id %d not found", id)
	}
	return nil
}

func (s *Sqlite) UpdateStudent(id int64, name string, email string, phone string, age int) error {
	stmt, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, phone = ?, age = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name, email, phone, age, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("student with id %d not found", id)
	}
	return nil
}
