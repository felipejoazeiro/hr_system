package sections

import (
	"app/internal/sections/repository"
	"app/internal/sections/handler"
	"app/pkg/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	repo := repository.NewDepartmentRepository(db.DB.DB)
	departmentHandler := handler.NewDepartmentHandler(repo)
	sectionsGroup := rg.Group("/sections")
	{
		sectionsGroup.GET("/", departmentHandler.GetAllDepartments)
		sectionsGroup.POST("/departments", departmentHandler.CreateDepartment)
		sectionsGroup.PUT("/departments/:id", departmentHandler.EditDepartment)
		
	}
}
