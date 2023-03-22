package bookcontroller

import (
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func (r GetBookByIDRequest) Validate() (err error) {
	if r.ID <= 0 {
		return errors.AddTrace(errors.MissingParam)
	}

	return nil
}

func (r SearchBookRequest) Validate() (err error) {
	if r.Query == "" {
		return errors.AddTrace(errors.MissingParam)
	}

	return nil
}
