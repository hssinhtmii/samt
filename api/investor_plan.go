package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	db "samt/db/sqlc"
)

// create investor_plan
type InvestorPlanReq struct {
	PhoneNumber                 string `json:"phone_number"`
	PlanName                    string `json:"plan_name"`
	YearOfPlanRegistrationDate  int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate int32  `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate   int32  `json:"day_of_plan_registration_date"`
}

func (server *Server) Create_Investor_plan(ctx *gin.Context) {
	var req InvestorPlanReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateInvestorPlanParams{
		PhoneNumber:                 req.PhoneNumber,
		PlanName:                    req.PlanName,
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:   req.DayOfPlanRegistrationDate,
	}
	inv_plan, err := server.store.CreateInvestorPlan(ctx, arg)

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
	ctx.JSON(http.StatusOK, inv_plan)
}

// delete investor_plan

func (server *Server) DeleteInvestor_plan(ctx *gin.Context) {
	var req InvestorPlanReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteInvestorPlanParams{
		PhoneNumber:                 req.PhoneNumber,
		PlanName:                    req.PlanName,
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:   req.DayOfPlanRegistrationDate,
	}

	err := server.store.DeleteInvestorPlan(ctx, arg)

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
	ctx.JSON(http.StatusOK, fmt.Sprintf("user with phone_number = %s has was Deleted", req.PhoneNumber))
}

// get all investors
type GetInvestorsRow struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (server *Server) GetInvestor(ctx *gin.Context) {
	var res []GetInvestorsRow
	c := make(chan GetInvestorsRow)

	GetInvestor, err := server.store.GetAllInvestors(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, v := range GetInvestor {
		go converttocorrectoutput(v, c)
		res = append(res, <-c)
	}

	ctx.JSON(http.StatusOK, res)
}

// ? get all investor with Specified phone_number plans
type Investor_PlanReq struct {
	PhoneNumber string `form:"phone_number" json:"phone_number"`
}

type InvestorPlanRes struct {
	PlanName                               string `json:"plan_name" binding:"required"`
	YearOfPlanRegistrationDate             int32  `json:"year_of_plan_registration_date" binding:"required"`
	MonthOfPlanRegistrationDate            int32  `json:"month_of_plan_registration_date" binding:"required"`
	DayOfPlanRegistrationDate              int32  `json:"day_of_plan_registration_date" binding:"required"`
	PlanCapacity                           string `json:"plan_capacity"`
	DirectEmployment                       int32  `json:"Direct_employment"`
	ApplicationOfTheProduct                string `json:"Application_of_the_product"`
	SellingPriceOfProducts                 string `json:"Selling_price_of_products"`
	AnalysisOfTheMarketSituation           string `json:"Analysis_of_the_market_situation"`
	TheAmountOfDomesticProduction          string `json:"The_amount_of_domestic_production"`
	CountrysNeed                           string `json:"countrys_need"`
	NominalCapacityOfExistingActiveUnits   string `json:"nominal_capacity_of_existing_active_units"`
	TheNominalCapacityOfProjectsInProgress string `json:"the_nominal_capacity_of_projects_in_progress"`
	Technicalknowledgecompany              string `json:"technicalknowledgecompany"`
	Water                                  string `json:"water"`
	Fuel                                   string `json:"fuel"`
	Electricity                            string `json:"electricity"`
	LandArea                               string `json:"land_area"`
	TechnicalKnowledge                     string `json:"Technical_knowledge"`
	TypeOfEquipmentRequired                string `json:"Type_of_equipment_required"`
	TypeAndAmountOfMajorRawMaterials       string `json:"Type_and_amount_of_major_raw_materials"`
	TheMainSourceOfRawMaterials            string `json:"The_main_source_of_raw_materials"`
	ForeignExchangeCapital                 string `json:"Foreign_exchange_capital"`
	RialCapital                            string `json:"Rial_capital"`
	Currency                               string `json:"currency"`
	ExchangeRate                           int32  `json:"exchange_rate"`
	RialWorkingCapital                     string `json:"rial_working_capital"`
	CurrencyWorkingCapital                 string `json:"currency_working_capital"`
	Period                                 string `json:"period"`
	TotalCapital                           string `json:"total_capital"`
	AnnualSales                            string `json:"Annual_sales"`
	PaybackTime                            string `json:"Payback_time"`
	Image                                  string `json:"image"`
}

func (server *Server) GetInvestor_Plans(ctx *gin.Context) {
	var req Investor_PlanReq
	var res []InvestorPlanRes
	c := make(chan InvestorPlanRes)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	PlanInvestor, err := server.store.GetAllInvestorPlans(ctx, req.PhoneNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, v := range PlanInvestor {
		go ConvertTOCorrectOutput(v, c)
		res = append(res, <-c)
	}
	ctx.JSON(http.StatusOK, res)

}

type AllPlansInvestorsReq struct {
	PlanName                    string `json:"plan_name"`
	YearOfPlanRegistrationDate  int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate int32  `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate   int32  `json:"day_of_plan_registration_date"`
}

