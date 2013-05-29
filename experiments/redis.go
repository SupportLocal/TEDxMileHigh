package main

import (
	"github.com/vmihailenco/redis"
	"log"
	"strconv"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"sync"
)

func main() {

	client := redis.NewTCPClient(":6379", "", -1)
	defer client.Close()

	ping := client.Ping()
	log.Printf("%#v %q", ping.Err(), ping.Val())

	set := client.Set("key", "1")
	log.Printf("%#v %v", set.Err(), set.Val())

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(wg sync.WaitGroup) {
			multi, err := client.MultiClient()
			defer multi.Close()
			fatal.If(err)

			watch := multi.Watch("key")
			fatal.If(watch.Err())
			log.Printf("%v", watch.Val())

			reqs, err := transaction(multi)
			fatal.If(err)
			log.Println(reqs)

			wg.Done()
		}(wg)
	}

	wg.Wait()

	get := client.Get("key")
	log.Printf("%#v %v", get.Err(), get.Val())

	del := client.Del("key")
	log.Printf("%#v %v", del.Err(), del.Val())
}

func transaction(multi *redis.MultiClient) (reqs []redis.Req, err error) {

	get := multi.Get("key")
	err = get.Err()

	if err != nil && err != redis.Nil {
		return nil, err
	}

	var val int64
	val, err = strconv.ParseInt(get.Val(), 10, 64)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	reqs, err = multi.Exec(func() {
		multi.Set("key", strconv.FormatInt(val+1, 10))
	})

	if err == redis.Nil {
		log.Println("transaction failed. repeat.")
		return transaction(multi)
	}

	return reqs, err
}
