package customerror

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
)

type CustomError struct {
	err  string
	file string
	line int
}

func NewCustomError(err string) *CustomError {
	_, fn, line, _ := runtime.Caller(1)
	return &CustomError{err, relative(fn), line}
}

func relative(path string) string {
	return strings.TrimPrefix(filepath.ToSlash(path), configs.PrefixPath)
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("file:[%s]  line:[%d]  error:[%s]", e.file, e.line, e.err)
}
