package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {

	http.HandleFunc("/", handler)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

// w表示response对象，r表示request对象
func handler(w http.ResponseWriter, r *http.Request) {
	results, err := searching(context.Background(), []string{"golang"})
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, results)
	}
}

// Result 关键字查找结果
type Result struct {
	Data string
	Hot  string
}

func searching(ctx context.Context, keywords []string) ([]Result, error) {
	g, ctx := errgroup.WithContext(ctx)
	var collection []Result
	for _, keyword := range keywords {
		query := keyword
		g.Go(func() error {
			res, err := google(query)
			if err != nil {
				return err
			}
			collection = append(collection, *res)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		log.Print(err)
		return nil, err
	}
	return collection, nil
}

func google(keyword string) (*Result, error) {
	return nil, fmt.Errorf("google failed：%s", keyword)
}
