package logger

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	*logrus.Logger
}

func NewLogger(writer io.Writer) *CustomLogger {
	Logger := &CustomLogger{logrus.New()}
	Formatter := new(logrus.JSONFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Logger.SetFormatter(Formatter)
	mw := io.MultiWriter(os.Stdout, writer)
	Logger.SetOutput(mw)
	return Logger
}

func (l *CustomLogger) relative(path string) string {
	return strings.TrimPrefix(filepath.ToSlash(path), configs.PrefixPath)
}

func (l *CustomLogger) LogError(ctx context.Context, err error) {
	l.WithFields(logrus.Fields{
		"id": l.GetIdFromContext(ctx),
	}).Error(err)
}

func (l *CustomLogger) GetIdFromContext(ctx context.Context) string {
	_, fn, line, _ := runtime.Caller(2)
	user, ok := ctx.Value(configs.RequestUser).(models.User)
	if !ok {
		l.WithFields(logrus.Fields{
			"id":   "NO_ID",
			"file": l.relative(fn),
			"line": line,
		}).Warn("can't get request id from context")
		return ""
	}
	return strconv.Itoa(user.ID)
}

func (l *CustomLogger) StartReq(r http.Request, rid string) {
	l.WithFields(logrus.Fields{
		"id":         rid,
		"usr_addr":   r.RemoteAddr,
		"req_URI":    r.RequestURI,
		"method":     r.Method,
		"user_agent": r.UserAgent(),
	}).Info("request started")
}

func (l *CustomLogger) EndReq(start time.Time, ctx context.Context) {
	l.WithFields(logrus.Fields{
		"id":              l.GetIdFromContext(ctx),
		"elapsed_time,μs": time.Since(start).Microseconds(),
	}).Info("request ended")
}

func (l *CustomLogger) LogWarning(ctx context.Context, msg string) {
	l.WithFields(logrus.Fields{
		"id": l.GetIdFromContext(ctx),
	}).Warn(msg)
}
