package query

type FindProjectByIDQuery struct {
	ID string `json:"id"`
}

type FindProjectByIDQueryResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
