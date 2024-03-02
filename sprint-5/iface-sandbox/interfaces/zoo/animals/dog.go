package animals

import (
	"fmt"

	"sandbox/interfaces/logger"
)

type Dog struct {
	AnimalStruct

	log logger.Logger
}

func NewDog(log logger.Logger, name string, fullness int) Dog {
	dog := Dog{
		log: log,
	}
	dog.name = name
	dog.fullness = fullness
	return dog
}

func (d Dog) Feed() {
	d.fullness++
	d.log.Info(fmt.Sprintf("So yummy! My fullness is %d.", d.fullness))
}

func (d Dog) SayHello() {
	d.log.Info(fmt.Sprintf("Hi! My name is %s. I'm dog.\n", d.name))
	d.fullness -= 5
}

func (d Dog) GetFullness() int {
	return d.fullness
}
