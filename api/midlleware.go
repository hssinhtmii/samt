package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"samt/token"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const (
	authorizationHK = "authorization"
	authorizationBearerType = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthorizationMiddleWare(TokenMaker token.Make_Token) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHK)
		if len(authorizationHeader) == 0 {
			err := errors.New("توکن تنظیم نشده است")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header has invalid format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		
		if authorizationType != authorizationBearerType {
			err := fmt.Errorf("authorization has invalid type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		access_token := fields[1]
		payload, err := TokenMaker.VerifyToken(access_token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return			
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
func Handler(router *gin.Engine) {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
		return 
	}

	serveraddress := os.Getenv("SERVERADDRESS")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(serveraddress, handler))
	http.Handle("/", router)
}