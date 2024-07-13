package handlers

import (
	"backend/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RecipeController struct {
	Service domain.IRecipeService
}

func (cr *RecipeController) Get(c *gin.Context) {
	err, data := cr.Service.Get()

	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{"data": data})
	return
}

func (cr *RecipeController) GetById(c *gin.Context) {
	id := c.Param("id")

	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	err, data := cr.Service.GetById(id_int)
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{"data": data})
	return
}

func (cr *RecipeController) Post(c *gin.Context) {
	fmt.Print("OKKKK")

	var input domain.Recipe

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	err := cr.Service.Post(input)
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{"Status": "OK"})
	return
}

func (cr *RecipeController) Delete(c *gin.Context) {
	id := c.Param("id")

	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	err = cr.Service.Delete(id_int)
	if err != nil {
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{"Status": "OK"})
	return
}
