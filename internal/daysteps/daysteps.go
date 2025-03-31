package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"errors"
	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",") //Spliting the string into a slice of strings.
	if len(parts) != 2 { //Checking the slice length is 2
		return 0, 0, errors.New("invalid format: eptected 'steps,duration'")
	}

	steps, err := strconv.Atoi(parts[[0]]) //converting steps number into the int type
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format: %v", err) //error if wrong speps format
	}
	if steps <= 0 {
		return 0, 0, errors.New("steps must be positive") //error if steps are not positive
	}

	if duration <= 0 {
		return 0, 0, errors.New("duration must be positive") //error if duration are not positive
	}

	return steps, duration, nill
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data) //Spliting the string into a slice of strings.
	if err != nil {
		fmt.Println("Error:", err)//printing errors
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceMeters := float64(steps) * StepLength
	distanceKMeters := distanceMeters / 100
	calories := WalkingSpentCalories(steps, duration, weight, height)
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKm, calories)
	return result 
}
