package main

import (
	"atest/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/func25/mongofunc/mongorely"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//helm install --dry-run --debug atest ./atestchart --set image.repo=gcr.io --set image.name=atest -f ./atestchart/values-base.yaml -f ./atestchart/values-staging.yaml
var redisClient *redis.Client
var mongoClient *mongo.Client

func main() {
	var err error
	if err := config.Init(); err != nil {
		// log.Fatal("cannot init config ", err)
		log.Println(err)
	}

	redisClient = connectRedis()
	defer redisClient.Close()

	if mongoClient, err = connectMongo(); err != nil {
		log.Println("cannot connect mongo")
	}

	go muddleRedis()
	go muddleMongo()

	startServer()

}

func connectRedis() *redis.Client {
	fmt.Println(config.EnvConfig)
	c := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.EnvConfig.RedisHost, "6379"),
		Password: config.EnvConfig.RedisPass,
	})

	return c
}

func connectMongo() (*mongo.Client, error) {
	cl, err := mongorely.Connect(context.Background(), mongorely.DbConfig{
		DbName:   "atest",
		UserName: "root",
		Password: config.EnvConfig.MongoPass,
		Host:     config.EnvConfig.MongoHost,
		Port:     "27017",
	})

	if err != nil {
		return nil, err
	}

	mod := mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		}, Options: nil,
	}
	_, err = cl.Database("atest").Collection(Hero{}.GetMongoCollName()).Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		logger.Error().Err(err).Msg("[mongo][index]: " + err.Error())
	}

	return cl, nil
}

func startServer() {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf(`i'm alive and my message is "%s"`, config.EnvConfig.Message)})
	})
	r.GET("/stress", func(c *gin.Context) {
		j := 0

		// max := 1000000
		// max, _ = mathfunc.RandInt(max, max*1000)
		// for i := 0; i < max; i++ {
		// 	j += i
		// }
		c.JSON(http.StatusOK, gin.H{"message": strconv.Itoa(j)})
	})
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": `ver 1`})
	})
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"private message": config.EnvConfig.PrivateMessage})
	})

	healthy := true
	// go func() {
	// 	time.Sleep(25 * time.Second)
	// 	healthy = false
	// }()
	ready := true
	// go func() {
	// 	time.Sleep(25 * time.Second)
	// 	ready = false
	// }()

	r.GET("/api/healthy", func(c *gin.Context) {
		if healthy {
			c.Status(http.StatusOK)
		} else {
			fmt.Println("--------- UNHEALTHY")
			c.Status(http.StatusBadGateway)
		}
	})

	r.GET("/api/ready", func(c *gin.Context) {
		if ready {
			c.Status(http.StatusOK)
		} else {
			fmt.Println("--------- UNREADY")
			c.Status(http.StatusBadGateway)
		}
	})

	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Writer.Header().Set("Connection", "keep-alive")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "404 not found"})
	})
	fmt.Println(config.EnvConfig)
	// if err := r.Run("0.0.0.0:80"); err != nil {
	if config.EnvConfig.Port == 0 {
		config.EnvConfig.Port = 80
	}
	if err := r.Run(fmt.Sprintf(":%d", config.EnvConfig.Port)); err != nil {
		log.Fatal("cannot run server", err)
	}
}
