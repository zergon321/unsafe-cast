package unsafecast

import "fmt"

type ErrorSizeUnmatch struct {
	fromLength int
	fromSize   int64
	toSize     int64
}

func (err *ErrorSizeUnmatch) Error() string {
	return fmt.Sprintf(
		"size mismatch: source length = '%d',"+
			"source size = '%d', destination size = '%d'",
		err.fromLength, err.fromSize, err.toSize)
}
