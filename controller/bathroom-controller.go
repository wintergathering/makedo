package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wintergathering/makedo/models"
	"github.com/wintergathering/makedo/reviewer"
)

type BathroomController interface {
	FindAll() []models.Bathroom
	Save(c *gin.Context) error
}

type controller struct {
	review reviewer.BathroomReviewer
}

func New(r reviewer.BathroomReviewer) BathroomController {
	return controller{
		review: r,
	}
}

func (cn controller) FindAll() []models.Bathroom {
	return cn.review.FindAll()
}

func (cn controller) Save(c *gin.Context) error {
	var bathroom models.Bathroom
	err := c.ShouldBindJSON(&bathroom)
	if err != nil {
		return err
	}
	cn.review.Save(bathroom)
	return nil
}
