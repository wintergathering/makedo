package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wintergathering/makedo/models"
	"github.com/wintergathering/makedo/reviewer"
)

type BathroomController interface {
	FindAll() []models.Bathroom
	Save(c *gin.Context) error
	ShowAll(c *gin.Context)
	ShowByID(c *gin.Context)
}

type controller struct {
	review reviewer.BathroomReviewer
}

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

//method to pass a specific id and see that bathroom
func (cn *controller) ShowByID(c *gin.Context) {

	id := c.Param("id")

	brs := cn.review.FindAll()

	//there's likely a better way to do this, esp with a DB
	//see restaurant app maybe
	for _, a := range brs {
		if a.ID == id {
			c.HTML(http.StatusOK, "id_template.html", a)
			//c.IndentedJSON(http.StatusOK, a)
			return
			//try debugging this with c.JSON if the page doesn't seem to be loading correctly
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "bathroom not found"})
}
