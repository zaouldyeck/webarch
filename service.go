package architecture

import (
	"fmt"
)

// Person is how a person is stored.
type Person struct {
	First string
}

// Accessor is how to retrieve a person.
// If the person does not exist, return zero value.
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return p, fmt.Errorf("no person with n of %d found", n)
	}
	return p, nil
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}
