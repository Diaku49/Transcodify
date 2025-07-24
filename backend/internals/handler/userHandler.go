package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	Jwt "github.com/Diaku49/FoodOrderSystem/backend/internals/JwtService"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/constants"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/email"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/model"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/repository"
	util "github.com/Diaku49/FoodOrderSystem/backend/utilities"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserRepository *repository.UserRepository
	Validate       *validator.Validate
	Key            string
}

func NewUH(db *gorm.DB) *UserHandler {
	validate := validator.New()
	key := os.Getenv("JWT_SECRET")
	return &UserHandler{
		UserRepository: &repository.UserRepository{DB: db},
		Validate:       validate,
		Key:            key,
	}
}

func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var payload model.UserSignupPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteJsonError(w, "couldn parse payload", http.StatusBadRequest, err)
		return
	}

	//Validating
	if err := uh.Validate.Struct(payload); err != nil {
		util.WriteJsonError(w, "Validation failed", http.StatusBadRequest, err)
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		util.WriteJsonError(w, "Validation failed", http.StatusBadRequest, err)
		return
	}

	user := model.User{
		UserName: payload.UserName,
		Email:    payload.Email,
		Password: string(hashPass),
	}
	err = uh.UserRepository.CreateUser(&user)
	if err != nil {
		http.Error(w, "db failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteJsonSuccess(w, http.StatusCreated, model.SuccessResponse{
		Message: "User successfully signedup.",
	})
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload model.UserLoginPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteJsonError(w, "couldnt parse payload", http.StatusBadRequest, err)
		return
	}

	if err := uh.Validate.Struct(payload); err != nil {
		util.WriteJsonError(w, "Validation failed", http.StatusBadRequest, err)
		return
	}

	// getting userInfo
	userCred, err := uh.UserRepository.GetUserByEmail(payload.Email)
	if err != nil {
		util.WriteJsonError(w, "couldnt fetch user", http.StatusBadRequest, err)
		return
	}

	//Check pass
	err = bcrypt.CompareHashAndPassword([]byte(userCred.Password), []byte(payload.Password))
	if err != nil {
		util.WriteJsonError(w, "wrong password", http.StatusBadRequest, err)
		return
	}

	exp := time.Now().Add(time.Hour * 24).Unix()
	token, err := Jwt.CreateJwt(uh.Key, userCred.ID, exp)
	if err != nil {
		util.WriteJsonError(w, "token creation faild", http.StatusInternalServerError, err)
		return
	}

	resp := model.UserLoginResponse{
		Message: "token created successfully.",
		Jwt:     token,
	}

	util.WriteJsonSuccess(w, http.StatusAccepted, resp)
}

func (uh *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userId, err := util.GetUserIDFromContext(r)
	if err != nil {
		util.WriteJsonError(w, "Not authorize", http.StatusInternalServerError, err)
		return
	}

	user, err := uh.UserRepository.GetProfileById(userId)
	if err != nil {
		util.WriteJsonError(w, "db error", http.StatusInternalServerError, err)
		return
	}

	resp := model.UserProfileResponse{
		UserInfo: model.UserInfo{
			UserName: user.UserName,
			Email:    user.Email,
		},
		Message: "user fetched successfully",
	}

	util.WriteJsonSuccess(w, http.StatusAccepted, resp)
}

func (uh *UserHandler) SendResetPasswordEmail(w http.ResponseWriter, r *http.Request) {
	var payload model.GetUserByEmailPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteJsonError(w, "couldnt parse payload", http.StatusBadRequest, err)
	}

	if err := uh.Validate.Struct(payload); err != nil {
		util.WriteJsonError(w, "Validation failed", http.StatusBadRequest, err)
	}

	userCred, err := uh.UserRepository.GetUserByEmail(payload.Email)
	if err != nil {
		util.WriteJsonError(w, err.Error(), http.StatusInternalServerError, err)
	}

	// make token
	token, err := Jwt.CreateJwt(string(constants.ResetPassTokenKey), userCred.ID, time.Now().Add(time.Hour*1).Unix())
	if err != nil {
		util.WriteJsonError(w, "token creation failed", http.StatusInternalServerError, err)
		return
	}

	// setup for mail
	data := model.ResetPasswordMailData{
		Email:    userCred.Email,
		ResetURL: constants.ResetPasswordURL + token,
		Year:     time.Now().Year(),
		Token:    token,
	}

	err = email.MailC.SendResetPasswordEmail(userCred.Email, "Change Password", data)
	if err != nil {
		util.WriteJsonError(w, "email send failed", http.StatusInternalServerError, err)
		return
	}

	util.WriteJsonSuccess(w, http.StatusAccepted, model.SuccessResponse{
		Message: "email sent successfully",
	})
}

func (uh *UserHandler) ChangePasswordByEmail(w http.ResponseWriter, r *http.Request) {
	var payload model.ChangePasswordByEmailPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteJsonError(w, "couldnt parse payload", http.StatusBadRequest, err)
	}

	if err := uh.Validate.Struct(payload); err != nil {
		util.WriteJsonError(w, "Validation failed", http.StatusBadRequest, err)
	}

	userId, err := util.GetUserIDFromContext(r)
	if err != nil {
		util.WriteJsonError(w, "Not authorize", http.StatusInternalServerError, err)
		return
	}

	err = uh.UserRepository.ChangePassword(userId, payload.Password)
	if err != nil {
		util.WriteJsonError(w, "db error", http.StatusInternalServerError, err)
		return
	}

	util.WriteJsonSuccess(w, http.StatusAccepted, model.SuccessResponse{
		Message: "password changed successfully",
	})
}
