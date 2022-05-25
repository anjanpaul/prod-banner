package main

import (
	"fmt"
	"math/rand"
	"mongo-fiber-api/app"
	"mongo-fiber-api/config"
	"mongo-fiber-api/handler"
	"mongo-fiber-api/routes"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	log "github.com/sirupsen/logrus"
)

const idleTimeout = 5 * time.Second

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
}

func main() {

	config.New()
	App := app.New()
	App.Bootstrap()
	defer func() { _ = App.Mongo.Disconnect() }()

	// setup App
	server := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	// setup middlewares
	server.Use(requestid.New())
	server.Use(recover.New())
	server.Use(cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("%s,%s", config.Params.CORSPermitted, config.Params.AirbringrDomain),
	}))

	server.Use(logger.New(logger.Config{
		Format:   "[${time}] ${status} ${locals:requestid} - ${latency} ${method} ${path}\n",
		TimeZone: "Asia/Dhaka",
	}))

	//App.SetCacheMiddleware(server)

	//routes
	Handler := handler.New()
	server.Get("/", Handler.Home)
	routes.Banner(server, Handler)
	server.Use(Handler.NotFound) // 404

	// Listen from a different goroutine
	go func() {
		if err := server.Listen(fmt.Sprintf(":%d", config.Params.Port)); err != nil {
			log.Fatal(err.Error())
		}
	}()

	rand.Seed(time.Now().UnixNano())
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = server.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here

	fmt.Println("Fiber was successful shutdown.")
}
