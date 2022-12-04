package models

type CursoRequest struct {
	DesiredCourse  	string 	`json:"desiredCourse"`
	RequiredCourse 	string 	`json:"requiredCourse"`
}

type CursoResponse struct {
	Course	string	`json:"course"`
	Order	int		`json:"order"`
}

type Cursos struct {
	UserId  string  		`json:"userId"`
	Courses []CursoRequest 	`json:"courses"`
}

type CursosResponse struct {
	UserId  string  		`json:"userId"`
	Courses []CursoResponse `json:"courses"`
}

