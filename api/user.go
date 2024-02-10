package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"net/http"
	db "samt/db/sqlc"
	"samt/token"
	"samt/unique"
	"time"
)

// insert into user
type CreateUserReq struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,min=11,max=11"`
	Email       string `json:"email"`
	Password    string `json:"password" binding:"required"`
}

// show result
type UserRes struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

// inserting into the user table
func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// produce hash password
	hashPassword, err := unique.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email: sql.NullString{
			String: req.Email,
			Valid:  true,
		},
		Password: hashPassword,
	}

	err = server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": "create user was successful"})
}

// get all users
func (server *Server) getAllUsers(ctx *gin.Context) {
	user, err := server.store.GetAllUsers(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var res []UserRes
	for _, v := range user {
		req := UserRes{
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email.String,
		}
		res = append(res, req)
	}
	ctx.JSON(http.StatusOK, res)
}

// get one user from phone number
type getUserReq struct {
	PhoneNumber string `json:"phone_number" binding:"required,min=11,max=11"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	user, err := server.store.GetUser(ctx, req.PhoneNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := UserRes{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email.String,
	}
	ctx.JSON(http.StatusOK, res)
}

type DeleteUserReq struct {
	PhoneNumber string `json:"phone_number" binding:"required,min=11,max=11"`
}

func (server *Server) DeleteUser(ctx *gin.Context) {
	var req DeleteUserReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	_, err := server.store.GetUser(ctx, req.PhoneNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "کاربر وجود ندارد،لطفا ثبت نام کنید"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteUser(ctx, req.PhoneNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("user with phone_number = %s has was Deleted", req.PhoneNumber))

}

// login user
type LoginUserReq struct {
	PhoneNumber string `json:"phone_number" binding:"required,min=11,max=11"`
	Password    string `json:"password" binding:"required"`
}

type LoginUserRes struct {
	SessionsId               uuid.UUID `json:"sessions_id"`
	Access_token             string    `json:"access_token"`
	Access_token_expired_at  time.Time `json:"access_token_expired_at"`
	Refresh_token            string    `json:"refresh_token"`
	Refresh_token_expired_at time.Time `json:"refresh_token_expired_at"`
	User                     UserRes   `json:"user"`
}

// show user for with token
func user_response(user db.User) UserRes {
	return UserRes{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email.String,
	}
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.PhoneNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "کاربر وجود ندارد،لطفا ثبت نام کنید"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = unique.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "رمز عبور اشتباه است"})
		return
	}
	// make a token for user
	make_token_time := 365 * 24 * time.Hour 
	access_token, access_payload, err := server.tokenMaker.CreateToken(
		user.PhoneNumber,
		make_token_time,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// refresh token codes
	refresh_token_time := 365 * 24 * time.Hour
	refresh_token, refresh_payload, err := server.tokenMaker.CreateToken(
		user.PhoneNumber,
		refresh_token_time,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	sessions, err := server.store.CreateSessions(ctx, db.CreateSessionsParams{
		ID:           refresh_payload.Id,
		PhoneNumber:  user.PhoneNumber,
		RefreshToken: refresh_token,
		IsBlocked:    false,
		ExpiredAt:    refresh_payload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := user_response(user)
	rsp := LoginUserRes{
		SessionsId:               sessions.ID,
		Access_token:             access_token,
		Access_token_expired_at:  access_payload.ExpiredAt,
		Refresh_token:            refresh_token,
		Refresh_token_expired_at: refresh_payload.ExpiredAt,
		User:                     res,
	}
	ctx.JSON(http.StatusOK, rsp)
}

// updating user ** user.sql.go to be fixed
type UpdateUserRow struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (server *Server) UpdateUser(ctx *gin.Context) {
	var req UpdateUserRow

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	getuser, err := server.store.GetUser(ctx, req.PhoneNumber)

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.PhoneNumber != req.PhoneNumber {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر دسترسی ندارد"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.UpdateUserParams{
		PhoneNumber: req.PhoneNumber,
	}

	if req.FirstName != "" {
		arg.FirstName.String = req.FirstName
	} else {
		arg.FirstName.String = getuser.FirstName
	}
	if req.LastName != "" {
		arg.LastName.String = req.LastName
	} else {
		arg.LastName.String = getuser.LastName
	}
	if req.Email != "" {
		arg.Email.String = req.Email
	} else {
		arg.Email.String = getuser.Email.String
	}
	if req.Password != "" {
		arg.Password.String = req.Password
	} else {
		arg.Password.String = getuser.Password
	}

	user, err := server.store.UpdateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := UserRes{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email.String,
	}

	ctx.JSON(http.StatusOK, res)
}
