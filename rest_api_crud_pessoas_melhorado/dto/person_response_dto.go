package dto

type PersonResponseDTO struct {
	ID   string `json:"id" bson:"_id"`
	Nome string `json:"nome" bson:"nome"`
}
