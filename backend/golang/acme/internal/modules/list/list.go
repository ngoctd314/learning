package list

import (
	"errors"
	"go-learn/acme/internal/modules/data"
)

var (
	// error thrown when there are no people in the database
	errPeopleNotFound = errors.New("no people found")
)

type Lister struct{}

func (l *Lister) Do() ([]*data.Person, error) {
	people, err := l.load()
	if err != nil {
		return nil, err
	}

	if len(people) == 0 {
		// special processing for 0 people returned
		return nil, errPeopleNotFound
	}

	return people, nil
}

// load all people
func (l *Lister) load() ([]*data.Person, error) {
	people, err := data.LoadAll()
	if err != nil {
		if err == data.ErrNotFound {
			// By converting the error we are encapsulating the implementation details from our users.
			return nil, errPeopleNotFound
		}
		return nil, err
	}

	return people, nil
}
