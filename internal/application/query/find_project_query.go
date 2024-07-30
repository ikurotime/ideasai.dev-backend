package query

type FindProjectQuery struct {
	ID string `json:"id"`
}

type FindProjectQueryResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
