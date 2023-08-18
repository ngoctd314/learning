package get

import (
	"errors"
	"go-learn/acme/internal/modules/data"
)

var (
	// error thrown when the requested person is not in the database
	errPersonNotFound = errors.New("person not found")
)

type Getter struct{}

func (g *Getter) Do(ID int) (*data.Person, error) {
	person, err := data.Load(ID)
	if err != nil {
		if err == data.ErrNotFound {
			return nil, errPersonNotFound
		}
	}

	return person, err
}
