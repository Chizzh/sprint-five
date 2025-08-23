package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	vals := strings.Split(datastring, ",")
	if len(vals) != 2 {
		return fmt.Errorf("failed to split string %s", datastring)
	}
	steps, err := strconv.Atoi(vals[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("invalid number of steps: %d (must be greater than 0)", steps)
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(vals[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("invalid duration: %d (must be greater than 0)", duration)
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if ds.Steps <= 0 {
		return "", fmt.Errorf("invalid number of steps: %d (must be greater than 0)", ds.Steps)
	}
	if ds.Duration <= 0 {
		return "", fmt.Errorf("invalid duration: %d (must be greater than 0)", ds.Duration)
	}
	if ds.Weight <= 0 {
		return "", fmt.Errorf("invalid weight: %.2f (must be greater than 0)", ds.Weight)
	}
	if ds.Height <= 0 {
		return "", fmt.Errorf("invalid height: %.2f (cannot be negative)", ds.Height)
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories)
	return result, nil
}
