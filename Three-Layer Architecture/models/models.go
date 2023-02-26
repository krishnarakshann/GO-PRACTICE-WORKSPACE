package models

import "time"

type Students struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	Cgpa       float64   `json:"cgpa"`
	Dob        string    `json:"date-of-birth"`
	Mobile_num string    `json:"mobile_number"`
	CreatedAt  time.Time `json:"created_at"`
}

type StudentReq struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Cgpa       float64 `json:"cgpa"`
	Dob        string  `json:"date-of-birth"`
	Mobile_num string  `json:"mobile_number"`
}
