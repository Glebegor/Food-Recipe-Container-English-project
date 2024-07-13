package routers

import (
	"backend/api/handlers"
	"backend/domain"
	"backend/repositories"
	"backend/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func SetupRouter(gin *gin.Engine, db *sqlx.DB, conf domain.Config, timeout time.Duration) {
	recipesRepository := repositories.RecipesRepository{Db: db}
	recipesService := services.RecipesService{Repo: &recipesRepository}
	recipesController := handlers.RecipeElementsController{Service: recipesService}

	recipesGroup := gin.Group("/recipes")
	{
		recipesGroup.GET("/", recipesController.Get)
		recipesGroup.GET("/:id", recipesController.GetById)
		//recipesGroup.GET("/:id/elements", recipesController.GetEl)
		recipesGroup.POST("/", recipesController.Post)
		recipesGroup.PUT("/:id", recipesController.Put)
		recipesGroup.DELETE("/:id", recipesController.Delete)
	}

	recipesElementsRepository := repositories.RecipesElementRepository{Db: db}
	recipesElementsService := services.RecipesElementService{Repo: &recipesElementsRepository}
	recipesElementsController := handlers.RecipeElementsController{Service: recipesElementsService}

	recipesElementsGroup := gin.Group("/recipesElements")
	{
		recipesElementsGroup.GET("/", recipesElementsController.Get)
		recipesElementsGroup.GET("/:id", recipesElementsController.GetById)
		recipesElementsGroup.POST("/", recipesElementsController.Post)
		recipesElementsGroup.PUT("/:id", recipesElementsController.Put)
		recipesElementsGroup.DELETE("/:id", recipesElementsController.Delete)
	}
}
