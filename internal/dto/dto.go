package dto

type RequestDTO struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Function       bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}
