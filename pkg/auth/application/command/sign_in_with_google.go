package command

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type SignInWithGoogleHandler struct {
	authenticationRepository domain.AuthenticationRepository
	jwtRepository            domain.JwtRepository
	userHub                  domain.UserHub
}

func NewSignInWithGoogleHandler(authenticationRepository domain.AuthenticationRepository, jwtRepository domain.JwtRepository, userHub domain.UserHub) SignInWithGoogleHandler {
	return SignInWithGoogleHandler{
		jwtRepository:            jwtRepository,
		authenticationRepository: authenticationRepository,
		userHub:                  userHub,
	}
}

func (h SignInWithGoogleHandler) SignInWithGoogle(ctx *appcontext.AppContext, req dto.SignInWithGoogleRequest) (*dto.SignInWithGoogleResponse, error) {
	ctx.Logger().Info("[command] new sign in with Google request", appcontext.Fields{"token": req.Token})

	ctx.Logger().Text("get user's data with Google token")
	googleUser, err := h.authenticationRepository.GetUserInfoWithToken(ctx, req.Token)
	if err != nil {
		ctx.Logger().Error("failed to get staff data with Google token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Info("Google's user found, find application's user with email via grpc", appcontext.Fields{"email": googleUser.Email})
	user, err := h.userHub.FindUserByEmail(ctx, googleUser.Email)
	if err != nil {
		ctx.Logger().Error("failed to find user by email via grpc", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, create new one")
		user, err = h.createNewUser(ctx, *googleUser)
		if err != nil {
			ctx.Logger().Error("failed to create new user", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Info("user found, generate access token", appcontext.Fields{"userID": user.ID})
	accessToken, err := h.jwtRepository.GenerateAccessToken(ctx, user.ID)
	if err != nil {
		ctx.Logger().Error("failed to generate access token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done sign in with Google request")
	return &dto.SignInWithGoogleResponse{
		AccessToken: accessToken,
	}, nil
}

func (h SignInWithGoogleHandler) createNewUser(ctx *appcontext.AppContext, googleUser domain.AuthenticationUser) (*domain.User, error) {
	ctx.Logger().Info("create new user with Google data via grpc", appcontext.Fields{"id": googleUser.UID, "email": googleUser.Email, "name": googleUser.Name})
	id, err := h.userHub.CreateUser(ctx, googleUser.Name, googleUser.Email)
	if err != nil {
		ctx.Logger().Error("failed to create new user via grpc", err, appcontext.Fields{})
		return nil, err
	}

	return &domain.User{
		ID:   id,
		Name: googleUser.Name,
	}, nil
}
