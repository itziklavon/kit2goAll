package mongo

import (
	"log"
	"time"

	"../configuration"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoAdress = []string{configuration.GetMongoPropertyValue("MONGODB_URL_1"), configuration.GetMongoPropertyValue("MONGODB_URL_2"), configuration.GetMongoPropertyValue("MONGODB_URL_3")}
var username = configuration.GetMongoPropertyValue("MONGODB_USERNAME")
var password = configuration.GetMongoPropertyValue("MONGODB_PASS")

func GetMongoConnection(dbName string) *mgo.Session {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    mongoAdress,
		Timeout:  60 * time.Second,
		Database: dbName,
		Username: username,
		Password: password,
	}
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal(err)
	}
	return mongoSession
}

func Find(dbName string, collection string) bson.M {
	var m bson.M
	mongoSession := GetMongoConnection(dbName)

	coll := mongoSession.DB(dbName).C(collection)
	err := coll.Find(nil).One(&m)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoSession.Close()
	return m
}
