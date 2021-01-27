package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	ctx, cancelFunc := context.WithCancel(context.Background())

	// 监听中断信号
	ch := make(chan os.Signal)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT)

	go func(ctx context.Context) {
		select {
		case <-ch:
			log.Print("信号，closed")
			listener.Close()
		case <-ctx.Done():
			cancelFunc()
		}
	}(ctx)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		fmt.Printf("received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// handle connections in a new goroutine.
		go handleConnection(ctx, conn)

	}
}

// handleConnection handle connections
func handleConnection(ctx context.Context, conn net.Conn) {
	var ch = make(chan string)
	go handleRead(ctx, conn, ch)
	go handleWrite(ctx, conn, ch)
}

// handleRead  read connection
func handleRead(ctx context.Context, conn net.Conn, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			conn.Close()
			return
		default:
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			ch <- data
		}
	}
}

// handleRead  read connection
func handleWrite(ctx context.Context, conn net.Conn, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("连接， disconnected")
			return
		default:
			data := <-ch
			fmt.Println("write message：" + data)
		}
	}
}
