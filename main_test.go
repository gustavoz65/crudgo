package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	controller "github.com/gustavoz65/api-go-gin/controllers"
	"github.com/gustavoz65/api-go-gin/database"
	"github.com/gustavoz65/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Paula",
		CPF:  "12345678910",
		RG:   "123456789",
	}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)

}

func TestVerificaSttscode(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controller.Saudacao)
	req, _ := http.NewRequest("GET", "/Gustavo", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")
	mockDaresposta := `{"API diz":"E ai Gustavo,Tudo Beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockDaresposta, string(respostaBody))

}

func TestListandotodososalunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controller.ExibeAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestBuscaAlunoPorCpf(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678910", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controller.BuscaAlunoPorId)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Paula", alunoMock.Nome)
	assert.Equal(t, "12345678910", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

// func TestBuscaPorAlunoRg(t *testing.T) {
// 	database.ConectaComBancoDeDados()
// 	r := SetupDasRotasDeTeste()
// 	r.GET("/alunos/rg/:rg", controller.BuscaAlunoPorRg)
// 	req, _ := http.NewRequest("GET", "/alunos/rg/00987657443", nil)
// 	resposta := httptest.NewRecorder()
// 	r.ServeHTTP(resposta, req)
// 	assert.Equal(t, http.StatusOK, resposta.Code)
// 	fmt.Println(resposta.Body.String())
// }

func TestEditaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controller.EditaAluno)
	aluno := models.Aluno{
		Nome: "Paula Travesti",
		CPF:  "94637395632",
		RG:   "985464763",
	}
	valorJSON, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJSON))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "Paula Travesti", alunoMockAtualizado.Nome)
	assert.Equal(t, "94637395632", alunoMockAtualizado.CPF)
	assert.Equal(t, "985464763", alunoMockAtualizado.RG)
	fmt.Println(alunoMockAtualizado.ID)
}
