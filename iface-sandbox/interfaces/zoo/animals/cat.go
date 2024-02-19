package animals

import (
	"fmt"

	"sandbox/interfaces/logger"
)

type Cat struct {
	AnimalStruct

	log logger.Logger
}

func NewCat(log logger.Logger, name string, fullness int) Cat {
	cat := Cat{
		log: log,
	}
	cat.name = name
	cat.fullness = fullness
	return cat
}

func (c *Cat) ShowYourself() {
	c.log.Error("HI!!!")
}

func (c *Cat) Feed() {
	c.fullness++
	c.log.Warn(fmt.Sprintf("So yummy! My fullness is %d.", c.fullness))
}

func (c *Cat) SayHello() {
	c.log.Error(fmt.Sprintf("Hi! My name is %s. I'm cat.", c.name))
	c.fullness -= 2
}

func (c *Cat) GetFullness() int {
	return c.fullness
}
