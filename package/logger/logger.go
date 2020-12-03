package logger

import (
	"context"
	"io"
	"net/http"
	"os"

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

func (l *CustomLogger) LogError(ctx context.Context, err error) {
	l.WithFields(logrus.Fields{}).Error(err)
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

func (l *CustomLogger) EndReq(since int64, ctx context.Context) {
	l.WithFields(logrus.Fields{
		"elapsed_time,μs": since,
	}).Info("request ended")
}

func (l *CustomLogger) LogWarning(ctx context.Context, msg string) {
	l.WithFields(logrus.Fields{}).Warn(msg)
}
