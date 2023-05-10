package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Rayelisson/api-go-gin/controller"
	docs "github.com/Rayelisson/api-go-gin/docs"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/alunos"
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.POST("/alunos", controller.CriaNovoAluno)
	r.GET("/alunos/:id", controller.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.PATCH("/alunos/:id", controller.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCPF)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
