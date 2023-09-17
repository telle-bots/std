package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// middlewareContext cancels context if connection was closed
func middlewareContext(noCancel bool) fiber.Handler {
	return func(fCtx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(fCtx) || noCancel {
			fCtx.SetUserContext(fCtx.Context())
			return fCtx.Next()
		}

		conn := fCtx.Context().Conn()

		ctx, cancel := context.WithCancel(fCtx.Context())
		go func() {
			_, _ = conn.Read(nil)
			cancel()
		}()
		fCtx.SetUserContext(ctx)

		return fCtx.Next()
	}
}
