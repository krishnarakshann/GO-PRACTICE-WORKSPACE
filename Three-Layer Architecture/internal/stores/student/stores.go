package stores

import (
	"context"
	"database/sql"
	"errors"
	"threelayer/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

const tablename = "student"

var students []models.Students

func (s *store) GetStudents(ctx context.Context) ([]models.Students, error) {

	query := "select Id,name,age,cgpa,dob,mobile_num,created_at from " + tablename
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s models.Students
		if err = rows.Scan(&s.Id, &s.Name, &s.Age, &s.Cgpa, &s.Dob, &s.Mobile_num, &s.CreatedAt); err != nil {
			return nil, err
		}
		s.Dob = s.Dob[:10]
		students = append(students, s)
	}
	return students, nil
}

func (s *store) PostStudent(ctx context.Context, studentReq models.StudentReq) (int, error) {
	query := "INSERT INTO " + tablename + "(name,age,cgpa,dob,mobile_num) values(?,?,?,?,?)"

	var args []interface{}
	args = append(args, studentReq.Name, studentReq.Age, studentReq.Cgpa, studentReq.Dob, studentReq.Mobile_num)

	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, errors.New("error in executing the query")
	}

	rows_aff, err := result.RowsAffected()
	if rows_aff != 1 {
		return int(0), errors.New("unable to execute the query")
	}

	l_id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(l_id), nil
}

func (s *store) GetStudentById(ctx context.Context, Id int) (*models.Students, error) {
	query := "select Id,name,age,cgpa,dob,mobile_num,created_at from " + tablename + " where Id = ?"
	var args []interface{}
	args = append(args, Id)

	var stud models.Students

	row := s.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&stud.Id, &stud.Name, &stud.Age, &stud.Cgpa, &stud.Dob, &stud.Mobile_num, &stud.CreatedAt); err != nil {
		return nil, err
	}
	stud.Dob = stud.Dob[:10]
	return &stud, nil
}
