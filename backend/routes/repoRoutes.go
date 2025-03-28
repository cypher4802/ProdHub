package routes

import (
	controllers "prodhub-backend/controller"
	"prodhub-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RepoRoutes(router *gin.Engine) {
	repo := router.Group("/repo")
	repo.Use(middleware.Authentication())

	{
		// Repository Routes
		repo.GET("/", controllers.GetAllPublicRepos)
		repo.POST("/create", controllers.CreateRepo)	
		repo.GET("/:id", controllers.GetRepo)
		repo.PUT("/:id", controllers.UpdateRepo)
		repo.DELETE("/:id", controllers.DeleteRepo)
		repo.POST("/upload",controllers.UploadFile)

		// Branch Routes
		repo.POST("/:id/branch", controllers.CreateBranch)
		repo.GET("/:id/branch/:branchName", controllers.GetBranch)
		repo.DELETE("/:id/branch/:branchName", controllers.DeleteBranch)

		// Version Routes
		repo.POST("/:id/branch/:branchName/version", controllers.AddVersion)

	}
}