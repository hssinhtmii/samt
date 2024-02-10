package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "samt/db/sqlc"
	"samt/token"
)

const TokenSymetrickey = "01234567890123456789012345678901"

type Server struct {
	store      *db.Store
	tokenMaker token.Make_Token
	router     *gin.Engine
}

func NewServer(store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(TokenSymetrickey)

	if err != nil {
		return nil, fmt.Errorf("توکن به درستی تولید نشد = %w", err)
	}
	
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	
	server.setupRouter()
	return server, nil
}
func (server *Server) setupRouter() {
	router := gin.Default()
	//user calls
	router.POST("/SignUp", server.CreateUser)
	router.POST("/GetUser", server.getUser)
	router.POST("/Login", server.LoginUser)
	router.DELETE("/DeleteUser", server.DeleteUser)
	router.GET("/AllUser", server.getAllUsers)
	router.PUT("/UpdateUser", server.UpdateUser)
	
	//plan calls
	router.POST("/getPlan", server.getPlan)
	router.POST("/plan", server.CreatePlan)
	router.PUT("/UpdatePlan", server.UpdatePlan)
	router.DELETE("/DeletePlan", server.deletePlan)
	router.GET("/allPlans", server.GetAllPlan)
	
	// investor_plan calls
	router.DELETE("/DeleteInvestorPlan", server.DeleteInvestor_plan)
	router.GET("/GetAllInvestor", server.GetInvestor)
	router.GET("/GetAllPlan_Investors", server.GetAllPlan_Investors)
	router.POST("/GetInvestor_Plans", server.GetInvestor_Plans)
	router.POST("/investor_plan", server.Create_Investor_plan)
	
	// upload calls
	router.POST("/CreateAdmin", server.CreateAdmin)
	
	// image calls
	router.POST("/poster/image", server.UploadPosterImage)
	router.POST("/plan/image", server.UploadplanImage)
	router.DELETE("/Deleteimage", server.Deleteimage)
	router.GET("/GetAllimage", server.GetAllimage)
	
	// token cal
	router.POST("/tokens/renew", server.renewAccess_token)
	
	server.router = router
	// calling CORS
	Handler(router)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
