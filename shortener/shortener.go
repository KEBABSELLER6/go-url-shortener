package shortener

import (
	"errors"

	"github.com/teris-io/shortid"
)

func GenerateShortId() (string, error) {
	sid, sidErr := shortid.New(1, shortid.DefaultABC, 2342)

	if sidErr == nil {
		id, err := sid.Generate()
		if err == nil {
			return id, nil
		} else {
			return "", errors.New("Can't generate id")
		}
	} else {
		return "", errors.New("Can't get generator")
	}
}
