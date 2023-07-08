package models

type Address struct {
	Rua    string `json:"rua"`
	Numero int    `json:"numero"`
	Cidade string `json:"cidade"`
}

type Person struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	Nome           string    `json:"nome"`
	DataNascimento string    `json:"dt_nascimento"`
	Endereco       []Address `json:"endereco"`
}
