package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"CPF" validate:"len=11, regexp=^[0-11]*$"`
	RG   string `json:"RG" validate:"len=9, regexp=^[0-9]*$"`
}

//estrutura para levar os campos a serem preenchidos pelo usuario

func ValidaDadosDeAlunos(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
