package http

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// listenCmd represents HTTP server listen command
type listenCmd struct {
	ctx context.Context
	cfg Config
	app *fiber.App
}

// Run HTTP server
func (s *listenCmd) Run() error {
	ln, err := net.Listen("tcp4", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return fmt.Errorf("net listen: %w", err)
	}
	return s.app.Listener(&listen{
		Listener: ln,
	})
}

// Stop HTTP server
func (s *listenCmd) Stop() {
	_ = s.app.Shutdown()
}

// listen is a wrapper over [net.Listener]
type listen struct {
	net.Listener
}

// Accept wraps connection
func (l *listen) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}
	return &connection{
		Conn: conn,
	}, nil
}

// connection is a wrapper over [net.Conn]
type connection struct {
	net.Conn
	sync.Mutex

	n int
	b []byte
}

// Read reads data from connection and handles a special case for connection status check
func (a *connection) Read(b []byte) (n int, err error) {
	a.Lock()
	defer a.Unlock()

	// Special case to check if connection is still active
	if len(b) == 0 {
		bp := make([]byte, 1)

		n, err = a.Conn.Read(bp)
		if err != nil || n == 0 {
			return 0, err
		}

		if a.b == nil {
			a.b = bp
		} else {
			a.b = append(a.b, bp[0])
		}
		a.n += n

		return 0, nil
	}

	// Return buffered data back
	if a.n > 0 {
		n = copy(b, a.b)
		a.b = a.b[n:]
		a.n -= n
		return n, nil
	}

	return a.Conn.Read(b)
}
