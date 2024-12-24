package routes

import (
	"BEKEN_UAS_PRAK/controllers"
	"BEKEN_UAS_PRAK/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// Endpoint kategori tanpa middleware
		api.GET("/kategori", controllers.GetAllKategori)
		api.POST("/kategori", controllers.CreateKategori)

		// Endpoint modul dengan middleware checkRole
		api.GET("/modul", middlewares.CheckRole("admin"), controllers.GetAllModul)
		api.POST("/modul", middlewares.CheckRole("admin"), controllers.CreateModul)
		api.GET("/user/:id/modul", controllers.GetUserModules)

		// Endpoint user dengan middleware checkJenisUser
		api.GET("/user", middlewares.CheckJenisUser("dosen"), controllers.GetAllUsers)
		api.POST("/user", middlewares.CheckJenisUser("dosen"), controllers.CreateUser)

		// Group route untuk admin
		admin := api.Group("/admin")
		admin.Use(middlewares.CheckRole("admin")) // Middleware hanya untuk admin
		{
			// CRUD pada users
			admin.POST("/users", controllers.CreateUser)
			admin.GET("/users", controllers.GetAllUsers)
			admin.PUT("/users/:id", controllers.UpdateJenisUser) // Update jenis_user
			admin.DELETE("/users/:id", controllers.DeleteUser)
		}

		// CRUD pada modul
		admin.POST("/modul", controllers.CreateModul)
		admin.GET("/modul", controllers.GetAllModul)
		admin.GET("/modul/:id", controllers.GetModulByID)
		admin.PUT("/modul/:id", controllers.UpdateModul)
		admin.DELETE("/modul/:id", controllers.DeleteModul)


		admin.GET("/user/:id/modul", controllers.GetUserModules)
		admin.PUT("/user/:id/change-jenis_user", controllers.UpdateJenisUser)


		// CRUD pada template_modul
		admin.POST("/template_modul", controllers.CreateTemplateModul)
		admin.GET("/template_modul", controllers.GetAllTemplateModul)
		admin.GET("/template_modul/:id", controllers.GetTemplateModulByID)
		admin.PUT("/template_modul/:id", controllers.UpdateTemplateModul)
		admin.DELETE("/template_modul/:id", controllers.DeleteTemplateModul)


		admin.POST("/user/:id/modul", controllers.AddModulToUser)    // Create
		admin.PUT("/user/:id/modul/:modul_id", controllers.UpdateUserModul) // Update
		admin.DELETE("/user/:id/modul/:modul_id", controllers.DeleteUserModul) // Delete

		

	}

	return r
}
