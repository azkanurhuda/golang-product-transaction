package handler

import (
	"backend/application/usecase"
	"backend/domain/repository"
	"backend/domain/service"
	"backend/interfaces/form"
	"backend/interfaces/presenter"
	"backend/interfaces/response"
	"net/http"
)

type AuthenticationHandler struct {
	userUseCase usecase.UserUseCase
}

func NewAuthenticationHandler(userRepository repository.User, jwtService service.Jwt) *AuthenticationHandler {
	return &AuthenticationHandler{
		userUseCase: usecase.NewUserUseCase(userRepository, jwtService),
	}
}

func (h *AuthenticationHandler) Signup(w http.ResponseWriter, r *http.Request) {
	f := form.SignUp{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := f.Validate(); err != nil {
		presenter.NewBadRequest(w)
		return
	}

	user, err := f.Entity()
	if err != nil {
		presenter.NewBadRequest(w)
		return
	}

	token, err := h.userUseCase.SignUp(r.Context(), user)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccessToken(token))
}

func (h *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	f := form.Login{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := f.Validate(); err != nil {
		presenter.NewBadRequest(w)
		return
	}

	token, err := h.userUseCase.Login(r.Context(), f.Email, f.Password)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccessToken(token))
}
