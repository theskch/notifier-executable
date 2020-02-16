package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pborman/getopt/v2"
	notif "github.com/theskch/notifier/http"
)

var (
	url      = getopt.StringLong("url", 0, "", "Target url for notifications (required)", "string")
	interval = getopt.DurationLong("interval", 'i', 5*time.Second, "Notification interval", "duration")
	help     = getopt.BoolLong("help", 'h', "Show context-sensitive help")
)

func main() {
	getopt.ParseV2()
	getopt.HelpColumn = 40

	if *help {
		getopt.PrintUsage(os.Stdout)
		os.Exit(1)
	}

	if *url == "" {
		getopt.PrintUsage(os.Stdout)
		os.Exit(1)
	}

	clientConfig := notif.ClientConfig{
		URL:   *url,
		Limit: 5,
	}

	client, err := notif.NewHTTPClient(clientConfig)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				message := scanner.Text()
				log.Printf("sending message: %s", message)
				client.SendMessage([]byte(message), callback)
			}

			time.Sleep(*interval)
			os.Stdin.Seek(0, 0)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT)

	<-done
	log.Print("executable stopped")
}

func callback(content []byte, err error) {
	if err != nil {
		log.Printf("error response received: %s", err)
	} else {
		log.Printf("response content: %s", string(content))
	}
}
