package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"meal-planner-importer/queries"
	"strconv"
	"strings"
)

var dayNameMap = map[string]int16{
	"poniedzialek": 1,
	"wtorek":       2,
	"środa":        3,
	"czwartek":     4,
	"piatek":       5,
	"sobota":       6,
	"niedziela":    7,
}

type MealService struct {
	queries queries.Queries
}

func NewMealService(db *sql.DB) *MealService {
	return &MealService{queries: *queries.New(db)}
}

func (service MealService) NewMeal(details []string) error {
	category, err := service.GetMealCategoryByName(details[9], details[0])
	if err != nil {
		return err
	}
	name := service.createNameFromMultiline(details[0])
	meal := service.GetMealByName(name)
	if meal.ID == 0 {
		meal, err = service.createMeal(name, category.ID, details[2], dayNameMap[details[10]])
	}
	if err != nil {
		return err
	}

	variant, err := service.createVariant(meal, details)
	if err != nil {
		return err
	}

	err = service.createMacro(variant.ID, meal.Name, details)
	if err != nil {
		return err
	}

	log.Printf("Added meal: %s\n", meal.Name)

	return nil
}

func (service MealService) GetMealCategoryByName(categoryName string, mealName string) (queries.MealCategory, error) {
	mealCategory, err := service.queries.GetMealCategory(ctx, strings.ToLower(categoryName))
	if err != nil {
		return mealCategory, errors.New(
			fmt.Sprintf(
				"Couldn't find meal category '%s' when creating meal '%s'. Error: %s\n",
				strings.ToLower(categoryName),
				mealName,
				err,
			),
		)
	}

	return mealCategory, nil
}

func (service MealService) GetMealByName(name string) queries.Meal {
	meal, err := service.queries.GetMeal(ctx, name)
	if err != nil {
		log.Printf("Error trying to find meal '%s', error: %s\n", name, err)
	}
	return meal
}

func (service MealService) createMacro(variantId int16, mealName string, rowValue []string) error {
	proteins, _ := strconv.ParseFloat(rowValue[4], 32)
	fats, _ := strconv.ParseFloat(rowValue[5], 32)
	carbs, _ := strconv.ParseFloat(rowValue[6], 32)
	fiber, _ := strconv.ParseFloat(rowValue[7], 32)
	_, err := service.queries.CreateMacro(ctx, queries.CreateMacroParams{
		MealVariantID: variantId,
		Proteins:      float32(proteins),
		Fats:          float32(fats),
		Carbs:         float32(carbs),
		Fiber:         float32(fiber),
	})
	if err != nil {
		return errors.New(fmt.Sprintf("couldn't create macro for meal %s\n", mealName))
	}
	return nil
}

func (service MealService) createVariant(meal queries.Meal, rowValue []string) (queries.MealVariant, error) {
	kcal, _ := strconv.ParseFloat(rowValue[3], 32)
	kcalDaily, _ := strconv.ParseInt(rowValue[12], 10, 16)
	variant, err := service.queries.CreateMealVariant(ctx, queries.CreateMealVariantParams{
		MealID:    meal.ID,
		Kcal:      float32(kcal),
		KcalDaily: int16(kcalDaily),
		Person:    rowValue[11],
	})
	if err != nil {
		return variant, errors.New(
			fmt.Sprintf(
				"Couldn't create meal variant for meal '%s'. Variant %d-%s. Error: %s\n",
				meal.Name,
				kcalDaily,
				rowValue[11],
				err,
			),
		)
	}
	return variant, nil
}

func (service MealService) createMeal(
	name string,
	mealCategoryId int16,
	description string,
	day int16,
) (queries.Meal, error) {
	meal, err := service.queries.CreateMeal(ctx, queries.CreateMealParams{
		MealCategoryID: mealCategoryId,
		Name:           name,
		Description:    description,
		Day:            day,
	})
	if err != nil {
		return meal, errors.New(
			fmt.Sprintf("Couldn't create meal '%s', error: %s\n", name, err),
		)
	}
	return meal, nil
}

func (service MealService) createNameFromMultiline(originalName string) string {
	nameParts := strings.Split(originalName, "\n")
	if len(nameParts) > 1 {
		nameParts[1] = strings.ToLower(nameParts[1])
	}
	name := strings.Join(nameParts, " ")
	return name
}
