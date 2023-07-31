package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	dnsapi "github.com/AlexS778/fqdnIPLookup/internal/api/dns"
	"github.com/AlexS778/fqdnIPLookup/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "Port number for the server")

	flag.Parse()

	if port == 0 {
		fmt.Println("Missing port flag -port, server is running at :8080")
	}

	connStr := os.Getenv("connStr")
	if connStr == "" {
		log.Fatal("No database connection string, set it as env variable: export connStr=")
	}
	waitTime := os.Getenv("waitTime")
	if waitTime == "" {
		log.Fatal("Specify wait time for querying dns servers every x seconds")
	}
	dbHanlder := db.InitDBContext(connStr)
	router := gin.Default()
	dnsapi.RegiserV1Routes(router, dbHanlder)
	timeDuration, err := time.ParseDuration(waitTime)
	if err != nil {
		log.Fatal(err)
	}
	go dnsapi.ContinuousUpdate(dbHanlder, timeDuration)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow requests from any origin
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Run(fmt.Sprintf(":%d", port))
}
