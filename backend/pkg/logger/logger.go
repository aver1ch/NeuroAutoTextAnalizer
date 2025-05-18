package logger

import (
	"log/slog"
	"os"
	"reflect"
	"strings"

	"github.com/lmittmann/tint"
)

func Init(format string, level string) *slog.Logger {
	format = strings.ToUpper(format)
	level = strings.ToUpper(level)

	opts := &slog.HandlerOptions{} //nolint:exhaustruct

	switch level {
	case "INFO":
		opts.Level = slog.LevelInfo
	case "WARN":
		opts.Level = slog.LevelWarn
	case "ERROR":
		opts.Level = slog.LevelError
	default:
		opts.Level = slog.LevelDebug
		opts.AddSource = true
	}

	var handler slog.Handler = slog.NewJSONHandler(os.Stdout, opts)
	if format == "TEXT" {
		handler = tint.NewHandler(os.Stdout, &tint.Options{ //nolint:exhaustruct
			AddSource:  opts.AddSource,
			Level:      opts.Level,
			TimeFormat: "15:04:05.000000000",
			NoColor:    false,
		})
	}

	log := slog.New(handler)

	log.Info("Init logger", slog.Group(
		"config",
		slog.String("format", format),
		slog.String("level", level),
	))

	return log
}

type Secret string

func (s Secret) LogValue() slog.Value {
	if s == "" {
		return slog.StringValue("[EMPTY]")
	}

	return slog.StringValue("[REDACTED]")
}

func NewAttr(k string, v slog.Value) slog.Attr {
	return slog.Attr{
		Key:   k,
		Value: v,
	}
}

func Nullable(value any) slog.Value {
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		return slog.AnyValue(val.Elem())
	}

	return slog.AnyValue(value)
}
