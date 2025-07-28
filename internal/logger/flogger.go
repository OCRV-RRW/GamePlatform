package flogger

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FiberHandler struct {
	slog.Handler
}

type LoggerMiddelware struct {
	Debug bool
}

func NewLoggerMiddelware(debug bool) LoggerMiddelware {
	return LoggerMiddelware{
		Debug: debug,
	}
}

func (m *LoggerMiddelware) Handle(c *fiber.Ctx) error {
	requestId := uuid.NewString()
	c.Locals(Method, c.Method())
	c.Locals(Path, c.Path())
	c.Locals(RequestId, requestId)
	c.Locals(SourceIP, c.IP())

	if m.Debug {
		slog.InfoContext(c.Context(), "")
	}
	return c.Next()
}

func (l FiberHandler) Handle(ctx context.Context, r slog.Record) error {
	if ctx.Value(RequestId) == nil {
		return l.Handler.Handle(ctx, r)
	}

	requestId := ctx.Value(RequestId).(string)
	sourceIp := ctx.Value(SourceIP).(string)
	path := ctx.Value(Path).(string)
	method := ctx.Value(Method).(string)

	requestGroup := slog.Group(
		string(Request),
		slog.String(string(RequestId), requestId),
		slog.String(string(SourceIP), sourceIp),
		slog.String(string(Method), method),
		slog.String(string(Path), path),
	)

	r.AddAttrs(requestGroup)

	return l.Handler.Handle(ctx, r)
}

type ContextKey string

const (
	Request   ContextKey = "request"
	RequestId ContextKey = "request_id"
	Method    ContextKey = "method"
	Path      ContextKey = "path"
	SourceIP  ContextKey = "source_ip"
)
