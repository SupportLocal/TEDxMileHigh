package handlers

import (
	"log"
	"net/http"
	"supportlocal/TEDxMileHigh/redis"
)

func BlockMessages(w http.ResponseWriter, r *http.Request) {
	var err error
	var block struct {
		Ids []int `json:"ids"`
	}

	if err = readJson(r, &block); err != nil {
		log.Printf("website: readJson failed %q", err)
	}

	if len(block.Ids) > 0 {
		messageRepo := redis.MessageRepo()
		for _, id := range block.Ids {
			if err = messageRepo.Block(id); err != nil {
				log.Printf("website: messageRepo.Block failed %q", err)
			}
		}
	}

	mustWriteJson(w, block)
}
