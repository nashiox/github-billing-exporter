package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/xerrors"
)

func Run(args *Args) error {
	var mode apiMode
	if args.Organization != "" {
		mode = orgMode
	} else if args.User != "" {
		mode = userMode
	}

	go getGitHubActionsBilling(mode, args)
	go getGitHubPackagesBilling(mode, args)
	go getGitHubSharedStorageBilling(mode, args)

	ctx, cancel := context.WithCancel(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "/metrics")
	})
	mux.Handle("/metrics", promhttp.Handler())

	httpServer := &http.Server{
		Addr:        ":" + strconv.Itoa(args.Port),
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	httpServer.RegisterOnShutdown(cancel)

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			xerrors.Errorf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		return xerrors.Errorf("shutdown error: %v\n", err)
	}

	log.Printf("gracefully stopped\n")
	return nil
}
