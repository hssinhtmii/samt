package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	db "samt/db/sqlc"
)

// insert plan into database
type PlanReq struct {
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

func (server *Server) CreatePlan(ctx *gin.Context) {
	var req PlanReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePlanParams{
		PlanName:                    req.PlanName,
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:   req.DayOfPlanRegistrationDate,
		PlanCapacity: sql.NullString{
			String: req.PlanCapacity,
			Valid:  true,
		},
		DirectEmployment: sql.NullInt32{
			Int32: req.DirectEmployment,
			Valid: true,
		},
		ApplicationOfTheProduct: sql.NullString{
			String: req.ApplicationOfTheProduct,
			Valid:  true,
		},
		SellingPriceOfProducts: sql.NullString{
			String: req.SellingPriceOfProducts,
			Valid:  true,
		},
		AnalysisOfTheMarketSituation: sql.NullString{
			String: req.AnalysisOfTheMarketSituation,
			Valid:  true,
		},
		TheAmountOfDomesticProduction: sql.NullString{
			String: req.TheAmountOfDomesticProduction,
			Valid:  true,
		},
		CountrysNeed: sql.NullString{
			String: req.CountrysNeed,
			Valid:  true,
		},

		NominalCapacityOfExistingActiveUnits: sql.NullString{
			String: req.NominalCapacityOfExistingActiveUnits,
			Valid:  true,
		},

		TheNominalCapacityOfProjectsInProgress: sql.NullString{
			String: req.TheNominalCapacityOfProjectsInProgress,
			Valid:  true,
		},

		Technicalknowledgecompany: sql.NullString{
			String: req.TechnicalKnowledge,
			Valid:  true,
		},
		Water: sql.NullString{
			String: req.Water,
			Valid:  true,
		},
		Fuel: sql.NullString{
			String: req.Fuel,
			Valid:  true,
		},
		Electricity: sql.NullString{
			String: req.Electricity,
			Valid:  true,
		},
		LandArea: sql.NullString{
			String: req.LandArea,
			Valid:  true,
		},
		TechnicalKnowledge: sql.NullString{
			String: req.TechnicalKnowledge,
			Valid:  true,
		},
		TypeOfEquipmentRequired: sql.NullString{
			String: req.TypeOfEquipmentRequired,
			Valid:  true,
		},
		TypeAndAmountOfMajorRawMaterials: sql.NullString{
			String: req.TypeAndAmountOfMajorRawMaterials,
			Valid:  true,
		},
		TheMainSourceOfRawMaterials: sql.NullString{
			String: req.TheMainSourceOfRawMaterials,
			Valid:  true,
		},
		ForeignExchangeCapital: sql.NullString{
			String: req.ForeignExchangeCapital,
			Valid:  true,
		},
		RialCapital: sql.NullString{
			String: req.RialCapital,
			Valid:  true,
		},
		Currency: sql.NullString{
			String: req.Currency,
			Valid:  true,
		},
		ExchangeRate: sql.NullInt32{
			Int32: req.ExchangeRate,
			Valid: true,
		},
		RialWorkingCapital: sql.NullString{
			String: req.RialCapital,
			Valid:  true,
		},
		CurrencyWorkingCapital: sql.NullString{
			String: req.CurrencyWorkingCapital,
			Valid:  true,
		},
		Period: sql.NullString{
			String: req.Period,
			Valid:  true,
		},
		TotalCapital: sql.NullString{
			String: req.TotalCapital,
			Valid:  true,
		},
		AnnualSales: sql.NullString{
			String: req.AnnualSales,
			Valid:  true,
		},
		PaybackTime: sql.NullString{
			String: req.PaybackTime,
			Valid:  true,
		},
		Image: sql.NullString{
			String: req.Image,
			Valid:  true,
		},
	}

	err := server.store.CreatePlan(ctx, arg)

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
	ctx.JSON(http.StatusOK, "create plan was successful")
}

// get all plans
func (server *Server) GetAllPlan(ctx *gin.Context) {
	var result []PlanReq
	c := make(chan PlanReq)

	plan, err := server.store.GetAllPlans(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, v := range plan {
		go convert_to_correct_output(v, c)
		result = append(result, <-c)
	}
	ctx.JSON(http.StatusOK, result)
}

type GetPlanReq struct {
	PlanName                    string `json:"plan_name"`
	YearOfPlanRegistrationDate  int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate int32  `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate   int32  `json:"day_of_plan_registration_date"`
}

func (server *Server) getPlan(ctx *gin.Context) {
	var req PlanReq
	c := make(chan PlanReq)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetPlanParams{
		PlanName:                    req.PlanName,
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:   req.DayOfPlanRegistrationDate,
	}
	plan, err := server.store.GetPlan(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	go convert_to_correct_output(plan, c)
	ctx.JSON(http.StatusOK, <-c)
}

type UpdatePlanReq struct {
	PlanName                               string `json:"plan_name"`
	YearOfPlanRegistrationDate             int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate            int32  `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate              int32  `json:"day_of_plan_registration_date"`
	PlanName2                              string `json:"plan_name2"`
	YearOfPlanRegistrationDate2            int32  `json:"year_of_plan_registration_date2"`
	MonthOfPlanRegistrationDate2           int32  `json:"month_of_plan_registration_date2"`
	DayOfPlanRegistrationDate2             int32  `json:"day_of_plan_registration_date2"`
	PlanCapacity                           string `json:"plan_capacity"`
	DirectEmployment                       int32  `json:"direct_employment"`
	ApplicationOfTheProduct                string `json:"application_of_the_product"`
	SellingPriceOfProducts                 string `json:"selling_price_of_products"`
	AnalysisOfTheMarketSituation           string `json:"analysis_of_the_market_situation"`
	TheAmountOfDomesticProduction          string `json:"the_amount_of_domestic_production"`
	CountrysNeed                           string `json:"countrys_need"`
	NominalCapacityOfExistingActiveUnits   string `json:"nominal_capacity_of_existing_active_units"`
	TheNominalCapacityOfProjectsInProgress string `json:"the_nominal_capacity_of_projects_in_progress"`
	Technicalknowledgecompany              string `json:"technicalknowledgecompany"`
	Water                                  string `json:"water"`
	Fuel                                   string `json:"fuel"`
	Electricity                            string `json:"electricity"`
	LandArea                               string `json:"land_area"`
	TechnicalKnowledge                     string `json:"technical_knowledge"`
	TypeOfEquipmentRequired                string `json:"type_of_equipment_required"`
	TypeAndAmountOfMajorRawMaterials       string `json:"type_and_amount_of_major_raw_materials"`
	TheMainSourceOfRawMaterials            string `json:"the_main_source_of_raw_materials"`
	ForeignExchangeCapital                 string `json:"foreign_exchange_capital"`
	RialCapital                            string `json:"rial_capital"`
	Currency                               string `json:"currency"`
	ExchangeRate                           int32  `json:"exchange_rate"`
	RialWorkingCapital                     string `json:"rial_working_capital"`
	CurrencyWorkingCapital                 string `json:"currency_working_capital"`
	Period                                 string `json:"period"`
	TotalCapital                           string `json:"total_capital"`
	AnnualSales                            string `json:"annual_sales"`
	PaybackTime                            string `json:"payback_time"`
	Image                                  string `json:"image"`
}

func (server *Server) UpdatePlan(ctx *gin.Context) {
	var req UpdatePlanReq
	c := make(chan PlanReq)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.GetPlanParams{
		PlanName:                    req.PlanName,
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:   req.DayOfPlanRegistrationDate,
	}
	getplan, err := server.store.GetPlan(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if req.PlanName2 == "" {
		req.PlanName2 = req.PlanName
	}
	if req.YearOfPlanRegistrationDate2 == 0 {
		req.YearOfPlanRegistrationDate2 = req.YearOfPlanRegistrationDate
	}
	if req.DayOfPlanRegistrationDate2 == 0 {
		req.DayOfPlanRegistrationDate2 = req.DayOfPlanRegistrationDate
	}
	if req.MonthOfPlanRegistrationDate2 == 0 {
		req.MonthOfPlanRegistrationDate2 = req.MonthOfPlanRegistrationDate
	}
	new_arg := db.UpdatePlanParams{
		PlanName:                     req.PlanName,
		YearOfPlanRegistrationDate:   req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate:  req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:    req.DayOfPlanRegistrationDate,
		PlanName2:                    req.PlanName2,
		YearOfPlanRegistrationDate2:  req.YearOfPlanRegistrationDate2,
		MonthOfPlanRegistrationDate2: req.MonthOfPlanRegistrationDate2,
		DayOfPlanRegistrationDate2:   req.DayOfPlanRegistrationDate2,
	}
	// checking for noy not
	new_arg.PlanCapacity.String = updateField(req.PlanCapacity, getplan.PlanCapacity.String)
	new_arg.DirectEmployment.Int32 = updateIntField(req.DirectEmployment, getplan.DirectEmployment.Int32)
	new_arg.ApplicationOfTheProduct.String = updateField(req.ApplicationOfTheProduct, getplan.ApplicationOfTheProduct.String)
	new_arg.SellingPriceOfProducts.String = updateField(req.SellingPriceOfProducts, getplan.SellingPriceOfProducts.String)
	new_arg.AnalysisOfTheMarketSituation.String = updateField(req.AnalysisOfTheMarketSituation, getplan.AnalysisOfTheMarketSituation.String)
	new_arg.TheAmountOfDomesticProduction.String = updateField(req.TheAmountOfDomesticProduction, getplan.TheAmountOfDomesticProduction.String)
	new_arg.CountrysNeed.String = updateField(req.CountrysNeed, getplan.CountrysNeed.String)
	new_arg.NominalCapacityOfExistingActiveUnits.String = updateField(req.NominalCapacityOfExistingActiveUnits, getplan.NominalCapacityOfExistingActiveUnits.String)
	new_arg.TheNominalCapacityOfProjectsInProgress.String = updateField(req.TheNominalCapacityOfProjectsInProgress, getplan.TheNominalCapacityOfProjectsInProgress.String)
	new_arg.Technicalknowledgecompany.String = updateField(req.Technicalknowledgecompany, getplan.Technicalknowledgecompany.String)
	new_arg.Water.String = updateField(req.Water, getplan.Water.String)
	new_arg.Fuel.String = updateField(req.Fuel, getplan.Water.String)
	new_arg.Electricity.String = updateField(req.Fuel, getplan.Electricity.String)
	new_arg.LandArea.String = updateField(req.LandArea, getplan.LandArea.String)
	new_arg.TechnicalKnowledge.String = updateField(req.TechnicalKnowledge, getplan.TechnicalKnowledge.String)
	new_arg.TypeAndAmountOfMajorRawMaterials.String = updateField(req.TypeAndAmountOfMajorRawMaterials, getplan.TypeAndAmountOfMajorRawMaterials.String)
	new_arg.TheMainSourceOfRawMaterials.String = updateField(req.TheMainSourceOfRawMaterials, getplan.TheMainSourceOfRawMaterials.String)
	new_arg.ForeignExchangeCapital.String = updateField(req.ForeignExchangeCapital, getplan.ForeignExchangeCapital.String)
	new_arg.RialCapital.String = updateField(req.RialCapital, getplan.RialCapital.String)
	new_arg.Currency.String = updateField(req.Currency, getplan.Currency.String)
	new_arg.ExchangeRate.Int32 = updateIntField(req.ExchangeRate, getplan.ExchangeRate.Int32)
	new_arg.RialWorkingCapital.String = updateField(req.RialWorkingCapital, getplan.RialWorkingCapital.String)
	new_arg.CurrencyWorkingCapital.String = updateField(req.CurrencyWorkingCapital, getplan.CurrencyWorkingCapital.String)
	new_arg.Period.String = updateField(req.Period, getplan.Period.String)
	new_arg.TotalCapital.String = updateField(req.TotalCapital, getplan.TotalCapital.String)
	new_arg.AnnualSales.String = updateField(req.AnnualSales, getplan.AnnualSales.String)
	new_arg.PaybackTime.String = updateField(req.PaybackTime, getplan.PaybackTime.String)
	new_arg.Image.String = updateField(req.Image, getplan.Image.String)

	up_plan, err := server.store.UpdatePlan(ctx, new_arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err: ": "after update"})
		return
	}
	go convert_to_correct_output(up_plan, c)
	ctx.JSON(http.StatusOK, <-c)
}

// delete plan function
type DeletePlanReq struct {
	PlanName                    string `json:"plan_name"`
	YearOfPlanRegistrationDate  int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate int32  `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate   int32  `json:"day_of_plan_registration_date"`
}

func (server *Server) deletePlan(ctx *gin.Context) {
	var req DeletePlanReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.DeletePlanParams{
		PlanName:                    req.PlanName,
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:   req.DayOfPlanRegistrationDate,
	}
	get_arg := db.GetPlanParams{
		PlanName:                    req.PlanName,
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
		DayOfPlanRegistrationDate:   req.DayOfPlanRegistrationDate,
	}

	_, err := server.store.GetPlan(ctx, get_arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "طرح وجود ندارد."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeletePlan(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			case "foreign_key_violation":
				server.DeleteInvestor_plan(ctx)
				ctx.JSON(http.StatusFailedDependency, "maybe later fix this bug")
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("plan with name = %s has was Deleted", req.PlanName))
}

// checking for not null
func updateField(newValue, defaultValue string) string {
	if newValue != "" {
		return newValue
	}
	return defaultValue
}

// checking for not null
func updateIntField(newValue, defaultValue int32) int32 {
	if newValue != 0 {
		return newValue
	}
	return defaultValue
}

// show output without "valid" value
func convert_to_correct_output(plan db.Plan, c chan PlanReq) {
	var res PlanReq
	res.PlanName = plan.PlanName
	res.YearOfPlanRegistrationDate = plan.YearOfPlanRegistrationDate
	res.MonthOfPlanRegistrationDate = plan.MonthOfPlanRegistrationDate
	res.DayOfPlanRegistrationDate = plan.DayOfPlanRegistrationDate
	res.PlanCapacity = plan.PlanCapacity.String
	res.DirectEmployment = plan.DirectEmployment.Int32
	res.ApplicationOfTheProduct = plan.ApplicationOfTheProduct.String
	res.SellingPriceOfProducts = plan.SellingPriceOfProducts.String
	res.AnalysisOfTheMarketSituation = plan.AnalysisOfTheMarketSituation.String
	res.TheAmountOfDomesticProduction = plan.TheAmountOfDomesticProduction.String
	res.CountrysNeed = plan.CountrysNeed.String
	res.NominalCapacityOfExistingActiveUnits = plan.NominalCapacityOfExistingActiveUnits.String
	res.TheNominalCapacityOfProjectsInProgress = plan.TheNominalCapacityOfProjectsInProgress.String
	res.Technicalknowledgecompany = plan.Technicalknowledgecompany.String
	res.Water = plan.Water.String
	res.Fuel = plan.Fuel.String
	res.Electricity = plan.Electricity.String
	res.LandArea = plan.LandArea.String
	res.TechnicalKnowledge = plan.TechnicalKnowledge.String
	res.TypeOfEquipmentRequired = plan.TypeOfEquipmentRequired.String
	res.TypeAndAmountOfMajorRawMaterials = plan.TypeAndAmountOfMajorRawMaterials.String
	res.TheMainSourceOfRawMaterials = plan.TheMainSourceOfRawMaterials.String
	res.ForeignExchangeCapital = plan.ForeignExchangeCapital.String
	res.RialCapital = plan.RialCapital.String
	res.Currency = plan.Currency.String
	res.ExchangeRate = plan.ExchangeRate.Int32
	res.RialWorkingCapital = plan.RialWorkingCapital.String
	res.CurrencyWorkingCapital = plan.CurrencyWorkingCapital.String
	res.Period = plan.Period.String
	res.TotalCapital = plan.TotalCapital.String
	res.AnnualSales = plan.AnnualSales.String
	res.PaybackTime = plan.PaybackTime.String
	res.Image = plan.Image.String

	c <- res
}

type SearchPlan struct {
	PlanName                    string `json:"plan_name"`
	YearOfPlanRegistrationDate  int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate int32  `json:"month_of_plan_registration_date"`
}

func (server *Server) searchPlan(ctx *gin.Context) {
	var req SearchPlan
	var result []PlanReq
	c := make(chan PlanReq)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.SearchPlanParams{
		YearOfPlanRegistrationDate:  req.YearOfPlanRegistrationDate,
		MonthOfPlanRegistrationDate: req.MonthOfPlanRegistrationDate,
	}
	if req.PlanName != "" {
		arg.PlanName = "%" + req.PlanName + "%"
	}

	plans, err := server.store.SearchPlan(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, v := range plans {
		go convert_to_correct_output(v, c)
		result = append(result, <-c)
	}
	ctx.JSON(http.StatusOK, result)

}

/*
	arg.PlanName,
	arg.YearOfPlanRegistrationDate,
	arg.MonthOfPlanRegistrationDate,
	arg.DayOfPlanRegistrationDate,
	arg.PlanName2,
	arg.YearOfPlanRegistrationDate2,
	arg.MonthOfPlanRegistrationDate2,
	arg.DayOfPlanRegistrationDate2,
	arg.PlanCapacity.String,
	arg.DirectEmployment.Int32,
	arg.ApplicationOfTheProduct.String,
	arg.SellingPriceOfProducts.String,
	arg.AnalysisOfTheMarketSituation.String,
	arg.TheAmountOfDomesticProduction.String,
	arg.CountrysNeed,
	arg.NominalCapacityOfExistingActiveUnits.String,
	arg.TheNominalCapacityOfProjectsInProgress.String,
	arg.Technicalknowledgecompany.String,
	arg.Water.String,
	arg.Fuel.String,
	arg.Electricity.String,
	arg.LandArea.String,
	arg.TechnicalKnowledge.String,
	arg.TypeOfEquipmentRequired.String,
	arg.TypeAndAmountOfMajorRawMaterials.String,
	arg.TheMainSourceOfRawMaterials.String,
	arg.ForeignExchangeCapital.String,
	arg.RialCapital.String,
	arg.Currency.String,
	arg.ExchangeRate.Int32,
	arg.RialWorkingCapital.String,
	arg.CurrencyWorkingCapital.String,
	arg.Period.String,
	arg.TotalCapital.String,
	arg.AnnualSales.String,
	arg.PaybackTime.String,
	arg.Image.String,
*/
