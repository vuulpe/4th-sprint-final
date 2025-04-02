package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",") //Spliting the string into a slice of strings.
	if len(parts) != 3 {              //Checking the slice length is 3
		return 0, "", 0, errors.New("invalid format: eptected 3 components")
	}
	steps, err := strconv.Atoi(parts[0]) //converting steps into int
	if err != nil {
		return 0, "", 0, fmt.Errorf("error of step parsing") //error
	}
	duration, err := time.ParseDuration(parts[2]) //parsing duretion
	if err != nil {
		return 0, "", 0, fmt.Errorf("error of duration parsing") //error
	}
	activiti := parts[1]
	return steps, activiti, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	return float64(steps) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return dist / hours
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "ошибка парсинга данных"
	}
	distance := distance(steps)
	speed := meanSpeed(steps, duration)
	durationHours := duration.Hours()
	var calories float64
	switch strings.ToLower(activity) {
	case "бег":
		calories = RunningSpentCalories(steps, weight, duration)
	case "ходьба":
		calories = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "неизвестный тип тренировки"
	}
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity,
		durationHours,
		distance,
		speed,
		calories,
	)
}

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.

const (
	runningCaloriesMeanSpeedMultiplier = 18.0
	runningCaloriesMeanSpeedShift      = 1.2
)

func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	speed := meanSpeed(steps, duration)
	calories := ((runningCaloriesMeanSpeedMultiplier * speed) - runningCaloriesMeanSpeedShift) * weight

	if calories < 0 {
		return 0
	}
	return calories
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	speed := meanSpeed(steps, duration)
	spentCalories := (walkingCaloriesWeightMultiplier * weight) + (speed*speed/height)*float64(walkingSpeedHeightMultiplier)*float64(duration.Hours())*float64(minInH)
	return spentCalories
}
