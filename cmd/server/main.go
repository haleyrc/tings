package main

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/haleyrc/log"
	"github.com/haleyrc/router"
	"github.com/haleyrc/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(log.NewJSONHandler(log.Debug()))

	router := router.New()

	router.GET("/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	var h http.Handler = router
	h = log.ParseRequestInfo(h)

	server := server.New("8086", h)
	logger.Info("server listening", slog.String("addr", server.Addr()))
	if err := server.ListenAndServe(ctx); err != nil {
		logger.Error("server shutdown unexpectedly", slog.Any("error", err))
	}
}
