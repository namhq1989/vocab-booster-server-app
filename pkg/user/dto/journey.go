package dto

import "github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"

type Journey struct {
	ID         string `json:"id"`
	Lang       string `json:"lang"`
	IsLearning bool   `json:"isLearning"`
}

func (Journey) FromDomain(journey domain.Journey) Journey {
	return Journey{
		ID:         journey.ID,
		Lang:       journey.Lang.String(),
		IsLearning: journey.IsLearning,
	}
}
