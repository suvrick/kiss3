package app

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/suvrick/kiss/internal/bot"
	"github.com/suvrick/kiss/internal/middlewares"
	"github.com/suvrick/kiss/internal/proxy"
	"github.com/suvrick/kiss/internal/user"
	"github.com/suvrick/kiss/pkg/db"
	"github.com/suvrick/kiss/pkg/db/client/postgres"
)

const (
	logFile = "app.log"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {

	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err.Error())
	}

	logPath := path.Join(dir, logFile)

	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	log.Println("[Run] initialize logger: OK")

	dbconfig := &db.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "suvrick",
		Password: "bl69unn",
		DBName:   "kissdb",
	}

	db, err := postgres.NewClient(dbconfig)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: log.Writer(),
	}))

	router.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("../ui/dist/ui/index.html")
		} else {
			c.File("../ui/dist/ui/" + path.Join(dir, file))
		}
	})

	router.Use(middlewares.CORSMiddleware())

	proxyRepository := proxy.NewProxyRepository(db)
	proxyService := proxy.NewProxyService(proxyRepository)
	proxyController := proxy.NewProxyController(proxyService)
	proxyController.Register(router)

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService)
	userController.Register(router)

	botRepository := bot.NewBotRepository(db)
	botService := bot.NewBotService(botRepository)
	botController := bot.NewBotController(botService)
	botController.Register(router)

	// taskRepository := repositories.NewTaskRepository(db)
	// taskService := services.NewTaskService(taskRepository)
	// taskController := controllers.NewTaskController(taskService)
	// taskController.Register(router)

	return router.Run(":8080")
}
