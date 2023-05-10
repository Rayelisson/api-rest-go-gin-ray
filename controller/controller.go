package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"

	"github.com/Rayelisson/api-go-gin/database"
	"github.com/Rayelisson/api-go-gin/models"
)

// ExibeTodosAlunos godc
// @Summary Create new product
// @Description Mostra um usuario com o nome
// @Tags  alunos
// @Accept  json
// @Produce  json
// @Param alunos body models.alunos
// @Success 200 {object} models.alunos
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [get]

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai  ," + nome + ", tudo beleza",
	})
}

// CriaNovoAluno godc
// @Summary Criar Novo Aluno
// @Description Rota de Criacao de Um novo Aluno
// @Tags  alunos
// @Accept  json
// @Produce  json
// @Param alunos body models.alunos
// @Success 200 {object} models.alunos
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [post]

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorID godc
// @Summary Buscar Aluno Por ID
// @Description Rota de Criacao de Um novo Aluno
// @Tags  alunos
// @Accept  json
// @Produce  json
// @Param alunos body models.alunos
// @Success 200 {object} models.alunos
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/:id [get]

func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno godc
// @Summary Deletar Um Aluno
// @Description Rota de Criacao de Um novo Aluno
// @Tags  alunos
// @Accept  json
// @Produce  json
// @Param alunos body models.alunos
// @Success 200 {object} models.alunos
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/:id [delete]

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditarAluno godc
// @Summary Editar Um Aluno
// @Description Rota de Criacao de Um novo Aluno
// @Tags  alunos
// @Accept  json
// @Produce  json
// @Param alunos body models.alunos
// @Success 200 {object} models.alunos
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/:id [patch]

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF godc
// @Summary Buscar um  Aluno Por CPF
// @Description Buscar a aluno por cpf
// @Tags  alunos
// @Accept  json
// @Produce  json
// @Param alunos body models.alunos
// @Success 200 {object} models.alunos
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/:id [get]

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)

}
