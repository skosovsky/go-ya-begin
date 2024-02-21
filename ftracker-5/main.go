package main

import (
	"fmt"
	"math"
	"time"
)

// Общие константы для вычислений.
const (
	MInKm      = 1000 // количество метров в одном километре
	MinInHours = 60   // количество минут в одном часе
	LenStep    = 0.65 // длина одного шага
	CmInM      = 100  // количество сантиметров в одном метре
)

// Training общая структура для всех тренировок.
type Training struct {
	TrainingType string        // тип тренировки
	Action       int           // количество повторов(шаги, гребки при плавании)
	LenStep      float64       // длина одного шага или гребка в м
	Duration     time.Duration // продолжительность тренировки
	Weight       float64       // вес пользователя в кг
}

// distance возвращает дистанцию, которую преодолел пользователь.
// Формула расчета: количество_повторов * длина_шага / м_в_км.
func (t Training) distance() float64 {

	lenStepInKm := t.LenStep / MInKm
	distance := float64(t.Action) * lenStepInKm

	return distance
}

// meanSpeed возвращает среднюю скорость бега или ходьбы.
func (t Training) meanSpeed() float64 {
	if t.Duration == 0 {
		return 0.0
	}

	speed := t.distance() / t.Duration.Hours()

	return speed
}

// Calories возвращает количество потраченных килокалорий на тренировке.
// Пока возвращаем 0, так как этот метод будет переопределяться для каждого типа тренировки.
func (t Training) Calories() float64 {
	return 0.0
}

// InfoMessage содержит информацию о проведенной тренировке.
type InfoMessage struct {
	TrainingType string        // тип тренировки
	Duration     time.Duration // длительность тренировки
	Distance     float64       // расстояние, которое преодолел пользователь
	Speed        float64       // средняя скорость, с которой двигался пользователь
	Calories     float64       // количество потраченных килокалорий на тренировке
}

// TrainingInfo возвращает структуру InfoMessage, в которой хранится вся информация о проведенной тренировке.
func (t Training) TrainingInfo() InfoMessage {
	return InfoMessage{
		TrainingType: t.TrainingType,
		Duration:     t.Duration,
		Distance:     t.distance(),
		Speed:        t.meanSpeed(),
		Calories:     t.Calories(),
	}
}

// String возвращает строку с информацией о проведенной тренировке.
func (i InfoMessage) String() string {
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %v мин\nДистанция: %.2f км.\nСр. скорость: %.2f км/ч\nПотрачено ккал: %.2f\n",
		i.TrainingType,
		i.Duration.Minutes(),
		i.Distance,
		i.Speed,
		i.Calories,
	)
}

// CaloriesCalculator интерфейс для структур: Running, Walking и Swimming.
type CaloriesCalculator interface {
	Calories() float64
	TrainingInfo() InfoMessage
}

// Константы для расчета потраченных килокалорий при беге.
const (
	CaloriesMeanSpeedMultiplier = 18   // множитель средней скорости бега
	CaloriesMeanSpeedShift      = 1.79 // коэффициент изменения средней скорости
)

// Running структура, описывающая тренировку Бег.
type Running struct {
	Training
}

// Calories возвращает количество потраченных килокалория при беге.
// Формула расчета: (18 * средняя_скорость_в_км/ч + 1.79) * вес_спортсмена_в_кг / м_в_км * время_тренировки_в_часах * мин_в_часе
// Это переопределенный метод Calories() из Training.
func (r Running) Calories() float64 {
	if r.Duration == 0 {
		return 0.0
	}

	speed := CaloriesMeanSpeedMultiplier*r.meanSpeed() + CaloriesMeanSpeedShift
	speedInKm := speed / MInKm
	durationInHours := r.Duration.Hours() * MinInHours
	calories := speedInKm * r.Weight * durationInHours

	return calories
}

// TrainingInfo возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (r Running) TrainingInfo() InfoMessage {
	return InfoMessage{
		TrainingType: r.TrainingType,
		Duration:     r.Duration,
		Distance:     r.distance(),
		Speed:        r.meanSpeed(),
		Calories:     r.Calories(),
	}
}

// Константы для расчета потраченных килокалорий при ходьбе.
const (
	CaloriesWeightMultiplier      = 0.035 // коэффициент для веса
	CaloriesSpeedHeightMultiplier = 0.029 // коэффициент для роста
	KmHInMsec                     = 0.278 // коэффициент для перевода км/ч в м/с
	UnknownPowerMultiplier        = 2
)

// Walking структура описывающая тренировку Ходьба.
type Walking struct {
	Training
	Height float64 // рост пользователя.
}

