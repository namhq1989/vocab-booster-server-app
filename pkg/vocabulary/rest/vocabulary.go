package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s server) registerVocabularyRoutes() {
	g := s.echo.Group("/api/vocabulary")

	g.GET("/search", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.SearchVocabularyRequest)
			performerID = ctx.GetUserID()
			lang        = ctx.GetLang()
		)

		resp, err := s.app.SearchVocabulary(ctx, performerID, lang, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.SearchVocabularyRequest](next)
	})

	g.PATCH("/:id/bookmark", func(c echo.Context) error {
		var (
			ctx          = c.Get("ctx").(*appcontext.AppContext)
			req          = c.Get("req").(dto.BookmarkVocabularyRequest)
			performerID  = ctx.GetUserID()
			vocabularyID = c.Param("id")
		)

		resp, err := s.app.BookmarkVocabulary(ctx, performerID, vocabularyID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.BookmarkVocabularyRequest](next)
	})

	g.GET("/bookmarked", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetUserBookmarkedVocabulariesRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetUserBookmarkedVocabularies(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetUserBookmarkedVocabulariesRequest](next)
	})

	g.GET("/word-of-the-day", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetWordOfTheDayRequest)
			performerID = ctx.GetUserID()
			lang        = ctx.GetLang()
		)

		resp, err := s.app.GetWordOfTheDay(ctx, performerID, lang, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetWordOfTheDayRequest](next)
	})

	g.GET("/:id/community-sentences", func(c echo.Context) error {
		var (
			ctx          = c.Get("ctx").(*appcontext.AppContext)
			req          = c.Get("req").(dto.GetCommunitySentencesRequest)
			performerID  = ctx.GetUserID()
			vocabularyID = c.Param("id")
			lang         = ctx.GetLang()
		)

		resp, err := s.app.GetCommunitySentences(ctx, performerID, vocabularyID, lang, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetCommunitySentencesRequest](next)
	})

	g.GET("/community-sentences/:sentenceId", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetCommunitySentenceRequest)
			performerID = ctx.GetUserID()
			sentenceID  = c.Param("sentenceId")
			lang        = ctx.GetLang()
		)

		resp, err := s.app.GetCommunitySentence(ctx, performerID, sentenceID, lang, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetCommunitySentenceRequest](next)
	})
}
