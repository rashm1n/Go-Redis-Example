package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

// AddItem adds an item to a certain key
func AddItem(key string, value string, client *redis.Client)  {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

// GetItem gets an item using a key
func GetItem(key string, client *redis.Client) string {
	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// SearchCacheKeys Searches for keys using KEYS pattern
func SearchCacheKeys(substring string, client *redis.Client) ([]string, error){
	var e error
	var keyList []string

	res := client.Keys(substring)

	if res.Err() != nil {
		fmt.Println("Error while searching keys for substring " + substring, res.Err().Error())
	} else {
		keyList, e = res.Result()
		if e == nil {
			return keyList,e
		}

		return keyList,e
	}
	return keyList,res.Err()
}

// SearchCacheKeysByScan searches for keys using SCAN command
func SearchCacheKeysByScan(substring string, client *redis.Client)  {
	iter := client.Scan(0, substring, 0).Iterator()
	for iter.Next() {
		fmt.Println(iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}
