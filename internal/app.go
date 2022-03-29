package internal

import (
	"context"
	"fmt"
	"github.com/yalagtyarzh/unibot/internal/config"
	"github.com/yalagtyarzh/unibot/internal/service"
	"github.com/yalagtyarzh/unibot/pkg/client/youtube"
	"github.com/yalagtyarzh/unibot/pkg/logging"
	tele "gopkg.in/telebot.v3"
	"net/http"
	"time"
)

type App interface {
	Run()
}

type app struct {
	cfg            *config.Config
	logger         *logging.Logger
	httpServer     *http.Server
	youtubeService service.YoutubeService
}

func NewApp(logger *logging.Logger, cfg *config.Config) (App, error) {
	ytClient := youtube.NewClient(cfg.Youtube.AccessToken, cfg.Youtube.APIURL, &http.Client{})
	youtubeService := service.NewYoutubeService(ytClient, logger)

	logger.Println("")
	return &app{
		cfg:            cfg,
		logger:         logger,
		youtubeService: youtubeService,
	}, nil
}

func (a *app) Run() {
	a.logger.Info(a.cfg.Telegram.Token == "")
	pref := tele.Settings{
		Token:   a.cfg.Telegram.Token,
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: a.OnBotError,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		a.logger.Fatal(err)
		return
	}

	b.Handle("/yt", func(c tele.Context) error {
		sendText := "Твой трек не найден"
		trackName := c.Message().Payload
		name, err := a.youtubeService.FindTrackByName(context.Background(), trackName)
		if err != nil {
			a.logger.Error(err)
			return c.Send(sendText)
		}

		sendText = name
		return c.Send(fmt.Sprintf("Вот твой трек %s", trackName))
	})

	b.Start()
}

func (a *app) OnBotError(err error, ctx tele.Context) {
	a.logger.Error(err)
}
