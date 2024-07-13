package handlers

import (
	"backend/domain"
	"github.com/gin-gonic/gin"
)

type RecipeController struct {
	Service domain.IRecipeService
}

func (cr *RecipeController) Get(c *gin.Context) {
	err, data := cr.Service.Get()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
	return
}

func (cr *RecipeController) GetById(c *gin.Context) {

}

func (cr *RecipeController) Post(c *gin.Context) {
	var input domain.Recipe

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

func (cr *RecipeController) Delete(c *gin.Context) {

}
