package router

import (
	"net/http"
	"time"

	"github.com/CTFBox/CTFBox/repository"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wader/gormstore"
	"go.uber.org/zap"
)

// Handlers
type Handlers struct {
	Repo          repository.Repository
	Logger        *zap.Logger
	SessionKey    []byte
	SessionOption sessions.Options
	Origin        string
}

func (h *Handlers) SetupRoute(db *gorm.DB) *echo.Echo {
	echo.NotFoundHandler = func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}

	// echo初期化
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	store := gormstore.New(db, h.SessionKey)
	e.Use(session.Middleware(store))
	// db cleanup every hour
	// close quit channel to stop cleanup
	quit := make(chan struct{})
	// defer close(quit)
	go store.PeriodicCleanup(1*time.Hour, quit)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	// API定義 (/api)
	api := e.Group("/api")
	{
		apiChallenges := api.Group("/challenges")
		{
			apiChallenges.GET("/", h.HandleGetListOfProblem)
			apiChallenges.GET("/:challengeId", h.HandleGetProblem)
			apiChallenges.POST("/:challengeId", h.HandlePostFlag)
		}

	}
	e.GET("/api/oauth2/authParams", h.HandleGetAuthParams)

	return e
}
