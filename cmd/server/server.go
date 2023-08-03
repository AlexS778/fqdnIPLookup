package main

import (
	"context"
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
		log.Println("Missing port flag -port, server is running at :8080")
	}

	connStr := os.Getenv("connStr")
	if connStr == "" {
		log.Fatal("No database connection string, set it as env variable: export connStr=")
	}
	waitTime := os.Getenv("waitTime")
	if waitTime == "" {
		log.Fatal("Specify wait time for querying dns servers every x seconds")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dbHanlder := db.InitDBContext(connStr)
	router := gin.Default()
	router.Use(cors.Default())
	dnsapi.RegiserV1Routes(router, dbHanlder)
	timeDuration, err := time.ParseDuration(waitTime)
	if err != nil {
		log.Fatal(err)
	}

	go dnsapi.ContinuousUpdate(ctx, dbHanlder, timeDuration)

	router.Run(fmt.Sprintf(":%d", port))
}
