package command

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
)

type UpdateMeHandler struct {
	userRepository domain.UserRepository
}

func NewUpdateMeHandler(userRepository domain.UserRepository) UpdateMeHandler {
	return UpdateMeHandler{
		userRepository: userRepository,
	}
}

func (h UpdateMeHandler) UpdateMe(ctx *appcontext.AppContext, performerID string, req dto.UpdateMeRequest) (*dto.UpdateMeResponse, error) {
	ctx.Logger().Info("[command] new update me request", appcontext.Fields{"performerID": performerID, "name": req.Name, "bio": req.Bio, "visibility": req.Visibility})

	ctx.Logger().Text("find user by id in db")
	user, err := h.userRepository.FindUserByID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to find user by id in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found")
		return nil, apperrors.User.UserNotFound
	}

	ctx.Logger().Text("set user's data")
	if err = user.SetName(req.Name); err != nil {
		ctx.Logger().Error("failed to set user's name", err, appcontext.Fields{})
		return nil, err
	}
	if err = user.SetBio(req.Bio); err != nil {
		ctx.Logger().Error("failed to set user's bio", err, appcontext.Fields{})
		return nil, err
	}
	if err = user.SetVisibility(req.Visibility); err != nil {
		ctx.Logger().Error("failed to set user's visibility", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update user in db")
	if err = h.userRepository.UpdateUser(ctx, *user); err != nil {
		ctx.Logger().Error("failed to update user in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done update me request")
	return &dto.UpdateMeResponse{}, nil
}
