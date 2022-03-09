package reviewer

import "github.com/wintergathering/makedo/models"

type BathroomReviewer interface {
	Save(models.Bathroom) models.Bathroom
	FindAll() []models.Bathroom
}

type bathroomReview struct {
	bathrooms []models.Bathroom
}

func New() BathroomReviewer {
	return &bathroomReview{}
}

func (br *bathroomReview) Save(b models.Bathroom) models.Bathroom {
	br.bathrooms = append(br.bathrooms, b)
	return b
}

func (br *bathroomReview) FindAll() []models.Bathroom {
	return br.bathrooms
}
