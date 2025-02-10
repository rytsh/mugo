package faker

import (
	"github.com/jaswdr/faker"
)

type Faker struct {
	faker.Faker
}

func New() Faker {
	return Faker{
		Faker: faker.New(),
	}
}
