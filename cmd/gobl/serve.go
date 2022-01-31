package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"

	"github.com/invopop/gobl"
)

const (
	defaultHTTPPort = 80

	// If you customize this server, you should change this.
	vendorName = "Invopop Ltd."
)

type serveOpts struct {
	httpPort int
}

func serve() *serveOpts {
	return &serveOpts{}
}

func (s *serveOpts) cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "serve",
		RunE: s.runE,
	}
	f := cmd.Flags()

	f.IntVarP(&s.httpPort, "port", "p", defaultHTTPPort, "HTTP port to listen on")

	return cmd
}

func (s *serveOpts) runE(cmd *cobra.Command, _ []string) error {
	ctx, cancel := context.WithCancel(cmd.Context())
	defer cancel()

	e := echo.New()

	e.GET("/", s.version())

	var startErr error
	go func() {
		err := e.Start(":" + strconv.Itoa(s.httpPort))
		if !errors.Is(err, http.ErrServerClosed) {
			startErr = err
		}
		cancel() // Ensure we stop the context if we have a startup error
	}()

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		return err
	}
	return startErr
}

func (s *serveOpts) version() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"gobl":    "Welcome",
			"version": gobl.VERSION,
			"vendor": map[string]interface{}{
				"name": vendorName,
			},
		})
	}
}
