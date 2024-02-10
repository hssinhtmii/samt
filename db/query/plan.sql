-- name: CreatePlan :exec
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
    );
-- name: GetAllPlans :many
SELECT *
FROM "plan";
-- name: GetPlan :one
SELECT *
FROM "plan"
WHERE plan_name = $1
    AND year_of_plan_registration_date = $2
    AND month_of_plan_registration_date = $3
    AND day_of_plan_registration_date = $4;
-- name: UpdatePlan :one
UPDATE plan
SET plan_name = COALESCE(sqlc.arg(plan_name2), plan_name),
    year_of_plan_registration_date = COALESCE(
        sqlc.arg(year_of_plan_registration_date2),
        year_of_plan_registration_date
    ),
    month_of_plan_registration_date = COALESCE(
        sqlc.arg(month_of_plan_registration_date2),
        month_of_plan_registration_date
    ),
    day_of_plan_registration_date = COALESCE(
        sqlc.arg(day_of_plan_registration_date2),
        day_of_plan_registration_date
    ),
    plan_capacity = COALESCE(sqlc.narg(plan_capacity), plan_capacity),
    direct_employment = COALESCE(
        sqlc.narg(direct_employment),
        direct_employment
    ),
    application_of_the_product = COALESCE(
        sqlc.narg(application_of_the_product),
        application_of_the_product
    ),
    selling_price_of_products = COALESCE(
        sqlc.narg(selling_price_of_products),
        selling_price_of_products
    ),
    analysis_of_the_market_situation = COALESCE(
        sqlc.narg(analysis_of_the_market_situation),
        analysis_of_the_market_situation
    ),
    the_amount_of_domestic_production = COALESCE(
        sqlc.narg(the_amount_of_domestic_production),
        the_amount_of_domestic_production
    ),
    countrys_need = COALESCE(sqlc.narg(countrys_need), countrys_need),
    nominal_capacity_of_existing_active_units = COALESCE(
        sqlc.narg(nominal_capacity_of_existing_active_units),
        nominal_capacity_of_existing_active_units
    ),
    the_nominal_capacity_of_projects_in_progress = COALESCE(
        sqlc.narg(the_nominal_capacity_of_projects_in_progress),
        the_nominal_capacity_of_projects_in_progress
    ),
    technicalknowledgecompany = COALESCE(
        sqlc.narg(technicalknowledgecompany),
        technicalknowledgecompany
    ),
    water = COALESCE(sqlc.narg(water), water),
    fuel = COALESCE(sqlc.narg(fuel), fuel),
    electricity = COALESCE(sqlc.narg(electricity), electricity),
    land_area = COALESCE(sqlc.narg(land_area), land_area),
    technical_knowledge = COALESCE(
        sqlc.narg(technical_knowledge),
        technical_knowledge
    ),
    type_of_equipment_required = COALESCE(
        sqlc.narg(type_of_equipment_required),
        type_of_equipment_required
    ),
    type_and_amount_of_major_raw_materials = COALESCE(
        sqlc.narg(type_and_amount_of_major_raw_materials),
        type_and_amount_of_major_raw_materials
    ),
    the_main_source_of_raw_materials = COALESCE(
        sqlc.narg(The_main_source_of_raw_materials),
        The_main_source_of_raw_materials
    ),
    foreign_exchange_capital = COALESCE(
        sqlc.narg(foreign_exchange_capital),
        foreign_exchange_capital
    ),
    rial_capital = COALESCE(sqlc.narg(rial_capital), rial_capital),
    currency = COALESCE(sqlc.narg(currency), currency),
    exchange_rate = COALESCE(sqlc.narg(exchange_rate), exchange_rate),
    rial_working_capital = COALESCE(
        sqlc.narg(rial_working_capital),
        rial_working_capital
    ),
    currency_working_capital = COALESCE(
        sqlc.narg(currency_working_capital),
        currency_working_capital
    ),
    period = COALESCE(sqlc.narg(period), period),
    total_capital = COALESCE(sqlc.narg(total_capital), total_capital),
    annual_sales = COALESCE(sqlc.narg(annual_sales), annual_sales),
    payback_time = COALESCE(sqlc.narg(payback_time), payback_time),
    image = COALESCE(sqlc.narg(image), image)
WHERE plan_name = $1
    AND year_of_plan_registration_date = $2
    AND month_of_plan_registration_date = $3
    AND day_of_plan_registration_date = $4
RETURNING *;
-- name: DeletePlan :exec
DELETE FROM plan
WHERE plan_name = $1
    AND year_of_plan_registration_date = $2
    AND month_of_plan_registration_date = $3
    AND day_of_plan_registration_date = $4;
-- name: SearchPlan :many
SELECT *
FROM plan
WHERE plan_name LIKE $1
    OR year_of_plan_registration_date = $2
    OR month_of_plan_registration_date = $3;
    -- OR year_of_plan_registration_date > $4
    -- OR month_of_plan_registration_date > $5;