package api

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/x-sushant-x/Zocket/config"
	repository "github.com/x-sushant-x/Zocket/repository/implementation"
	"github.com/x-sushant-x/Zocket/service"
	"github.com/x-sushant-x/Zocket/socket"
)

func StartServer() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://z-frontend-kappa.vercel.app",
		AllowHeaders:     "Authorization, Content-Type",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	webSocketClient := socket.NewWebSocketClient()

	api := app.Group("/api")

	userRepo := repository.NewUserRepo(config.DB)

	{
		authService := service.NewAuthService(userRepo)
		authController := NewAuthController(authService)
		auth := api.Group("/auth")
		auth.Post("/signup", authController.Signup)
		auth.Post("/login", authController.Login)
	}

	{
		userService := service.NewUserService(userRepo)
		userController := NewUserController(userService)
		user := api.Group("/user")
		user.Get("/list", userController.GetAllUsers)
	}

	{
		taskRepo := repository.NewTaskRepo(config.DB)
		taskServce := service.NewTaskService(taskRepo, webSocketClient)
		taskController := NewTaskController(taskServce)

		task := api.Group("/task")
		task.Post("/", taskController.CreateTask)
		task.Get("/list", taskController.GetAllTasks)
		task.Put("/status", taskController.UpdateTaskStatus)
	}

	app.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		webSocketClient.Add(conn)

		defer webSocketClient.Remove(conn)

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}))

	log.Fatal(app.Listen(":4000"))
}
