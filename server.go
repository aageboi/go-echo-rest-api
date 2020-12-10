package main

import (
	"flag"
	"fmt"
	"os"
	"log"

	"github.com/aageboi/go-echo-rest-api/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/joho/godotenv"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
)

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func main() {
	env := flag.String("env", "local", "Input either local or production")
	port := flag.String("port", "3000", "Input port number")

	flag.Parse()
	fmt.Println("env", *env)
	fmt.Println("port", *port)

	e := echo.New()

	e.Use(middleware.Logger())

	// Redis connection
	rds := redis.NewClient(&redis.Options{
		Addr:     goDotEnvVariable("REDIS_HOST"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rds.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}

	// Database connection
	db, err := mgo.Dial(goDotEnvVariable("MONGO_HOST"))
	if err != nil {
		fmt.Println(db, err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db, REDIS: rds}

	// Routes
	e.GET("/articles/:id", h.FindArticleByID)
	e.GET("/articles", h.FindAllArticle)

	e.Use(ServerHeader)
	e.Logger.Fatal(e.Start(":" + *port))
}
