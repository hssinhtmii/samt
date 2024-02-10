// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: plan.sql

package db

import (
	"context"
	"database/sql"
)

const createPlan = `-- name: CreatePlan :exec
INSERT INTO plan (
        "plan_name",
        "year_of_plan_registration_date",
        "month_of_plan_registration_date",
        "day_of_plan_registration_date",
        "plan_capacity",
        "direct_employment",
        "application_of_the_product",
        "selling_price_of_products",
        "analysis_of_the_market_situation",
        "the_amount_of_domestic_production",
        "countrys_need",
        "nominal_capacity_of_existing_active_units",
        "the_nominal_capacity_of_projects_in_progress",
        "technicalknowledgecompany",
        "water",
        "fuel",
        "electricity",
        "land_area",
        "technical_knowledge",
        "type_of_equipment_required",
        "type_and_amount_of_major_raw_materials",
        "the_main_source_of_raw_materials",
        "foreign_exchange_capital",
        "rial_capital",
        "currency",
        "exchange_rate",
        "rial_working_capital",
        "currency_working_capital",
        "period",
        "total_capital",
        "annual_sales",
        "payback_time",
        "image"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21,
        $22,
        $23,
        $24,
        $25,
        $26,
        $27,
        $28,
        $29,
        $30,
        $31,
        $32,
        $33
    )
`

