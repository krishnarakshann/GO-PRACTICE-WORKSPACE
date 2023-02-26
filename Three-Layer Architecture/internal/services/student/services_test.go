package services

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"threelayer/internal/stores"
	"threelayer/models"
	"time"
)

func initiateTest(t *testing.T) (ctrl *gomock.Controller, mockStore *stores.MockStudentStorer, s *Services) {
	ctrl = gomock.NewController(t)
	mockStore = stores.NewMockStudentStorer(ctrl)
	s = New(mockStore)
	return
}

func TestServices_GetStudents(t *testing.T) {
	ctrl, mockStore, service := initiateTest(t)
	defer ctrl.Finish()

	testcases := []struct {
		description    string
		expectedOutput []models.Students
		expectedError  error
		mock           *gomock.Call
	}{
		{
			description: "Success case ",
			expectedOutput: []models.Students{
				{
					Id:         1,
					Name:       "abc",
					Age:        23,
					Cgpa:       9.9,
					Dob:        "2002-02-20",
					Mobile_num: "7995805080",
					CreatedAt:  time.Date(2023, time.February, 03, 13, 02, 58, 0, time.UTC),
				},
			},
			expectedError: nil,
			mock: mockStore.EXPECT().GetStudents(gomock.Any()).Return([]models.Students{
				{
					Id:         1,
					Name:       "abc",
					Age:        23,
					Cgpa:       9.9,
					Dob:        "2002-02-20",
					Mobile_num: "7995805080",
					CreatedAt:  time.Date(2023, time.February, 03, 13, 02, 58, 0, time.UTC),
				},
			}, nil),
		},
	}

	for i, tc := range testcases {
		fmt.Println("Executing Testcase ", i+1)

		ctx := context.Background()
		students, error := service.GetStudents(ctx)

		if !reflect.DeepEqual(tc.expectedError, error) {
			t.Errorf("Exepcetd %v \n but got %v \n ", tc.expectedError, error)
		}

		if !reflect.DeepEqual(tc.expectedOutput, students) {
			t.Errorf("Expected %v \n but got %v \n ", tc.expectedOutput, students)
		}
	}
}
