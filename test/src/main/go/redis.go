package main

import (
	"./log"
	"./redis"
)

func main() {
	log.SetLogOutput("/redis/redis.log")
	log.Debug(":main: GS_BRAND_ID in redis is: " + redis.GetSysParam("GS_BRAND_ID"))
	redis.HSet("SysParams", "GS_BRAND_ID", "8")
	log.Debug(":main: GS_BRAND_ID in redis is: " + redis.GetSysParam("GS_BRAND_ID"))
	redis.HSet("SysParams", "GS_BRAND_ID", "7")
	keys := redis.Keys()
	for _, key := range keys {
		log.Debug(":main: keys in redis, under keys is: " + key)
	}
}
