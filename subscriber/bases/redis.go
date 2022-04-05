package bases

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

var ADDRREDIS = ""
var PASSWORDREDIS = ""

func SaveLogRedis(bjsonLog []byte) {
	fmt.Println("Go Redis Tutorial")
	client := redis.NewClient(&redis.Options{
		Addr:     ADDRREDIS,
		Password: PASSWORDREDIS,
		DB:       0,
	})

	json, err := json.Marshal(bjsonLog)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}
