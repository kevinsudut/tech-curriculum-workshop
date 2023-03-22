package response

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func buildResponse(data interface{}, err error) []byte {
	var (
		respData   []byte
		errMarshal error
	)

	if err != nil {
		errorMsg := err.Error()

		errLib, ok := err.(*errors.Errs)
		if ok {
			errorMsg = errLib.ErrMessage()
		}

		respData, errMarshal = jsoniter.Marshal(Response{
			Error: errorMsg,
		})
	} else {
		respData, errMarshal = jsoniter.Marshal(Response{
			Data: data,
		})
	}

	if errMarshal != nil {
		fmt.Println("ERR Marshal", err)
	}

	return respData
}
