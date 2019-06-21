package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func info(host string, port int, dbNumber int) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       dbNumber,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	response, err := client.Info().Result()
	if err != nil {
		log.Fatalf("Errror: %s", err)
	}
	log.Println(response)
}

func main() {
	hostFlag := flag.String("h", "127.0.0.1", "Server hostname (default: 127.0.0.1)")
	portFlag := flag.Int("p", 6379, "Server port (default: 6379)")
	dbNumberFlag := flag.Int("n", 0, "Database number (default: 0)")
	flag.Parse()
	info(*hostFlag, *portFlag, *dbNumberFlag)
}
