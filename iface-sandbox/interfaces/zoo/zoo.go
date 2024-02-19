package zoo

import (
	"fmt"

	"sandbox/interfaces/logger"
	"sandbox/interfaces/zoo/animals"
)

//go:generate mockgen -s ./animals.go -d ./mock/mock_animals
type Animal interface {
	Feed()
	SayHello()
	GetFullness() int
}

func EmulateZoo(log logger.Logger) {
	var animalsSlice []Animal

	for _, name := range []string{"Bars", "Murzik", "Igor"} {
		newCat := animals.NewCat(log, name, 100)
		animalsSlice = append(animalsSlice, &newCat)
	}
	for _, name := range []string{"Sharik", "Rex", "Ignat"} {
		newDog := animals.NewDog(log, name, 100)
		animalsSlice = append(animalsSlice, newDog)
	}

	for _, a := range animalsSlice {
		a.SayHello()
		log.Info(fmt.Sprintf("My fullness is %d for now.", a.GetFullness()))
		a.Feed()
	}
}
