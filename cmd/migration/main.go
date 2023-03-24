package main

import (
	"fmt"

	"github.com/kevinsudut/tech-curriculum-workshops/lib/database"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func main() {
	err := database.Migration()
	if err != nil {
		errorMsg := err.Error()

		errLib, ok := err.(*errors.Errs)
		if ok {
			errorMsg = errLib.ErrMessage()
		}

		fmt.Println(errorMsg)
	}
}
