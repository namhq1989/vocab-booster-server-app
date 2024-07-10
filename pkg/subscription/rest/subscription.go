package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s server) registerSubscriptionRoutes() {
	g := s.echo.Group("/api/subscription")

	g.GET("/plans", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetSubscriptionPlansRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetSubscriptionPlans(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetSubscriptionPlansRequest](next)
	})
}
