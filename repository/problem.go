package repository

import "github.com/gofrs/uuid"

type Problem struct {
	ID uuid.UUID `db:"id" json:"challengeID"`
	// add labels
	Title  string
	Flag   string
	Score  int
	Solved int
	Star   int
}

type ProblemRepository interface {
	GetProblem(id string) (*Problem, error)
	GetListOfProblem() ([]*Problem, error)
	JudgeFlag(id string, flag string) (bool, error)
}
