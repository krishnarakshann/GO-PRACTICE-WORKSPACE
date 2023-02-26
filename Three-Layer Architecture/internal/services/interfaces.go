package services

import (
	"context"
	"threelayer/models"
)

type StudentService interface {
	GetStudents(context.Context) ([]models.Students, error)
	PostStudent(context.Context, models.StudentReq) (*models.Students, error)
	GetStudentById(context.Context, int) (*models.Students, error)
}
