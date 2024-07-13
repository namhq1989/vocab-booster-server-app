package dto

type GetStatsRequest struct{}

type GetStatsResponse struct {
	Point          int64 `json:"point"`
	CompletionTime int   `json:"completionTime"`
}
