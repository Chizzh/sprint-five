package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("invalid number of steps: %d (must be greater than 0)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight: %.2f (must be greater than 0)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid height: %.2f (must be greater than 0)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration: %d (must be greater than 0)", duration)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("invalid number of steps: %d (must be greater than 0)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight: %.2f (must be greater than 0)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid height: %.2f (must be greater than 0)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration: %d (must be greater than 0)", duration)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	durationHours := duration.Hours()
	meanSpeed := distance / durationHours
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distance := (float64(steps) * stepLength) / float64(mInKm)
	return distance
}
