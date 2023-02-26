package stores

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"regexp"
	"testing"
	"threelayer/models"
	"time"
)

func TestStore_GetStudents(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error occured '#{err}' while opening the connection")
	}
	store := New(db)
	defer db.Close()

	testcases := []struct {
		description string
		expouptut   []models.Students
		experror    error
		mock        *sqlmock.ExpectedQuery
	}{
		{
			description: "Success case",
			expouptut: []models.Students{
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
			experror: nil,
			mock: mock.ExpectQuery(regexp.QuoteMeta("select Id,name,age,cgpa,dob,mobile_num,created_at from " + tablename)).WillReturnRows(sqlmock.NewRows([]string{"Id",
				"name", "age", "cgpa", "dob", "mobile_num", "created_at"}).AddRow(1, "abc", 23, 9.9, "2002-02-20", "7995805080",
				time.Date(2023, time.February, 03, 13, 02, 58, 0, time.UTC))),
		},
		{
			description: "Error case",
			expouptut:   nil,
			experror:    errors.New("Query Error"),
			mock:        mock.ExpectQuery(regexp.QuoteMeta("select Id,name,age,cgpa,dob,mobile_num,created_at from " + tablename)).WillReturnError(errors.New("Query Error")),
		},
	}
	for i, tc := range testcases {
		ctx := context.Background()
		students, err := store.GetStudents(ctx)

		fmt.Printf(" Executing Testcase %d \n", i+1)

		if !reflect.DeepEqual(err, tc.experror) {
			t.Errorf("Expected Error  %v \n but got Error \n %v \n ", tc.experror, err)
		}

		if !reflect.DeepEqual(students, tc.expouptut) {
			t.Errorf("Expected %v but got %v \n", tc.expouptut, students)
		}

	}
}

func initiateTest(t *testing.T) (db *sql.DB, mock sqlmock.Sqlmock, store *store) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Unexpected Error occured")
	}
	store = New(db)
	return db, mock, store
}

func TestStore_PostStudent(t *testing.T) {

	db, mock, store := initiateTest(t)

	defer db.Close()

	testcases := []struct {
		description    string
		input          models.StudentReq
		expectedOutput int
		expectedError  error
		mock           *sqlmock.ExpectedExec
	}{
		//{
		//	description: "Successcase",
		//	input: models.StudentReq{
		//		Name:       "abc",
		//		Age:        23,
		//		Cgpa:       9.9,
		//		Dob:        "2002-02-20",
		//		Mobile_num: "9948818186",
		//	},
		//	expectedOutput: 1,
		//	expectedError:  nil,
		//	mock:           mock.ExpectExec(regexp.QuoteMeta("INSERT INTO student(name,age,cgpa,dob,mobile_num) values(?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1)),
		//},
		//{
		//	description: "Error case of arguement count mismatch",
		//	input: models.StudentReq{
		//		Name:       "abc",
		//		Age:        23,
		//		Cgpa:       9.9,
		//		Dob:        "2002-02-20",
		//		Mobile_num: "9948818186",
		//	},
		//	expectedError:  errors.New("error in executing the query"),
		//	expectedOutput: 0,
		//	mock:           mock.ExpectExec(regexp.QuoteMeta("INSERT INTO student(name,age,cgpa,dob,mobile_num) values(?,?,?,?,?)")).WillReturnError(errors.New("error in executing the query")),
		//},
		{
			description: "Error case ",
			input: models.StudentReq{
				Name:       "abc",
				Age:        23,
				Cgpa:       9.9,
				Dob:        "2002-02-20",
				Mobile_num: "9948818186",
			},
			expectedError:  errors.New("unable to execute the query"),
			expectedOutput: 0,
			mock: mock.ExpectExec(regexp.QuoteMeta("INSERT INTO student(name,age,cgpa,dob,mobile_num) values(?,?,?,?,?)")).
				WillReturnResult(sqlmock.NewResult(0, 0)).WillReturnError(nil),
		},
	}

	for i, tc := range testcases {
		fmt.Println("Executing testcase ", i+1)

		ctx := context.Background()
		response, err := store.PostStudent(ctx, tc.input)

		if !reflect.DeepEqual(tc.expectedError, err) {
			t.Errorf("Expected %v \n but got %v \n", tc.expectedError, err)
		}

		if !reflect.DeepEqual(tc.expectedOutput, response) {
			t.Errorf("Expected %v \n but got %v \n ", tc.expectedOutput, response)
		}

	}

}
