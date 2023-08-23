package entity

type Pet struct {
	ID        int64         `json:"id"`
	Category  PetCategory   `json:"category"`
	Name      string        `json:"name"`
	PhotoUrls []string      `json:"photoUrls"`
	Tags      []PetCategory `json:"tags"`
	Status    string        `json:"status"`
}

type PetCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
