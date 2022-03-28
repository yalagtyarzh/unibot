package internal

import (
	"github.com/yalagtyarzh/unibot/internal/config"
	"github.com/yalagtyarzh/unibot/pkg/logging"
	tele "gopkg.in/telebot.v3"
	"net/http"
	"time"
)

type App interface {
	Run()
}

type app struct {
	cfg        *config.Config
	logger     *logging.Logger
	httpServer *http.Server
}

func NewApp(logger *logging.Logger, cfg *config.Config) (App, error) {
	logger.Println("")
	return &app{
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (a *app) Run() {
	a.logger.Info(a.cfg.Telegram.Token == "")
	pref := tele.Settings{
		Token:  a.cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		a.logger.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
