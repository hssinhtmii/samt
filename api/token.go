package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type renewAccess_tokenReq struct {
	Refresh_token string `json:"refresh_token"`
}

type renewAccess_tokenRes struct {
	Access_token            string    `json:"access_token"`
	Access_token_expired_at time.Time `json:"access_token_expired_at"`
}

func (server *Server) renewAccess_token(ctx *gin.Context) {
	var req renewAccess_tokenReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refresh_payload, err := server.tokenMaker.VerifyToken(req.Refresh_token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.GetSessions(ctx, refresh_payload.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if session.IsBlocked {
		err := fmt.Errorf("session is blocked")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	
	if session.PhoneNumber != refresh_payload.PhoneNumber {
		err := fmt.Errorf("کاربر مطابقت ندارد")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	if session.RefreshToken != req.Refresh_token {
		err := fmt.Errorf("توکن مطابقت ندارد یا اشتباه است")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	if time.Now().Local().After(session.ExpiredAt) {
		err := fmt.Errorf("توکن منقضی شده است")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return		
	}

	// make a token for user
	make_token_time := 365 * 24 * time.Hour
	access_token, access_payload, err := server.tokenMaker.CreateToken(
		refresh_payload.PhoneNumber,
		make_token_time,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := renewAccess_tokenRes{
		Access_token:             access_token,
		Access_token_expired_at:  access_payload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
