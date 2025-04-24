package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/pieash9/basic-golang/students-api/internal/config"
	"github.com/pieash9/basic-golang/students-api/internal/types"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS student (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			age INTEGER NOT NULL
		);
	`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO student (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM student WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(
		&student.Id,
		&student.Name,
		&student.Email,
		&student.Age,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student with id %d not found", id)
		}
		return types.Student{}, fmt.Errorf("query error : %w", err)
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM student")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)

		if err != nil {
			return nil, err
		}

		students = append(students, student)

	}

	return students, nil
}

func (s *Sqlite) DeleteAStudent(id int64) (types.Student, error) {
	student, err := s.GetStudentById(id)
	if err != nil {
		return types.Student{}, err
	}

	stmt, err := s.Db.Prepare("DELETE FROM student where id = ?")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return types.Student{}, err
	}

	return student, nil
}

func (s *Sqlite) UpdateStudent(id int64, name string, email string, age int) (types.Student, error) {
	_, err := s.GetStudentById(id)
	if err != nil {
		return types.Student{}, err
	}

	stmt, err := s.Db.Prepare("UPDATE student SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, email, age, id)
	if err != nil {
		return types.Student{}, err
	}

	return s.GetStudentById(id)
}
