package errors

import (
	"database/sql"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Fields map[string]interface{}

type Errs struct {
	Err error

	Code   string
	Reason string

	Traces []string
}

func New(args ...interface{}) *Errs {
	err := &Errs{}

	for _, arg := range args {
		switch arg.(type) {
		case Errs:
			errTemp := arg.(Errs)
			err = &errTemp
		case *Errs:
			err = arg.(*Errs)
		case string:
			err.Err = errors.New(arg.(string))
		case error:
			err.Err = arg.(error)
		default:
			err.Err = errors.New("unknown error")
		}
	}

	if err.Err == nil {
		err.Err = errors.New("unknown error")
		if len(err.Reason) > 0 {
			err.Err = errors.New(err.Reason)
		}
	}

	return err
}

func (e *Errs) Error() string {
	if e == nil {
		return ""
	}

	if e.Err == nil {
		return ""
	}

	return e.Err.Error()
}

func (e *Errs) ErrMessage() string {
	if e.Code == "" {
		e.Code = "50000"
	}

	return fmt.Sprintf("Error - %s - with message [%s] %s in %s", time.Now(), e.Code, e.Error(), strings.Join(e.Traces, " "))
}

func AddTrace(err interface{}) *Errs {
	errs := New(err)
	errs.Traces = append(errs.Traces, getLineOfCode(2))

	return errs
}

func IsMatchByCode(err1 error, err2 error) bool {
	if err1 == nil || err2 == nil {
		return false
	}

	var (
		errs1 *Errs
		errs2 *Errs
		ok    bool
	)

	if err1 != nil {
		errs1, ok = err1.(*Errs)
		if !ok {
			return false
		}
	}

	if err2 != nil {
		errs2, ok = err2.(*Errs)
		if !ok {
			return false
		}
	}

	if errs1 == nil || errs2 == nil {
		return false
	}

	return errs1.Code == errs2.Code
}

func getLineOfCode(skip int) string {
	_, file, line, _ := runtime.Caller(skip)

	path := strings.Split(file, "/")
	if len(path) > 4 {
		start := len(path) - 4
		file = "/" + strings.Join(path[start:], "/")
	}

	return fmt.Sprintf("%s[%d]", file, line)
}

func IsNotFound(err error) bool {
	return err.Error() == sql.ErrNoRows.Error()
}
