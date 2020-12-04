package repository

import "github.com/gofrs/uuid"

type Problem struct {
	ID uuid.UUID `db:"id"`
	// add labels
	title string 
	score: int,
	solved: int
	star: int
}

type ProblemRepository interface {
	GetProblem(id string) (*Problem, error)
	GetListOfProblem() ([]*Problem, error)
	JudgeFlag(id string) (bool, error)
}
