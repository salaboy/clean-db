package function

import (
	"context"
	"fmt"
	"net/http"
	"github.com/go-redis/redis"
	"os"
)

var redisHost = os.Getenv("REDIS_HOST")

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":6379",
		Password: "",
		DB:       0,
	})

	response, err := client.FlushDB().Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(res, "Status: " + response)

}
