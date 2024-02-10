	CREATE TABLE "user" (
		"first_name" varchar(200) NOT NULL,
		"last_name" varchar(200) NOT NULL,
		"phone_number" varchar(11) PRIMARY KEY,
		"email" varchar DEFAULT NULL,
		"password" varchar NOT NULL
	  );
	  CREATE TABLE "investor_plan" (
		"phone_number" varchar(11),
		"plan_name" varchar,
		"year_of_plan_registration_date" int,
		"month_of_plan_registration_date" int,
		"day_of_plan_registration_date" int,
		PRIMARY KEY (
		  "phone_number",
		  "plan_name",
		  "year_of_plan_registration_date",
		  "month_of_plan_registration_date",
		  "day_of_plan_registration_date"
		)
	  );
	  CREATE TABLE "uploader" (
		"username" varchar PRIMARY KEY,
		"password" varchar NOT NULL
	  );
	  CREATE TABLE "image" ("img" varchar NOT NULL UNIQUE, is_poster boolean  NOT NULL default FALSE);
	  CREATE TABLE "plan" (
		"plan_name" varchar NOT NULL,
		"year_of_plan_registration_date" int NOT NULL,
		"month_of_plan_registration_date" int NOT NULL,
		"day_of_plan_registration_date" int NOT NULL,
		"plan_capacity" varchar,
		"direct_employment" int,
		"application_of_the_product" varchar,
		"selling_price_of_products" varchar,
		"analysis_of_the_market_situation" varchar,
		"the_amount_of_domestic_production" varchar,
		"countrys_need" varchar,
		"nominal_capacity_of_existing_active_units" varchar,
		"the_nominal_capacity_of_projects_in_progress" varchar,
		"technicalknowledgecompany" varchar,
		"water" varchar,
		"fuel" varchar,
		"electricity" varchar,
		"land_area" varchar,
		"technical_knowledge" varchar,
		"type_of_equipment_required" varchar,
		"type_and_amount_of_major_raw_materials" varchar,
		"the_main_source_of_raw_materials" varchar,
		"foreign_exchange_capital" varchar,
		"rial_capital" varchar,
		"currency" varchar,
		"exchange_rate" int,
		"rial_working_capital" varchar,
		"currency_working_capital" varchar,
		"period" varchar,
		"total_capital" varchar,
		"annual_sales" varchar,
		"payback_time" varchar,
		"image" varchar,
		PRIMARY KEY (
		  "plan_name",
		  "year_of_plan_registration_date",
		  "month_of_plan_registration_date",
		  "day_of_plan_registration_date"
		)
	  );
	  CREATE TABLE "sessions" (
		"id" uuid PRIMARY KEY,
		"phone_number" VARCHAR(11) NOT NULL,
		"refresh_token" VARCHAR NOT NULL,
		"is_blocked" BOOLEAN NOT NULL DEFAULT FALSE,
		"expired_at" TIMESTAMPTZ NOT NULL,
		"created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
	  );
	  ALTER TABLE "sessions"
	  ADD FOREIGN KEY ("phone_number") REFERENCES "user" ("phone_number") ON DELETE CASCADE;  
	  ALTER TABLE "sessions"
	  ADD FOREIGN KEY ("phone_number") REFERENCES "user" ("phone_number") ON DELETE CASCADE ON UPDATE CASCADE;
	  CREATE INDEX ON "plan" ("plan_name");
	  CREATE INDEX ON "plan" ("year_of_plan_registration_date");
	  CREATE INDEX ON "plan" ("month_of_plan_registration_date");
	  CREATE INDEX ON "plan" ("day_of_plan_registration_date");
	  CREATE INDEX ON "user" ("phone_number");
	  ALTER TABLE "investor_plan"
	  ADD FOREIGN KEY ("phone_number") REFERENCES "user" ("phone_number") ON DELETE CASCADE ON UPDATE CASCADE;
	  ALTER TABLE "investor_plan"
	  ADD FOREIGN KEY (
		  "plan_name",
		  "year_of_plan_registration_date",
		  "month_of_plan_registration_date",
		  "day_of_plan_registration_date"
		) REFERENCES "plan" (
		  "plan_name",
		  "year_of_plan_registration_date",
		  "month_of_plan_registration_date",
		  "day_of_plan_registration_date"
		) ON DELETE CASCADE ON UPDATE CASCADE;