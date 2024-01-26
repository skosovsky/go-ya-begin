package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // Средняя длина шага.
	mInKm     = 1000  // Количество метров в километре.
	minInH    = 60    // Количество минут в часе.
	kmhInMsec = 0.278 // Коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // Количество сантиметров в метре.
	msgTmpl   = "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"
)

// distance возвращает дистанцию (в километрах), которую преодолел пользователь за время тренировки.
//
//	action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
//	action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
//	duration float64 — длительность тренировки в часах.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	distance := distance(action)
	return distance / duration
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
//	lengthPool int — длина бассейна в метрах.
//	countPool int — сколько раз пользователь переплыл бассейн.
//	duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}

	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// message подготавливает и возвращает строку с информацией о тренировке.
//
//	trainingType string — вид тренировки (Бег, Ходьба, Плавание).
//	duration float64 — длительность тренировки в часах.
func message(trainingType string, duration, distance, speed, calories float64) string {
	return fmt.Sprintf(msgTmpl, trainingType, duration, distance, speed, calories)
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
//	action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
//	trainingType string — вид тренировки (Бег, Ходьба, Плавание).
//	duration float64 — длительность тренировки в часах.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	switch trainingType {
	case "Бег":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := RunningSpentCalories(action, weight, duration)
		return message(trainingType, duration, distance, speed, calories)
	case "Ходьба":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height)
		return message(trainingType, duration, distance, speed, calories)
	case "Плавание":
		distance := distance(action)
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return message(trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18   // Множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79 // Среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных калорий при беге.
//
//	action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
//	weight float64 — вес пользователя.
//	duration float64 — длительность тренировки в часах.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	meanSpeedKmH := meanSpeed(action, duration)
	spentCalories := (runningCaloriesMeanSpeedMultiplier * meanSpeedKmH * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH

	return spentCalories
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // Множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // Множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
//	action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
//	duration float64 — длительность тренировки в часах.
//	weight float64 — вес пользователя.
//	height float64 — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	meanSpeedMS := meanSpeed(action, duration) * kmhInMsec
	spentCalories := (walkingCaloriesWeightMultiplier*weight + (math.Pow(meanSpeedMS, 2)/height*cmInM)*walkingSpeedHeightMultiplier*weight) * duration * minInH

	return spentCalories
}

// Константы для расчета калорий, расходуемых при плавании.
const (
	swimmingCaloriesMeanSpeedShift   = 1.1 // Среднее количество сжигаемых калорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier = 2   // Множитель веса при плавании.
)

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
//	lengthPool int — длина бассейна в метрах.
//	countPool int — сколько раз пользователь переплыл бассейн.
//	duration float64 — длительность тренировки в часах.
//	weight float64 — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	meanSpeedKmH := swimmingMeanSpeed(lengthPool, countPool, duration)
	spentCalories := (meanSpeedKmH + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration

	return spentCalories
}
