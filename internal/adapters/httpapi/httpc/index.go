package httpc

import (
	"context"
	"net/http"
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/services"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
)

type St struct {
	lg       interfaces.Logger
	eChan    chan<- error
	server   *http.Server
	services *services.Services
}

func New(lg interfaces.Logger, listen string, eChan chan<- error, services *services.Services) *St {
	api := &St{
		lg:       lg,
		eChan:    eChan,
		services: services,
	}

	api.server = &http.Server{
		Addr:    listen,
		Handler: api.router(),
	}

	return api
}

func (a *St) Start() {
	go func() {
		a.lg.Infow("Start rest-api", "addr", a.server.Addr)

		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			a.lg.Errorw("Http server is closed", err)
			a.eChan <- err
		}
	}()
}

func (a *St) Shutdown(timeout time.Duration) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
	defer ctxCancel()

	return a.server.Shutdown(ctx)
}
