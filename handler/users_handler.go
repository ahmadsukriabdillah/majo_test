package handler

import (
	"majo_test/services"
	"majo_test/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	s services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{s: s}
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *UserHandler) SignIn(ctx *gin.Context) {
	var req UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.NewResponseError(ctx, err)
		return
	}
	user, err := s.s.DoLogin(req.Username, req.Password)
	if err != nil {
		utils.NewResponseError(ctx, err)
		return
	}

	utils.NewResponseOk(ctx, gin.H{
		"token": user,
	})
}
