package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade float64
}

// method dengan value reciver
func (s Student) GetInfo() string {
	return fmt.Sprintf("%s, %d tahun, ipk= %.2f", s.Name, s.Age, s.Grade)
}

// method dengan pointer reciver (modifikasi)
func (s *Student) UpdateGrade(newGrade float64) {
	s.Grade = newGrade
}

func main() {
	student := Student{
		Name:  "Didan",
		Age:   20,
		Grade: 3.75}

	fmt.Println(student.GetInfo())

	student.UpdateGrade(3.8)
	fmt.Println("setelah update", student.GetInfo())
}
