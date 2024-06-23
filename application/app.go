package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb   *redis.Client
}

// add the owner of this method(receiver) to the function signature
// similar to .this in JavaScript
func(a *App) Start(ctx context.Context) error {
	// first define the http server
	server := &http.Server{
		Addr: ":8080",
		Handler: a.router,
	}

	// call ping for the redis client
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("redis failed to start: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis connection: ", err)
		}
	}()

	fmt.Println("Starting server...")

	ch := make(chan error, 1)

	go func() {
	err = server.ListenAndServe()
	if err != nil {
		ch <- fmt.Errorf("server failed to start: %w", err)
	}
		close(ch)
}()
	select {
	case err = <-ch: 
	return err 
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
	return nil
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb: redis.NewClient(&redis.Options{}),
	}
	return app
}

