package handlers

import (
	"github.com/mant1COREX/pet-project/internal/entity"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func isTaskIputValid(task entity.Task) error{
	return validate.Struct(task)
}

