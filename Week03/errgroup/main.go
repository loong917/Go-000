package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello")
		})
		// 启动web服务，监听9090端口
		err := http.ListenAndServe("127.0.0.1:8080", nil)
		if err != nil {
			return fmt.Errorf("ListenAndServe:  %v", err)
		}
		return nil
	})

	g.Go(func() error {
		err := ListenSignal(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("error group: %v", err.Error())
	}

}

// ListenSignal 监听信号
func ListenSignal(ctx context.Context) error {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGQUIT)
	select {
	case <-ch:
		return errors.New("receive signal")
	case <-ctx.Done():
		return errors.New("game over")
	}
}
