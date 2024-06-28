package app

import (
	"context"
	"fmt"
	"log"
	"main/internal/bootstrap"
	"main/internal/config"
	"main/internal/repositories/autorepository/autosqlx"
	"main/internal/services/autoservice"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg config.Config) error {
	db, err := bootstrap.InitSqlxDB(cfg)
	if err != nil {
		return err
	}

	autoService := autoservice.New(autosqlx.New(db))

	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.ServerHost, cfg.ServerPort),
		Handler: autoService.GetHandler(),
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	gracefullyShutdown(ctx, cancel, server)
	return nil
}

func gracefullyShutdown(ctx context.Context, cancel context.CancelFunc, server *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)
	<-ch
	if err := server.Shutdown(ctx); err != nil {
		// log.Warning(err)
		log.Println(err)
	}
	cancel()
}
