-- name: CreateInvestorPlan :one
INSERT INTO investor_plan (
        phone_number,
        plan_name,
        year_of_plan_registration_date,
        month_of_plan_registration_date,
        day_of_plan_registration_date
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: DeleteInvestorPlan :exec
DELETE FROM investor_plan
WHERE phone_number = $1
    AND plan_name = $2
    AND year_of_plan_registration_date = $3
    AND month_of_plan_registration_date = $4
    AND day_of_plan_registration_date = $5;
-- name: GetAllInvestors :many
SELECT u."first_name",
    u."last_name",
    u."email",
    u."phone_number"
FROM "user" u
    INNER JOIN "investor_plan" ip ON u."phone_number" = ip."phone_number";
-- name: GetAllInvestorPlans :many
SELECT *
FROM "plan" p
    INNER JOIN "investor_plan" ip ON p."plan_name" = ip."plan_name"
    AND p."year_of_plan_registration_date" = ip."year_of_plan_registration_date"
    AND p."month_of_plan_registration_date" = ip."month_of_plan_registration_date"
    AND p."day_of_plan_registration_date" = ip."day_of_plan_registration_date"
WHERE ip."phone_number" = $1;
-- name: GetAllPlansInvestor :many
SELECT u."first_name",
    u."last_name",
    u."email",
    u."phone_number"
FROM "user" u
    INNER JOIN "investor_plan" ip ON ip."phone_number" = u."phone_number"
WHERE ip."plan_name" = $1
    AND ip."year_of_plan_registration_date" = $2
    AND ip."month_of_plan_registration_date" = $3
    AND ip."day_of_plan_registration_date" = $4;