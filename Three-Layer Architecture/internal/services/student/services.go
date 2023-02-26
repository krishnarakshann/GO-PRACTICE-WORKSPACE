package services

import (
	"context"
	"threelayer/internal/stores"
	"threelayer/models"
)

type Services struct {
	store stores.StudentStorer
}

func New(store stores.StudentStorer) *Services {
	return &Services{store: store}
}

func (s *Services) GetStudents(ctx context.Context) ([]models.Students, error) {

	students, err := s.store.GetStudents(ctx)
	if err != nil {
		return nil, err
	}
	return students, err
}

func (s *Services) PostStudent(ctx context.Context, studentReq models.StudentReq) (*models.Students, error) {

	id, err := s.store.PostStudent(ctx, studentReq)
	if err != nil {
		return nil, err
	}
	student, err := s.store.GetStudentById(ctx, id)
	if err != nil {
		return nil, err
	}
	return student, err
}

func (s *Services) GetStudentById(ctx context.Context, Id int) (*models.Students, error) {
	student, err := s.store.GetStudentById(ctx, Id)
	if err != nil {
		return nil, err
	}
	return student, err
}
