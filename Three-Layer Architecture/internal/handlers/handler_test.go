package handlers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"threelayer/internal/services"
	"threelayer/models"
	"time"
)

func initiateTest(t *testing.T) (ctrl *gomock.Controller, mockService *services.MockStudentService, h *Handler) {

	ctrl = gomock.NewController(t)
	mockService = services.NewMockStudentService(ctrl)
	h = New(mockService)
	return
}

func TestHandler_GetStudents(t *testing.T) {
	ctrl, mock, handler := initiateTest(t)
	defer ctrl.Finish()

	testcases := []struct {
		description    string
		expectedOutput []byte
		mock           *gomock.Call
	}{
		{
			description:    "success case",
			expectedOutput: []byte(`{"code":200,"status":"Success","data":{"students":[{"id":1,"name":"krishna","age":22,"cgpa":8.9,"date-of-birth":"2002-02-20","mobile_number":"7995805080","created_at":"2023-02-25T12:22:08Z"}]}}`),
			mock: mock.EXPECT().GetStudents(gomock.Any()).Return([]models.Students{
				{
					Id:         1,
					Name:       "krishna",
					Age:        22,
					Cgpa:       8.9,
					Dob:        "2002-02-20",
					Mobile_num: "7995805080",
					CreatedAt:  time.Date(2023, time.February, 25, 12, 22, 8, 0, time.UTC),
				},
			}, nil),
		},
	}
	for i, tc := range testcases {

		fmt.Println("Executing Testcase ", i+1)

		r := httptest.NewRequest(http.MethodGet, "/getall", nil)
		w := httptest.NewRecorder()

		handler.GetStudents(w, r)

		result := w.Result()
		b, _ := io.ReadAll(result.Body)

		if !reflect.DeepEqual(b, tc.expectedOutput) {
			t.Errorf("Expected %v \n but got %v \n", string(tc.expectedOutput), string(b))
		}

	}
}
