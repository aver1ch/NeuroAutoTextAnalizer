package handlers

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

func Logging(log *slog.Logger) echo.MiddlewareFunc {
	log = log.WithGroup("http_server")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			if err != nil {
				log.ErrorContext(c.Request().Context(),
					"error during request",
					slog.Any("error", err),
					slog.Group(
						"request",
						slog.String("method", c.Request().Method),
						slog.String("path", c.Request().URL.Path),
						slog.Int("status", c.Response().Status),
						slog.Duration("dur", time.Since(start)),
						slog.String("remote_ip", c.Request().RemoteAddr),
						slog.String("user_agent", c.Request().UserAgent()),
					))
			} else {
				log.DebugContext(c.Request().Context(),
					"request", slog.Group(
						"request",
						slog.String("method", c.Request().Method),
						slog.String("path", c.Request().URL.Path),
						slog.Int("status", c.Response().Status),
						slog.Duration("dur", time.Since(start)),
						slog.String("remote_ip", c.Request().RemoteAddr),
						slog.String("user_agent", c.Request().UserAgent()),
					))
			}
			return err
		}
	}
}