func (server *Server) GetAllPlan_Investors(ctx *gin.Context) {
	var req AllPlansInvestorsReq
	var res []GetInvestorsRow
	c := make(chan GetInvestorsRow)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllPlansInvestorParams{
		PlanName:                    req.PlanName,
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:   req.DayOfPlanRegistrationDate,
	}

	investor_inf, err := server.store.GetAllPlansInvestor(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, v := range investor_inf {
		go converttocorrectoutput((db.GetAllInvestorsRow)(v), c)
		res = append(res, <-c)
	}

	ctx.JSON(http.StatusOK, res)
}

func ConvertTOCorrectOutput(req db.GetAllInvestorPlansRow, c chan InvestorPlanRes) {
	var res InvestorPlanRes
	res.PlanName = req.PlanName
	res.YearOfPlanRegistrationDate = req.YearOfPlanRegistrationDate
	res.MonthOfPlanRegistrationDate = req.MonthOfPlanRegistrationDate
	res.DayOfPlanRegistrationDate = req.DayOfPlanRegistrationDate
	res.PlanCapacity = req.PlanCapacity.String
	res.DirectEmployment = req.DirectEmployment.Int32
	res.ApplicationOfTheProduct = req.ApplicationOfTheProduct.String
	res.SellingPriceOfProducts = req.SellingPriceOfProducts.String
	res.AnalysisOfTheMarketSituation = req.AnalysisOfTheMarketSituation.String
	res.TheAmountOfDomesticProduction = req.TheAmountOfDomesticProduction.String
	res.CountrysNeed = req.CountrysNeed.String
	res.NominalCapacityOfExistingActiveUnits = req.NominalCapacityOfExistingActiveUnits.String
	res.TheNominalCapacityOfProjectsInProgress = req.TheNominalCapacityOfProjectsInProgress.String
	res.Technicalknowledgecompany = req.TechnicalKnowledge.String
	res.Water = req.Water.String
	res.Fuel = req.Fuel.String
	res.Electricity = req.Electricity.String
	res.LandArea = req.LandArea.String
	res.TechnicalKnowledge = req.TechnicalKnowledge.String
	res.TypeOfEquipmentRequired = req.TypeOfEquipmentRequired.String
	res.TypeAndAmountOfMajorRawMaterials = req.TypeAndAmountOfMajorRawMaterials.String
	res.TheMainSourceOfRawMaterials = req.TheMainSourceOfRawMaterials.String
	res.ForeignExchangeCapital = req.ForeignExchangeCapital.String
	res.RialCapital = req.RialCapital.String
	res.Currency = req.Currency.String
	res.ExchangeRate = req.ExchangeRate.Int32
	res.RialWorkingCapital = req.RialWorkingCapital.String
	res.CurrencyWorkingCapital = req.CurrencyWorkingCapital.String
	res.Period = req.Period.String
	res.TotalCapital = req.TotalCapital.String
	res.AnnualSales = req.AnnualSales.String
	res.PaybackTime = req.PaybackTime.String
	res.Image = req.Image.String
	
	c <- res
}

// ? convert to correct shape
func converttocorrectoutput(req db.GetAllInvestorsRow, c chan GetInvestorsRow) {
	var res GetInvestorsRow
	res.FirstName = req.FirstName
	res.LastName = req.LastName
	res.PhoneNumber = req.PhoneNumber
	res.Email = req.Email.String
	c <- res
}
