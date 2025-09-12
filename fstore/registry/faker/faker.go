package faker

import (
	"github.com/jaswdr/faker"

	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddStruct("faker", New())
}

type Faker struct {
	faker.Faker
}

func New() Faker {
	return Faker{
		Faker: faker.New(),
	}
}
