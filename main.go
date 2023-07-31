package main

import (
	"log"
	"os"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	logFileName := "logfile_start_time" + time.Now().Format("2006-01-02_15:04:05")
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	for range time.Tick(time.Minute) {
		log.SetOutput(f)
		googleRes := ping_url("www.google.com")
		routerRes := ping_url("192.168.1.1")
		log.Printf("Google: %v ms\n", googleRes.Milliseconds())
		log.Printf("Router: %v ms\n", routerRes.Milliseconds())
	}

}

func ping_url(url string) time.Duration {
	pinger, err := ping.NewPinger(url)
	if err != nil {
		panic(err)
	}
	pinger.Count = 5
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
	return pinger.Statistics().AvgRtt
}
