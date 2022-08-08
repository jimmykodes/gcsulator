package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/jimmykodes/prmrtr"

	"github.com/jimmykodes/gcsulator/internal/handlers"
)

func main() {
	addr := flag.String("addr", ":8080", "server address")
	if *addr == "" {
		log.Println("invalid address")
		return
	}

	flag.Parse()

	router := prmrtr.NewRouter(prmrtr.NotFoundHandlerOption(http.HandlerFunc(handlers.NotFound)))

	router.HandleFunc("/:bucket/:object", handlers.Get)
	router.HandleFunc("/storage/v1/b/:bucket/o/:object", handlers.Delete)
	router.HandleFunc("/upload/storage/v1/b/:bucket/o", handlers.Post)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "\t", r.URL.Path, "\t", r.URL.RawQuery)
		router.ServeHTTP(w, r)
	})

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	svr := &http.Server{
		Addr:    *addr,
		Handler: handler,
	}

	go func() {
		<-ctx.Done()
		log.Println("shutting down")
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := svr.Shutdown(timeoutCtx); err != nil {
			log.Println("error shutting down server", err)
		}
	}()

	log.Println("running at", *addr)
	if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("server error", err)
	}
}
