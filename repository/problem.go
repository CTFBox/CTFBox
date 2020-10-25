package repository

import "github.com/gofrs/uuid"

type Problem struct {
	ID uuid.UUID `db:"id"`
	// add labels
}

type ProblemRepository interface {
	GetProblem(id string) (*Problem, error)
}
