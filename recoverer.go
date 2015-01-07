package recoverer

import (
	"fmt"
)

func Catch(err *error) {
	if recovering := recover(); recovering != nil {
		if errorCast, ok := recovering.(error); ok {
			*err = errorCast
		} else {
			*err = fmt.Errorf("%+v", recovering)
		}
	}
}
