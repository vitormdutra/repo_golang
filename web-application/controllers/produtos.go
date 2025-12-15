package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"web-application/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}
		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.AtualizarProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertido)
	}
	http.Redirect(w, r, "/", 301)
}
