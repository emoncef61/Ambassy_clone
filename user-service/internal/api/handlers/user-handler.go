package handlers

import (
	"auth-service/internal/api/models"
	"auth-service/internal/customerrors"
	"auth-service/internal/db/repositories"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepository *repositories.UserRepository
}

func NewUserHandler(ur *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: ur,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		customerrors.ErrorResponse(w, r, http.StatusBadRequest, "invalid request payload")
		return
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PersonalInfo.Password), 14)
	if err != nil {
		customerrors.ErrorResponse(w, r, http.StatusInternalServerError, "password hashing failed")
		return
	}

	user.PersonalInfo.Password = hashedPassword

	if err := uh.UserRepository.Create(&user); err != nil {
		customerrors.ServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
