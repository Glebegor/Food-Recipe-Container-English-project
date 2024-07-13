package routers

import (
	"backend/api/handlers"
	"backend/domain"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func SetupRouter(gin *gin.Engine, db *sqlx.DB, conf domain.Config, timeout time.Duration) {
	// User

	recipesController := handlers.RecipeElementsController{}

	recipesElementsController := handlers.RecipeElementsController{}

}
