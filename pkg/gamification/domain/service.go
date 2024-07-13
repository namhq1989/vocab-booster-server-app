package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type Service interface {
	AddPoint(ctx *appcontext.AppContext, point Point) error
}
