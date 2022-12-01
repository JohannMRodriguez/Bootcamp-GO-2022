package models

type Curso struct {
	DesiredCourse  string `json:"desiredCourse"`
	RequiredCourse string `json:"requiredCourse"`
}

type Cursos struct {
	UserId  string  `json:"userId"`
	Courses []Curso `json:"courses"`
}

