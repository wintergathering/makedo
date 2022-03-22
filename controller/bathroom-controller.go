package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/wintergathering/makedo/models"
	"github.com/wintergathering/makedo/reviewer"
)

type BathroomController interface {
	FindAll() []models.Bathroom
	Save(c *gin.Context) error
	ShowAll(c *gin.Context)
}

type controller struct {
	review reviewer.BathroomReviewer
}

var validate *validator.Validate

func New(r reviewer.BathroomReviewer) BathroomController {
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
	// err = validate.Struct(bathroom)
	// if err != nil {
	// 	return err
	// }
	cn.review.Save(bathroom)
	return nil
}

func (cn *controller) ShowAll(c *gin.Context) {
	bathrooms := cn.review.FindAll()
	data := gin.H{
		"title":     "Places to Poop",
		"bathrooms": bathrooms,
	}
	c.HTML(http.StatusOK, "index.html", data)
}
