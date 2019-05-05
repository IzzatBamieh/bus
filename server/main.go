package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := NewLogger()

	logger.Info("Starting...")

	provider, err := NewProvider()
	if err != nil {
		logger.Fatal(err)
	}
	bus, err := NewBus(logger, provider)
	if err != nil {
		logger.Fatal(err)
	}
	service, err := NewLambdaNodeForgeService(provider, nil)
	if err != nil {
		logger.Fatal(err)
	}
	handler := NewHandler(bus, logger, service)
	router := httprouter.New()
	router.OPTIONS("/services", handler.PostServices)
	router.POST("/services", handler.PostServices)
	router.GET("/services", handler.GetServices)
	router.DELETE("/services/:name", handler.DeleteService)

	logger.Info("HTTP server started at 127.0.0.1:9000")
	logger.Fatal(http.ListenAndServe("127.0.0.1:9000", router))
}

func testSendReceive(logger *Logger, bus *Bus) {
	go func() {
		for {
			logger.Info("Sending message:", "foo")
			if err := bus.Send("test_01", []byte("foo")); err != nil {
				logger.Fatal(err)
			}
			time.Sleep(2 * time.Second)
		}
	}()
	receiver, err := bus.Receive("test_01", "consumer_01")
	if err != nil {
		logger.Fatal(err)
	}
	for {
		message, err := receiver.Next()
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("Received message:", string(message.Body))
		if err = message.Ack(); err != nil {
			logger.Fatal(err)
		}
	}
}
