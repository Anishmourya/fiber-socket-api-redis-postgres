package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/websocket"
	"gofibe/fiber/db"
	"gofibe/fiber/routes"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	app := fiber.New()

	app.Settings.Prefork = true
	app.Settings.CaseSensitive = true
	app.Settings.StrictRouting = true
	app.Settings.ReadTimeout = 30 * time.Second
	app.Settings.TemplateFolder = "./templates"
	app.Settings.TemplateExtension = ".html"

	app.Use(logger.New())

	app.Use(helmet.New())

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	
	fmt.Println(client.Ping())

	app.Get("/", routes.Index)

	app.Get("/db_test", func(c *fiber.Ctx) {

		db.ExampleDB_Model()

	});


	// Optional middleware
	app.Use("/ws", func(c *fiber.Ctx) {
		if c.Get("host") == "localhost:3000" {
			c.Locals("Host", "Localhost:3000")
			c.Next()
			return
		}
		c.Status(403).Send("Request origin not allowed")
	});

	// Upgraded websocket request
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
       socket_id := fmt.Sprint("my_channel_",rand.Intn(100000))
		pubsub := client.Subscribe(socket_id)
		defer pubsub.Close()

		for {
			mt, msg, err := c.ReadMessage()

			if err := client.Publish(socket_id, msg).Err(); err != nil {
				panic(err)
			}

			if err != nil {
				log.Println("read:", err)
				break
			}
			pubsubmsg,err := pubsub.ReceiveMessage()

			b := []byte( pubsubmsg.Payload)

			if err != nil {
				panic(err)
			}

			err = c.WriteMessage(mt, b)
			if err != nil {
				log.Println("write:", err)
				break
			}

		}


	}))



	// listen address ws://localhost:3000/ws
	app.Listen(3000)
}
