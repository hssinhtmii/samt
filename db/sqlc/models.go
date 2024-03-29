// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Image struct {
	Img      string `json:"img"`
	IsPoster bool   `json:"is_poster"`
}

type InvestorPlan struct {
	PhoneNumber                 string `json:"phone_number"`
	PlanName                    string `json:"plan_name"`
	YearOfPlanRegistrationDate  int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate int32  `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate   int32  `json:"day_of_plan_registration_date"`
}

type Plan struct {
	PlanName                               string         `json:"plan_name"`
	YearOfPlanRegistrationDate             int32          `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate            int32          `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate              int32          `json:"day_of_plan_registration_date"`
	PlanCapacity                           sql.NullString `json:"plan_capacity"`
	DirectEmployment                       sql.NullInt32  `json:"direct_employment"`
	ApplicationOfTheProduct                sql.NullString `json:"application_of_the_product"`
	SellingPriceOfProducts                 sql.NullString `json:"selling_price_of_products"`
	AnalysisOfTheMarketSituation           sql.NullString `json:"analysis_of_the_market_situation"`
	TheAmountOfDomesticProduction          sql.NullString `json:"the_amount_of_domestic_production"`
	CountrysNeed                           sql.NullString `json:"countrys_need"`
	NominalCapacityOfExistingActiveUnits   sql.NullString `json:"nominal_capacity_of_existing_active_units"`
	TheNominalCapacityOfProjectsInProgress sql.NullString `json:"the_nominal_capacity_of_projects_in_progress"`
	Technicalknowledgecompany              sql.NullString `json:"technicalknowledgecompany"`
	Water                                  sql.NullString `json:"water"`
	Fuel                                   sql.NullString `json:"fuel"`
	Electricity                            sql.NullString `json:"electricity"`
	LandArea                               sql.NullString `json:"land_area"`
	TechnicalKnowledge                     sql.NullString `json:"technical_knowledge"`
	TypeOfEquipmentRequired                sql.NullString `json:"type_of_equipment_required"`
	TypeAndAmountOfMajorRawMaterials       sql.NullString `json:"type_and_amount_of_major_raw_materials"`
	TheMainSourceOfRawMaterials            sql.NullString `json:"the_main_source_of_raw_materials"`
	ForeignExchangeCapital                 sql.NullString `json:"foreign_exchange_capital"`
	RialCapital                            sql.NullString `json:"rial_capital"`
	Currency                               sql.NullString `json:"currency"`
	ExchangeRate                           sql.NullInt32  `json:"exchange_rate"`
	RialWorkingCapital                     sql.NullString `json:"rial_working_capital"`
	CurrencyWorkingCapital                 sql.NullString `json:"currency_working_capital"`
	Period                                 sql.NullString `json:"period"`
	TotalCapital                           sql.NullString `json:"total_capital"`
	AnnualSales                            sql.NullString `json:"annual_sales"`
	PaybackTime                            sql.NullString `json:"payback_time"`
	Image                                  sql.NullString `json:"image"`
}

type Sessions struct {
	ID           uuid.UUID `json:"id"`
	PhoneNumber  string    `json:"phone_number"`
	RefreshToken string    `json:"refresh_token"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Uploader struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	PhoneNumber string         `json:"phone_number"`
	Email       sql.NullString `json:"email"`
	Password    string         `json:"password"`
}
