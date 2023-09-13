package handler

import (
	"backend/application/usecase"
	"backend/domain/repository"
	"backend/domain/service"
	"backend/interfaces/middleware"
	"backend/interfaces/presenter"
	"backend/interfaces/response"
	"net/http"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userRepository repository.User, jwtService service.Jwt) *UserHandler {
	return &UserHandler{
		userUseCase: usecase.NewUserUseCase(userRepository, jwtService),
	}
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		presenter.NewUnauthorized(w)
		return
	}

	user, err := h.userUseCase.Me(r.Context(), userID)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewUser(user))
}
