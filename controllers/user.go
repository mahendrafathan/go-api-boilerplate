package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mahendrafathan/go-api-boilerplate/helpers"
	"github.com/mahendrafathan/go-api-boilerplate/models"
	"github.com/mahendrafathan/go-api-boilerplate/repository"
)

var error helpers.CustomError

type UserController struct {
	UserRepo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) UserController {
	return UserController{
		UserRepo: repo,
	}
}

type ResponseOutput struct {
	User  models.User
	Token string
}

func (u UserController) SignupUser(w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	json.NewDecoder(r.Body).Decode(&User)

	if len(User.Name) < 3 {
		error.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters long!")
		return
	}

	if len(User.Username) < 3 {
		error.ApiError(w, http.StatusBadRequest, "Username should be at least 3 characters long!")
		return
	}

	if len(User.Email) < 3 {
		error.ApiError(w, http.StatusBadRequest, "Email should be at least 3 characters long!")
		return
	}

	if len(User.Password) < 3 {
		error.ApiError(w, http.StatusBadRequest, "Password should be at least 3 characters long!")
		return
	}

	if result := u.UserRepo.Create(User); result.Error != nil {
		error.ApiError(w, http.StatusInternalServerError, "Failed To Add new User in database! \n"+result.Error.Error())
		return
	}

	payload := helpers.Payload{
		Username: User.Username,
		Email:    User.Email,
		Id:       User.ID,
	}

	token, err := helpers.GenerateJwtToken(payload)
	if err != nil {
		error.ApiError(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
		return
	}

	helpers.RespondWithJSON(w, ResponseOutput{
		Token: token,
		User:  User,
	})

}

func (u UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	User := models.User{}

	type Credentials struct {
		Id       string
		Password string
	}

	credentials := Credentials{}
	json.NewDecoder(r.Body).Decode(&credentials)
	if len(credentials.Id) < 3 {
		error.ApiError(w, http.StatusBadRequest, "Invalid Username/Email!")
		return
	}

	if len(credentials.Password) < 3 {
		error.ApiError(w, http.StatusBadRequest, "Invalid Password!")
		return
	}

	if results := u.UserRepo.FindByUsernameOrEmail(credentials.Id).First(&User); results.Error != nil || results.RowsAffected < 1 {
		error.ApiError(w, http.StatusNotFound, "Invalid Username/Email, Please Signup!")
		return
	}

	if err := User.Validate(credentials.Password); err != nil {
		error.ApiError(w, http.StatusNotFound, "Invalid Credentials!")
		return
	}

	payload := helpers.Payload{
		Username: User.Username,
		Email:    User.Email,
		Id:       User.ID,
	}

	token, err := helpers.GenerateJwtToken(payload)
	if err != nil {
		error.ApiError(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
		return
	}

	helpers.RespondWithJSON(w, ResponseOutput{
		Token: token,
		User:  User,
	})

}
