package main

import (
	"RedisTest/redis"
	"fmt"
	r "github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go redis tutorial")

	rc := r.NewClient(&r.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	//redis.AddItem("/org1/api1/1.0.0","API1",rc)
	//redis.AddItem("/org123/api2/1.0.0","API2",rc)
	//redis.AddItem("/org1243/api3/1.0.0","API3",rc)
	//redis.AddItem("/org189/api4/1.0.0","API4",rc)
	//redis.AddItem("/org145/api5/1.0.0","API5",rc)
	//redis.AddItem("/org1345/api6/1.0.0","API6",rc)
	//redis.AddItem("/appl122/api7/1.0.0","API7",rc)
	//redis.AddItem("/appl223/api8/1.0.0","API8",rc)
	//redis.AddItem("/myo123/api9/1.0.0","API9",rc)
	//redis.AddItem("/myo321/api10/1.0.0","API10",rc)

	//fmt.Println(redis.GetItem("/myo321/api10/1.0.0", rc))

	//list, err := redis.SearchCacheKeys("/org*", rc)
	//if err != nil {
	//	fmt.Println("error")
	//	return
	//}
	//fmt.Println(list[0])

	redis.SearchCacheKeysByScan("/org*", rc)
}

