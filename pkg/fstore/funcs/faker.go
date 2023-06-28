package funcs

import (
	"github.com/jaswdr/faker"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.AddGroup("faker", new(Faker).init)
}

type Faker struct {
	faker.Faker
}

func (f *Faker) init() *Faker {
	f.Faker = faker.New()

	return f
}
