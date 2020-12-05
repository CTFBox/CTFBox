package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) HandleGetProblem(c echo.Context) error {
	// not implemented yet

	problemID := c.Param("challengeId")
	problem, err := h.Repo.GetProblem(problemID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, problem)
}

func (h *Handlers) HandleGetListOfProblem(c echo.Context) error {
	// not implemented yet
	problems, err := h.Repo.GetListOfProblem()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, problems)
}

func (h *Handlers) HandlePostFlag(c echo.Context) error {
	// not implemented yet
	return nil
}
