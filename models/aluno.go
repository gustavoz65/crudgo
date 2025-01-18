package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"CPF"`
	RG   string `json:"RG"`
}

//estrutura para levar os campos a serem preenchidos pelo usuario
