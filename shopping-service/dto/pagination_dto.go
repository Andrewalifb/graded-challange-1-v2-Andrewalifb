package dto

import "go.mongodb.org/mongo-driver/mongo/options"

type (
	PaginationRequest struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}

	PaginationData struct {
		Limit int64 `json:"limit"`
		Page  int64 `json:"page"`
	}

	PaginationResponse struct {
		Limit   int   `json:"limit"`
		Page    int   `json:"page"`
		MaxPage int64 `json:"max_page"`
		Count   int64 `json:"count"`
	}
)

func NewMongoPaginate(limit, page int) *PaginationData {
	return &PaginationData{
		Limit: int64(limit),
		Page:  int64(page),
	}
}

func (mp *PaginationData) GetPaginatedOpts() *options.FindOptions {
	l := mp.Limit
	skip := mp.Page*mp.Limit - mp.Limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	return &fOpt
}
