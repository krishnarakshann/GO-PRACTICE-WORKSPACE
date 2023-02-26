package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"threelayer/internal/services"
	"threelayer/models"
)

type Handler struct {
	svc services.StudentService
}

func New(svc services.StudentService) *Handler {
	return &Handler{svc: svc}
}

type listResponse struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Data   Students `json:"data"`
}
type Students struct {
	Students []models.Students `json:"students"`
}

type successResponse struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Data   SingleStudent `json:"data"`
}

type SingleStudent struct {
	Student *models.Students `json:"student"`
}

func (h *Handler) GetStudents(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	w.Header().Set("content-Type", "application/json")

	students, err := h.svc.GetStudents(ctx)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}

	l := listResponse{
		Code:   200,
		Status: "Success",
		Data: Students{
			Students: students,
		},
	}
	b, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`Error in marshalling the data`))
	}
	w.Write(b)

}

func (h *Handler) PostStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")

	var student models.StudentReq

	reqBody, error := io.ReadAll(r.Body)

	if error != nil {
		w.WriteHeader(400)
		w.Write([]byte(`Error in reading the request body`))
		return
	}

	err := json.Unmarshal(reqBody, &student)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`Error in reading the request body`))
		return
	}

	ctx := context.Background()
	response, error := h.svc.PostStudent(ctx, student)
	if error != nil {
		w.WriteHeader(400)
		w.Write([]byte(`There is some error`))
		return
	}
	s := successResponse{
		Code:   201,
		Status: "Success",
		Data: SingleStudent{
			Student: response,
		},
	}
	b, err := json.Marshal(s)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func (h *Handler) GetStudentById(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	w.Header().Set("content-Type", "application/json")

	path := r.URL.Path
	slice := strings.Split(path, "/")
	str_id := slice[2]
	Id, err := strconv.Atoi(str_id)

	student, err := h.svc.GetStudentById(ctx, Id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}

	s := successResponse{
		Code:   201,
		Status: "Success",
		Data: SingleStudent{
			Student: student,
		},
	}

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`Error in marshalling the data`))
	}
	w.Write(b)

}
