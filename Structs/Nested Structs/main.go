package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type College struct {
	Students []Student `json:"Students"`
}
type Student struct {
	Name    string             `json:"name"`
	Age     int                `json:"age"`
	Hobbies []string           `json:"hobbies"`
	sgpa    map[string]float64 `json:"sgpa"`
	Dob     DOB                `json:"date-of-birth"`
	Job     []Jobs             `json:"jobs"`
}
type DOB struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Jobs struct {
	Company  string `json:"company"`
	Duration int    `json:"duration"`
}

func main() {

	s := College{
		Students: []Student{
			{
				Name:    "Krishna rakshann",
				Age:     21,
				Hobbies: []string{"reading,writing,playing"},
				sgpa: map[string]float64{
					"First Sem":  8.8,
					"Second Sem": 8.9,
				},
				Dob: DOB{
					City:  "Guntur",
					State: "Andhra Pradesh",
				},
				Job: []Jobs{
					{
						Company:  "Zopsmart",
						Duration: 6,
					},
					{
						Company:  "Samsung",
						Duration: 1,
					},
				},
			},
		},
	}

	sbyte, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(sbyte))

	//aceesing the nested struct fields
	fmt.Println(s.Students[0])
	fmt.Println(s.Students[0].Name)
	fmt.Println(s.Students[0].Job[0].Company)
}
