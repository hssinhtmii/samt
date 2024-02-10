package api

import (
	"database/sql"
	"mime/multipart"
	"net/http"
	"os"
	db "samt/db/sqlc"
	"samt/unique"

	"github.com/gin-gonic/gin"
)

type uploadImageReq struct {
	Img       *multipart.FileHeader `json:"image"`
}

func (server *Server) UploadplanImage(ctx *gin.Context) {
	var req uploadImageReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, file_err := ctx.FormFile("image")

	if file_err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(file_err))
		return
	}

	path := "./images/" + unique.GenerateUniqueFileName() + ".jpg"

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.UploadImageParams{
		Img:      path,
		IsPoster: false,
	}
	upload_err := server.store.UploadImage(ctx, arg)

	if upload_err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(upload_err))
	}

	ctx.JSON(http.StatusOK, arg.Img)
}
func (server *Server) UploadPosterImage(ctx *gin.Context) {
	var req uploadImageReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, file_err := ctx.FormFile("image")

	if file_err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(file_err))
		return
	}

	path := "./images/" + unique.GenerateUniqueFileName() + ".jpg"

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.UploadImageParams{
		Img:      path,
		IsPoster: true,
	}
	upload_err := server.store.UploadImage(ctx, arg)

	if upload_err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(upload_err))
	}

	ctx.JSON(http.StatusOK, arg.Img)
}


type ImageReq struct {
	Image string `json:"image"`
}

func (server *Server) Deleteimage(ctx *gin.Context) {
	var req ImageReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	delete_err := os.Remove(req.Image)

	if delete_err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "دوباره تلاش کنید"})
	}

	err := server.store.DeleteImage(ctx, req.Image)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": "عکس با موفقیت حذف شد"})
}

type AllImageReq struct {
	Is_poster bool `json:"isposter"`
}

func (server *Server) GetAllimage(ctx *gin.Context) {
	var req AllImageReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	images, err := server.store.GetAllImage(ctx, req.Is_poster)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, images)
}
