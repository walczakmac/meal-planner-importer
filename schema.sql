create table macro
(
    id              smallserial primary key,
    meal_variant_id smallint not null,
    proteins        real     not null,
    fats            real     not null,
    carbs           real     not null,
    fiber           real     not null
);

create table meal
(
    id               smallserial primary key,
    meal_category_id smallint not null,
    "name"           varchar  not null,
    description      varchar  not null,
    "day"            smallint not null
);

create table meal_variant
(
    id         smallserial primary key,
    meal_id    smallint not null,
    kcal       real     not null,
    kcal_daily smallint not null,
    person     varchar  not null
);

create table meal_category
(
    id   smallserial primary key,
    name varchar not null
);