// Calories возвращает количество потраченных килокалорий при ходьбе.
// Формула расчета: (0.035 * вес_спортсмена_в_кг + (средняя_скорость_в_метрах_в_секунду**2 / рост_в_метрах) * 0.029 * вес_спортсмена_в_кг) * время_тренировки_в_часах * мин_в_ч.
// Это переопределенный метод Calories() из Training.
func (w Walking) Calories() float64 {
	if w.Duration == 0 || w.Weight == 0 {
		return 0.0
	}

	weight := CaloriesWeightMultiplier * w.Weight
	speed := math.Pow(w.meanSpeed()*KmHInMsec, UnknownPowerMultiplier)
	heightInM := w.Height / CmInM
	durationInHours := w.Duration.Hours() * MinInHours
	calories := (weight + speed/heightInM*CaloriesSpeedHeightMultiplier*w.Weight) * durationInHours

	return calories
}

// TrainingInfo возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (w Walking) TrainingInfo() InfoMessage {
	return InfoMessage{
		TrainingType: w.TrainingType,
		Duration:     w.Duration,
		Distance:     w.distance(),
		Speed:        w.meanSpeed(),
		Calories:     w.Calories(),
	}
}

// Константы для расчета потраченных килокалорий при плавании.
const (
	SwimmingLenStep                  = 1.38 // длина одного гребка
	SwimmingCaloriesMeanSpeedShift   = 1.1  // коэффициент изменения средней скорости
	SwimmingCaloriesWeightMultiplier = 2    // множитель веса пользователя
)

// Swimming структура, описывающая тренировку Плавание.
type Swimming struct {
	Training
	LengthPool int // длина бассейна.
	CountPool  int // количество пересечений бассейна.
}

// meanSpeed возвращает среднюю скорость при плавании.
// Формула расчета: длина_бассейна * количество_пересечений / м_в_км / продолжительность_тренировки
// Это переопределенный метод Calories() из Training.
func (s Swimming) meanSpeed() float64 {
	if s.Duration == 0 {
		return 0.0
	}

	distance := float64(s.LengthPool) * float64(s.CountPool)
	distanceInKm := distance / MInKm

	return distanceInKm / s.Duration.Hours()
}

// Calories возвращает количество калорий, потраченных при плавании.
// Формула расчета: (средняя_скорость_в_км/ч + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * вес_спортсмена_в_кг * время_тренировки_в_часах
// Это переопределенный метод Calories() из Training.
func (s Swimming) Calories() float64 {
	calories := (s.meanSpeed() + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * s.Weight * s.Duration.Hours()

	return calories
}

// TrainingInfo returns info about swimming training.
// Это переопределенный метод TrainingInfo() из Training.
func (s Swimming) TrainingInfo() InfoMessage {
	return InfoMessage{
		TrainingType: s.TrainingType,
		Duration:     s.Duration,
		Distance:     s.distance(),
		Speed:        s.meanSpeed(),
		Calories:     s.Calories(),
	}
}

// ReadData возвращает информацию о проведенной тренировке.
func ReadData(training CaloriesCalculator) string {
	return fmt.Sprint(training.TrainingInfo())
}

func main() {
	swimming := Swimming{
		Training: Training{
			TrainingType: "Плавание",
			Action:       2000, //nolint:gomnd // it's data
			LenStep:      SwimmingLenStep,
			Duration:     90 * time.Minute, //nolint:gomnd // it's data
			Weight:       85,               //nolint:gomnd // it's data
		},
		LengthPool: 50, //nolint:gomnd // it's data
		CountPool:  5,  //nolint:gomnd // it's data
	}

	fmt.Println(ReadData(swimming)) //nolint:forbidigo // it's output

	walking := Walking{
		Training: Training{
			TrainingType: "Ходьба",
			Action:       20000, //nolint:gomnd // it's data
			LenStep:      LenStep,
			Duration:     3*time.Hour + 45*time.Minute,
			Weight:       85, //nolint:gomnd // it's data
		},
		Height: 185, //nolint:gomnd // it's data
	}

	fmt.Println(ReadData(walking)) //nolint:forbidigo // it's output

	running := Running{
		Training: Training{
			TrainingType: "Бег",
			Action:       5000, //nolint:gomnd // it's data
			LenStep:      LenStep,
			Duration:     30 * time.Minute, //nolint:gomnd // it's data
			Weight:       85,               //nolint:gomnd // it's data
		},
	}

	fmt.Println(ReadData(running)) //nolint:forbidigo // it's output
}
