package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	vals := strings.Split(datastring, ",")
	if len(vals) != 3 {
		return fmt.Errorf("ошибка при разделении строки %s", datastring)
	}
	steps, err := strconv.Atoi(vals[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("invalid number of steps: %d (must be greater than 0)", steps)
	}
	t.Steps = steps
	trainingType := vals[1]
	t.TrainingType = trainingType
	duration, err := time.ParseDuration(vals[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("invalid duration: %d (must be greater than 0)", duration)
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	durationHours := t.Duration.Hours()

	var calculatedCalories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calculatedCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		calculatedCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки %s", t.TrainingType)
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		durationHours,
		distance,
		meanSpeed,
		calculatedCalories)
	return result, nil
}
