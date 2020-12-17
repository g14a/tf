package sig

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	ProviderSig = make(chan os.Signal, 1)
	ResourceSig = make(chan os.Signal, 1)
)

func Handler(c chan os.Signal) {
	signal.Notify(c,
		syscall.SIGHUP, // kill -SIGHUP XXXX
		syscall.SIGINT, // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT)

	go func() {
		for sig := range c {
			log.Printf("[SERVER] Closing due to Signal: %s", sig)
			log.Printf("[SERVER] Graceful shutdown")

			fmt.Println("Done.")

			// Exit cleanly
			os.Exit(0)
		}
	}()
}
