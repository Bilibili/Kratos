package log

import (
	c "context"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	// DefaultCaller is a Valuer that returns the file and line.
	DefaultCaller = Caller(3)

	// DefaultTimestamp is a Valuer that returns the current wallclock time.
	DefaultTimestamp = Timestamp(time.RFC3339)
)

// Valuer is returns a log value.
type Valuer func(ctx c.Context) interface{}

// Value return the function value.
func Value(ctx c.Context, v interface{}) interface{} {
	if v, ok := v.(Valuer); ok {
		return v(ctx)
	}
	return v
}

// Caller returns returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int) Valuer {
	return func(c.Context) interface{} {
		_, file, line, _ := runtime.Caller(depth)
		if strings.LastIndex(file, "github.com/go-kratos/kratos/log") > 0 {
			_, file, line, _ = runtime.Caller(depth + 1)
		}
		idx := strings.LastIndexByte(file, '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

// Timestamp returns a timestamp Valuer with a custom time format.
func Timestamp(layout string) Valuer {
	return func(c.Context) interface{} {
		return time.Now().Format(layout)
	}
}

func bindValues(ctx c.Context, keyvals []interface{}) {
	for i := 1; i < len(keyvals); i += 2 {
		if v, ok := keyvals[i].(Valuer); ok {
			keyvals[i] = v(ctx)
		}
	}
}

func containsValuer(keyvals []interface{}) bool {
	for i := 1; i < len(keyvals); i += 2 {
		if _, ok := keyvals[i].(Valuer); ok {
			return true
		}
	}
	return false
}
