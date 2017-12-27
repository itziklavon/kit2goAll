package main

import (
	"./http"
	"./log"
)

func main() {
	log.SetLogOutput("/http-sender/sender.log")
	values := map[string]string{"user_name": "qqtst_y824", "password": "android100", "language": "en", "brand_id": "1"}
	http.POST("http://172.17.30.17:8080/player/1.2.2/player/login", values, nil)
}
