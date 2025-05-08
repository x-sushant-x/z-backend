package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/x-sushant-x/Zocket/config"
	"github.com/x-sushant-x/Zocket/controller"
	"github.com/x-sushant-x/Zocket/model"
	repository "github.com/x-sushant-x/Zocket/repository/implementation"
	"github.com/x-sushant-x/Zocket/service"
	"github.com/x-sushant-x/Zocket/socket"
)

func init() {
	godotenv.Load()
}

func main() {
	app := fiber.New()

	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{})
	config.DB.AutoMigrate(&model.Task{})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://z-frontend-kappa.vercel.app",
		AllowHeaders:     "*",
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
		authController := controller.NewAuthController(authService)
		auth := api.Group("/auth")
		auth.Post("/signup", authController.Signup)
		auth.Post("/login", authController.Login)
	}

	{
		userService := service.NewUserService(userRepo)
		userController := controller.NewUserController(userService)
		user := api.Group("/user")
		user.Get("/list", userController.GetAllUsers)
	}

	{
		taskRepo := repository.NewTaskRepo(config.DB)
		taskServce := service.NewTaskService(taskRepo, webSocketClient)
		taskController := controller.NewTaskController(taskServce)

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
