package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/wintergathering/makedo/models"
	"github.com/wintergathering/makedo/reviewer"
	"github.com/wintergathering/makedo/validators"
)

type BathroomController interface {
	FindAll() []models.Bathroom
	Save(c *gin.Context) error
}

type controller struct {
	review reviewer.BathroomReviewer
}

var validate *validator.Validate

func New(r reviewer.BathroomReviewer) BathroomController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolPlace)
	return &controller{
		review: r,
	}
}

func (cn *controller) FindAll() []models.Bathroom {
	return cn.review.FindAll()
}

func (cn *controller) Save(c *gin.Context) error {
	var bathroom models.Bathroom
	err := c.ShouldBindJSON(&bathroom)
	if err != nil {
		return err
	}
	err = validate.Struct(bathroom)
	if err != nil {
		return err
	}
	cn.review.Save(bathroom)
	return nil
}
