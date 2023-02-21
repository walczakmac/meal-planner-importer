
-- name: CreateMacro :one
insert into macro(meal_variant_id, proteins, fats, carbs, fiber) values ($1, $2, $3, $4, $5) returning *;

-- name: CreateMeal :one
insert into meal(meal_category_id, name, description, day) values($1, $2, $3, $4) returning *;

-- name: CreateMealVariant :one
insert into meal_variant(meal_id, kcal, kcal_daily, person) values ($1, $2, $3, $4) returning *;

-- name: CreateIngredient :one
insert into ingredient(meal_variant_id, product_id, amount, unit, snack) values ($1, $2, $3, $4, $5) returning *;

-- name: GetMealCategory :one
select * from meal_category where name = $1;

-- name: GetMeal :one
select * from meal where name = $1;

-- name: GetProduct :one
select * from product where name = $1;