package handlers

import (
	"backend/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RecipeElementsController struct {
	Service domain.IRecipesElementService
}

func (cr *RecipeElementsController) Get(c *gin.Context) {
	err, data := cr.Service.Get()
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{"data": data})
	return
}

func (cr *RecipeElementsController) GetById(c *gin.Context) {
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)

	err, data := cr.Service.GetById(id_int)
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{"data": data})
	return
}

func (cr *RecipeElementsController) Post(c *gin.Context) {
	var input domain.RecipesElement

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	err := cr.Service.Post(input)
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}
}

func (cr *RecipeElementsController) Delete(c *gin.Context) {
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)

	err = cr.Service.Delete(id_int)
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{"message": "Deleted"})
	return
}
