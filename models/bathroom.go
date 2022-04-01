package models

type Bathroom struct {
	ID     string `json:"id"`
	Place  string `json:"place" binding:"required,min=2"`
	Rating uint   `json:"rating" binding:"required"`
	Review string `json:"review"`
	Author Person `json:"author" binding:"required"`
}

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}
