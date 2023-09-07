package entity

type Pet struct {
	ID        int64         `json:"id" db:"id"`
	Category  PetCategory   `json:"category" db:"category"`
	Name      string        `json:"name" db:"name"`
	PhotoUrls []string      `json:"photoUrls" db:"photoUrls"`
	Tags      []PetCategory `json:"tags"`
	Status    string        `json:"status" db:"status"`
}

type PetCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
