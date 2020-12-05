package repository

func (repo *GormRepository) GetProblem(id string) (*Problem, error) {
	// not implemented yet

	problem := Problem{}
	repo.db.First(&problem, "id = ?", id)

	return &problem, nil
}

func (repo *GormRepository) GetListOfProblem() ([]*Problem, error) {
	// not implemented yet
	problems := []*Problem{}
	repo.db.Select([]string{"id", "title", "score", "solved"}).Find(&problems)
	return problems, nil
}

func (repo *GormRepository) JudgeFlag(id string, flag string) (bool, error) {
	// not implemented yet
	problem := Problem{}
	repo.db.First(&problem, "id = ?", id)

	res := (problem.Flag == flag)
	return res, nil
}
