package main

import (
	"bytes"
	"encoding/json"
	"gin-api-rest/controllers"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTest()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/vitor", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria ser igual")
	mockDaResposta := `{"API diz:":"Eai vitor, Tudo certo?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectarComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTest()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectarComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTest()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/31294912", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectarComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTest()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do aluno Teste", alunoMock.Nome, "Os nomes não estão iguais")
	assert.Equal(t, "12345678901", alunoMock.CPF, "Os CPFs não estão iguais")
	assert.Equal(t, "123456789", alunoMock.RG, "Os RGs não estão iguais")
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectarComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTest()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestAtualizarAlunoHandler(t *testing.T) {
	database.ConectarComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTest()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome do aluno Teste", CPF: "12345678911", RG: "123456700"}
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "12345678911", alunoMockAtualizado.CPF)
	assert.Equal(t, "Nome do aluno Teste", alunoMockAtualizado.Nome)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
}
