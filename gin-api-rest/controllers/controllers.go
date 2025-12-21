package controllers

import (
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(ctx *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	ctx.JSON(200, alunos)
}

func BuscaAlunoPorID(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(ctx *gin.Context) {
	var aluno models.Aluno
	cpf := ctx.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorRG(ctx *gin.Context) {
	var aluno models.Aluno
	rg := ctx.Param("rg")
	database.DB.Where(&models.Aluno{RG: rg}).First(&aluno)

	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, aluno)
}

func Saudacao(ctx *gin.Context) {
	nome := ctx.Params.ByName("nome")
	ctx.JSON(200, gin.H{
		"API diz:": "Eai " + nome + ", Tudo certo?",
	})
}

func EditaAluno(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errr": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	ctx.JSON(http.StatusOK, aluno)
}

func DeletaAluno(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	database.DB.Delete(&aluno, id)

	ctx.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func CriaNovoAluno(ctx *gin.Context) {
	var aluno models.Aluno
	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	ctx.JSON(http.StatusOK, aluno)
}
