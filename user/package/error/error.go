package customerror

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/configs"
)

type CustomError struct {
	err  string
	file string
	line int
	code int
}

func NewCustomError(err error, code int, skip int) *CustomError {
	var fn string
	var line int
	_, fn, line, _ = runtime.Caller(skip)
	return &CustomError{err.Error(), relative(fn), line, code}
}

func relative(path string) string {

	return strings.TrimPrefix(filepath.ToSlash(path), configs.PrefixPath)
}

func ParseCode(err error) int {
	index := strings.Index(err.Error(), "]")
	code, err := strconv.Atoi(err.Error()[6:index])
	if err != nil {
		return http.StatusInternalServerError
	}
	return code
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code:[%d] file:[%s]  line:[%d]  error:[%s]", e.code, e.file, e.line, e.err)
}
