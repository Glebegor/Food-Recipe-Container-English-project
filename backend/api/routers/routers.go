package routers


func SetupRouter(gin *gin.Engine, db *sqlx.DB, conf Config) {
    // User
    recips := gin.Group("/recips")
    {
        recips.GET("/", )
        recips.GET("/:id", )
        recips.POST("/", )
        recips.PUT("/:id", )
        recips.DELETE("/:id", )
        recips.GET("/:id/elementsIds", )
    }

    recipsElements := gin.Group("/recipsElements")
    {
        recipsElements.GET("/", )
        recipsElements.GET("/:id", )
        recipsElements.POST("/", )
        recipsElements.PUT("/:id", )
        recipsElements.DELETE("/:id", )
    }
}