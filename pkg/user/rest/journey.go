package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s server) registerJourneyRoutes() {
	g := s.echo.Group("/api/user/journey")

	g.GET("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetJourneysRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetJourneys(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetJourneysRequest](next)
	})

	g.POST("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.CreateJourneyRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.CreateJourney(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.CreateJourneyRequest](next)
	})

	g.PATCH("/switch", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.SwitchJourneyRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.SwitchJourney(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.SwitchJourneyRequest](next)
	})
}
