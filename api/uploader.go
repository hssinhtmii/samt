package api

import (
	"net/http"
	db "samt/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CreateAdminReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (server *Server) CreateAdmin(ctx *gin.Context) {
	var req CreateAdminReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateAdminParams{
		Username: req.Username,
		Password: req.Password,
	}

	err := server.store.CreateAdmin(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "create admin was successful")
}
