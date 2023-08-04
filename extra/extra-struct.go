// Another struct embedding
package main

import "fmt"

type School struct {
	SchoolName  string
	SchoolMotto string
	SchoolStuff []string
}

type Teacher struct {
	School      // has features of School
	TeacherName string
	Activity    string
}

type Student struct {
	School      // has feature of School
	StudentName string
	Objective   string
}

func (t Teacher) Marks() {
	fmt.Printf("School Name: %v, Motto: %v, Stuff: %v ", t.School.SchoolName, t.School.SchoolMotto, t.School.SchoolStuff)
	fmt.Printf("Teacher Name: %v, Activity: %v", t.TeacherName, t.Activity)
}

func (s Student) Study() {
	fmt.Printf("School Name: %v, Motto: %v, Stuff: %v ", s.School.SchoolName, s.School.SchoolMotto, s.School.SchoolStuff)
	fmt.Printf("Student Name: %v, Objective: %v", s.StudentName, s.Objective)
}

func main() {
	teacher := Teacher{
		School: School{
			SchoolName:  "Soroti Dem",
			SchoolMotto: "Math is key",
			SchoolStuff: []string{"Peter", "William"},
		},
		TeacherName: "Omoding",
		Activity:    "Teach",
	}

	teacher.Marks()

	student := Student{
		School: School{
			SchoolName:  "Agule",
			SchoolMotto: "Education is key",
			SchoolStuff: []string{"okwii", "emma"},
		},
		StudentName: "Okwiir",
		Objective:   "Study",
	}
	student.Study()
}