type CreatePlanParams struct {
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

func (q *Queries) CreatePlan(ctx context.Context, arg CreatePlanParams) error {
	_, err := q.db.ExecContext(ctx, createPlan,
		arg.PlanName,
		arg.YearOfPlanRegistrationDate,
		arg.MonthOfPlanRegistrationDate,
		arg.DayOfPlanRegistrationDate,
		arg.PlanCapacity,
		arg.DirectEmployment,
		arg.ApplicationOfTheProduct,
		arg.SellingPriceOfProducts,
		arg.AnalysisOfTheMarketSituation,
		arg.TheAmountOfDomesticProduction,
		arg.CountrysNeed,
		arg.NominalCapacityOfExistingActiveUnits,
		arg.TheNominalCapacityOfProjectsInProgress,
		arg.Technicalknowledgecompany,
		arg.Water,
		arg.Fuel,
		arg.Electricity,
		arg.LandArea,
		arg.TechnicalKnowledge,
		arg.TypeOfEquipmentRequired,
		arg.TypeAndAmountOfMajorRawMaterials,
		arg.TheMainSourceOfRawMaterials,
		arg.ForeignExchangeCapital,
		arg.RialCapital,
		arg.Currency,
		arg.ExchangeRate,
		arg.RialWorkingCapital,
		arg.CurrencyWorkingCapital,
		arg.Period,
		arg.TotalCapital,
		arg.AnnualSales,
		arg.PaybackTime,
		arg.Image,
	)
	return err
}

const deletePlan = `-- name: DeletePlan :exec
DELETE FROM plan
WHERE plan_name = $1
    AND year_of_plan_registration_date = $2
    AND month_of_plan_registration_date = $3
    AND day_of_plan_registration_date = $4
`

type DeletePlanParams struct {
	PlanName                    string `json:"plan_name"`
	YearOfPlanRegistrationDate  int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate int32  `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate   int32  `json:"day_of_plan_registration_date"`
}

func (q *Queries) DeletePlan(ctx context.Context, arg DeletePlanParams) error {
	_, err := q.db.ExecContext(ctx, deletePlan,
		arg.PlanName,
		arg.YearOfPlanRegistrationDate,
		arg.MonthOfPlanRegistrationDate,
		arg.DayOfPlanRegistrationDate,
	)
	return err
}

const getAllPlans = `-- name: GetAllPlans :many
SELECT plan_name, year_of_plan_registration_date, month_of_plan_registration_date, day_of_plan_registration_date, plan_capacity, direct_employment, application_of_the_product, selling_price_of_products, analysis_of_the_market_situation, the_amount_of_domestic_production, countrys_need, nominal_capacity_of_existing_active_units, the_nominal_capacity_of_projects_in_progress, technicalknowledgecompany, water, fuel, electricity, land_area, technical_knowledge, type_of_equipment_required, type_and_amount_of_major_raw_materials, the_main_source_of_raw_materials, foreign_exchange_capital, rial_capital, currency, exchange_rate, rial_working_capital, currency_working_capital, period, total_capital, annual_sales, payback_time, image
FROM "plan"
`

func (q *Queries) GetAllPlans(ctx context.Context) ([]Plan, error) {
	rows, err := q.db.QueryContext(ctx, getAllPlans)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Plan
	for rows.Next() {
		var i Plan
		if err := rows.Scan(
			&i.PlanName,
			&i.YearOfPlanRegistrationDate,
			&i.MonthOfPlanRegistrationDate,
			&i.DayOfPlanRegistrationDate,
			&i.PlanCapacity,
			&i.DirectEmployment,
			&i.ApplicationOfTheProduct,
			&i.SellingPriceOfProducts,
			&i.AnalysisOfTheMarketSituation,
			&i.TheAmountOfDomesticProduction,
			&i.CountrysNeed,
			&i.NominalCapacityOfExistingActiveUnits,
			&i.TheNominalCapacityOfProjectsInProgress,
			&i.Technicalknowledgecompany,
			&i.Water,
			&i.Fuel,
			&i.Electricity,
			&i.LandArea,
			&i.TechnicalKnowledge,
			&i.TypeOfEquipmentRequired,
			&i.TypeAndAmountOfMajorRawMaterials,
			&i.TheMainSourceOfRawMaterials,
			&i.ForeignExchangeCapital,
			&i.RialCapital,
			&i.Currency,
			&i.ExchangeRate,
			&i.RialWorkingCapital,
			&i.CurrencyWorkingCapital,
			&i.Period,
			&i.TotalCapital,
			&i.AnnualSales,
			&i.PaybackTime,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlan = `-- name: GetPlan :one
SELECT plan_name, year_of_plan_registration_date, month_of_plan_registration_date, day_of_plan_registration_date, plan_capacity, direct_employment, application_of_the_product, selling_price_of_products, analysis_of_the_market_situation, the_amount_of_domestic_production, countrys_need, nominal_capacity_of_existing_active_units, the_nominal_capacity_of_projects_in_progress, technicalknowledgecompany, water, fuel, electricity, land_area, technical_knowledge, type_of_equipment_required, type_and_amount_of_major_raw_materials, the_main_source_of_raw_materials, foreign_exchange_capital, rial_capital, currency, exchange_rate, rial_working_capital, currency_working_capital, period, total_capital, annual_sales, payback_time, image
FROM "plan"
WHERE plan_name = $1
    AND year_of_plan_registration_date = $2
    AND month_of_plan_registration_date = $3
    AND day_of_plan_registration_date = $4
`

type GetPlanParams struct {
	PlanName                    string `json:"plan_name"`
	YearOfPlanRegistrationDate  int32  `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate int32  `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate   int32  `json:"day_of_plan_registration_date"`
}

func (q *Queries) GetPlan(ctx context.Context, arg GetPlanParams) (Plan, error) {
	row := q.db.QueryRowContext(ctx, getPlan,
		arg.PlanName,
		arg.YearOfPlanRegistrationDate,
		arg.MonthOfPlanRegistrationDate,
		arg.DayOfPlanRegistrationDate,
	)
	var i Plan
	err := row.Scan(
		&i.PlanName,
		&i.YearOfPlanRegistrationDate,
		&i.MonthOfPlanRegistrationDate,
		&i.DayOfPlanRegistrationDate,
		&i.PlanCapacity,
		&i.DirectEmployment,
		&i.ApplicationOfTheProduct,
		&i.SellingPriceOfProducts,
		&i.AnalysisOfTheMarketSituation,
		&i.TheAmountOfDomesticProduction,
		&i.CountrysNeed,
		&i.NominalCapacityOfExistingActiveUnits,
		&i.TheNominalCapacityOfProjectsInProgress,
		&i.Technicalknowledgecompany,
		&i.Water,
		&i.Fuel,
		&i.Electricity,
		&i.LandArea,
		&i.TechnicalKnowledge,
		&i.TypeOfEquipmentRequired,
		&i.TypeAndAmountOfMajorRawMaterials,
		&i.TheMainSourceOfRawMaterials,
		&i.ForeignExchangeCapital,
		&i.RialCapital,
		&i.Currency,
		&i.ExchangeRate,
		&i.RialWorkingCapital,
		&i.CurrencyWorkingCapital,
		&i.Period,
		&i.TotalCapital,
		&i.AnnualSales,
		&i.PaybackTime,
		&i.Image,
	)
	return i, err
}

const updatePlan = `-- name: UpdatePlan :one
UPDATE plan
SET plan_name = COALESCE($5, plan_name),
    year_of_plan_registration_date = COALESCE(
        $6,
        year_of_plan_registration_date
    ),
    month_of_plan_registration_date = COALESCE(
        $7,
        month_of_plan_registration_date
    ),
    day_of_plan_registration_date = COALESCE(
        $8,
        day_of_plan_registration_date
    ),
    plan_capacity = COALESCE($9, plan_capacity),
    direct_employment = COALESCE(
        $10,
        direct_employment
    ),
    application_of_the_product = COALESCE(
        $11,
        application_of_the_product
    ),
    selling_price_of_products = COALESCE(
        $12,
        selling_price_of_products
    ),
    analysis_of_the_market_situation = COALESCE(
        $13,
        analysis_of_the_market_situation
    ),
    the_amount_of_domestic_production = COALESCE(
        $14,
        the_amount_of_domestic_production
    ),
    countrys_need = COALESCE($15, countrys_need),
    nominal_capacity_of_existing_active_units = COALESCE(
        $16,
        nominal_capacity_of_existing_active_units
    ),
    the_nominal_capacity_of_projects_in_progress = COALESCE(
        $17,
        the_nominal_capacity_of_projects_in_progress
    ),
    technicalknowledgecompany = COALESCE(
        $18,
        technicalknowledgecompany
    ),
    water = COALESCE($19, water),
    fuel = COALESCE($20, fuel),
    electricity = COALESCE($21, electricity),
    land_area = COALESCE($22, land_area),
    technical_knowledge = COALESCE(
        $23,
        technical_knowledge
    ),
    type_of_equipment_required = COALESCE(
        $24,
        type_of_equipment_required
    ),
    type_and_amount_of_major_raw_materials = COALESCE(
        $25,
        type_and_amount_of_major_raw_materials
    ),
    the_main_source_of_raw_materials = COALESCE(
        $26,
        The_main_source_of_raw_materials
    ),
    foreign_exchange_capital = COALESCE(
        $27,
        foreign_exchange_capital
    ),
    rial_capital = COALESCE($28, rial_capital),
    currency = COALESCE($29, currency),
    exchange_rate = COALESCE($30, exchange_rate),
    rial_working_capital = COALESCE(
        $31,
        rial_working_capital
    ),
    currency_working_capital = COALESCE(
        $32,
        currency_working_capital
    ),
    period = COALESCE($33, period),
    total_capital = COALESCE($34, total_capital),
    annual_sales = COALESCE($35, annual_sales),
    payback_time = COALESCE($36, payback_time),
    image = COALESCE($37, image)
WHERE plan_name = $1
    AND year_of_plan_registration_date = $2
    AND month_of_plan_registration_date = $3
    AND day_of_plan_registration_date = $4
RETURNING plan_name, year_of_plan_registration_date, month_of_plan_registration_date, day_of_plan_registration_date, plan_capacity, direct_employment, application_of_the_product, selling_price_of_products, analysis_of_the_market_situation, the_amount_of_domestic_production, countrys_need, nominal_capacity_of_existing_active_units, the_nominal_capacity_of_projects_in_progress, technicalknowledgecompany, water, fuel, electricity, land_area, technical_knowledge, type_of_equipment_required, type_and_amount_of_major_raw_materials, the_main_source_of_raw_materials, foreign_exchange_capital, rial_capital, currency, exchange_rate, rial_working_capital, currency_working_capital, period, total_capital, annual_sales, payback_time, image
`

type UpdatePlanParams struct {
	PlanName                               string         `json:"plan_name"`
	YearOfPlanRegistrationDate             int32          `json:"year_of_plan_registration_date"`
	MonthOfPlanRegistrationDate            int32          `json:"month_of_plan_registration_date"`
	DayOfPlanRegistrationDate              int32          `json:"day_of_plan_registration_date"`
	PlanName2                              string         `json:"plan_name2"`
	YearOfPlanRegistrationDate2            int32          `json:"year_of_plan_registration_date2"`
	MonthOfPlanRegistrationDate2           int32          `json:"month_of_plan_registration_date2"`
	DayOfPlanRegistrationDate2             int32          `json:"day_of_plan_registration_date2"`
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

func (q *Queries) UpdatePlan(ctx context.Context, arg UpdatePlanParams) (Plan, error) {
	row := q.db.QueryRowContext(ctx, updatePlan,
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
	)
	var i Plan
	err := row.Scan(
		&i.PlanName,
		&i.YearOfPlanRegistrationDate,
		&i.MonthOfPlanRegistrationDate,
		&i.DayOfPlanRegistrationDate,
		&i.PlanCapacity,
		&i.DirectEmployment,
		&i.ApplicationOfTheProduct,
		&i.SellingPriceOfProducts,
		&i.AnalysisOfTheMarketSituation,
		&i.TheAmountOfDomesticProduction,
		&i.CountrysNeed,
		&i.NominalCapacityOfExistingActiveUnits,
		&i.TheNominalCapacityOfProjectsInProgress,
		&i.Technicalknowledgecompany,
		&i.Water,
		&i.Fuel,
		&i.Electricity,
		&i.LandArea,
		&i.TechnicalKnowledge,
		&i.TypeOfEquipmentRequired,
		&i.TypeAndAmountOfMajorRawMaterials,
		&i.TheMainSourceOfRawMaterials,
		&i.ForeignExchangeCapital,
		&i.RialCapital,
		&i.Currency,
		&i.ExchangeRate,
		&i.RialWorkingCapital,
		&i.CurrencyWorkingCapital,
		&i.Period,
		&i.TotalCapital,
		&i.AnnualSales,
		&i.PaybackTime,
		&i.Image,
	)
	return i, err
}
