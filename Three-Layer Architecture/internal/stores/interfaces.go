package stores

import (
	"context"
	"threelayer/models"
)

type StudentStorer interface {
	GetStudents(context.Context) ([]models.Students, error)
	PostStudent(context.Context, models.StudentReq) (int, error)
	GetStudentById(context.Context, int) (*models.Students, error)
}
