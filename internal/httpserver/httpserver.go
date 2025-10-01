package httpserver

import (
	"cfg-server/internal/svclogger"
	"context"
	"fmt"
	"time"

	fiberprom "github.com/carousell/fiber-prometheus-middleware"
	gojson "github.com/goccy/go-json"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HTTP struct {
	Port     string `yaml:"port" json:"port" default:"8080"`
	Timeouts struct {
		Read     time.Duration `yaml:"read" json:"read" default:"30s"`
		Write    time.Duration `yaml:"write" json:"write" default:"30s"`
		Idle     time.Duration `yaml:"idle" json:"idle" default:"30s"`
		Shutdown time.Duration `yaml:"shutdown" json:"shutdown" default:"30s"`
	} `yaml:"timeouts" json:"timeouts"`
}

type Server struct {
	ctx     context.Context
	Handler *fiber.App
	notify  chan error
}

// New -.
func New(c context.Context, log *svclogger.Log, opts *HTTP) *Server {
	app := fiber.New(fiber.Config{
		JSONEncoder:           gojson.Marshal,
		JSONDecoder:           gojson.Unmarshal,
		DisableStartupMessage: true,
		ReadTimeout:           opts.Timeouts.Read,
		WriteTimeout:          opts.Timeouts.Write,
		IdleTimeout:           opts.Timeouts.Idle,
	})

	app.Use(recover.New())

	//Logger settings
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger:   log.Logger,
		Fields:   []string{"latency", "status", "method", "url", "ua", "ip", "bytesSent"},
		SkipURIs: mySkipper(),
		Messages: []string{"-"},
	}))

	//metrics settings
	prometheus := fiberprom.NewPrometheus("")
	prometheus.Use(app)

	s := &Server{
		notify:  make(chan error, 1),
		Handler: app,
		ctx:     c,
	}

	s.start(fmt.Sprint(":", opts.Port))

	return s
}

func (s *Server) start(aPort string) {
	go func() {
		s.notify <- s.Handler.Listen(aPort)
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown(shutdownTimeout time.Duration) error {
	ctx, cancel := context.WithTimeout(s.ctx, shutdownTimeout)
	defer cancel()

	return s.Handler.ShutdownWithContext(ctx)
}
