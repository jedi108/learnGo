// Simple HTTP echo server written in Go.
package main

import (
	"bytes"
	"context"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var server = &EchoServer{
	Writer: os.Stdout,
	Logger: log.New(os.Stderr, "", log.LstdFlags),
}

func init() {
	flag.StringVar(&server.Addr, "addr", ":8081", "bind address")
}

func main() {
	flag.Parse()

	errCh := make(chan error)
	go func() {
		if err := server.Serve(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()
	log.Printf("Starting echo server at %s\n", server.Addr)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	select {
	case err, ok := <-errCh:
		if ok {
			log.Println("Echo server error:", err)
		}

	case sig := <-sigCh:
		log.Printf("Signal %s received\n", sig)
		if err := server.Shutdown(); err != nil {
			log.Println("Failed to shutdown echo server:", err)
		}
		log.Println("Echo server shutdown")
	}

}

type EchoServer struct {
	Addr      string
	Writer    io.Writer
	Logger    *log.Logger
	MustClose bool
	server    *http.Server
}

func (e *EchoServer) Serve() error {
	if e.Writer == nil {
		e.Writer = ioutil.Discard
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", e.echo)

	e.server = &http.Server{
		Addr:    e.Addr,
		Handler: mux,
	}

	if err := e.server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}

	return nil
}

func (e *EchoServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return e.ShutdownWithContext(ctx)
}

func (e *EchoServer) ShutdownWithContext(ctx context.Context) error {
	if err := e.server.Shutdown(ctx); err != nil {
		if e.MustClose {
			e.server.Close()
		}
		return err
	}

	return nil
}

func (e *EchoServer) echo(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	_ = r.Write(&buf)
	_, _ = e.Writer.Write([]byte("----------------------------\n"))
	_, _ = e.Writer.Write(buf.Bytes())
	_, _ = w.Write(buf.Bytes())
}
