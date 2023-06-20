package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

type MongoClient struct {
	Client *mongo.Client
}

// 定义单例模式获取mongo
var once sync.Once
var client *MongoClient

func NewMongo(uri string, pool int) *MongoClient {
	once.Do(func() {
		client = newMongo(uri, pool)
	})
	return client
}

func newMongo(uri string, pool int) *MongoClient {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 连接MongoDB
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(uri),
		options.Client().SetMinPoolSize(uint64(pool)),
	)
	if err != nil {
		panic(err)
	}

	// 检测MongoDB是否连接成功
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB连接成功！")
	return &MongoClient{
		Client: client,
	}
}

func (c *MongoClient) GetCollection(db, collection string) *mongo.Collection {
	// 获取numbers集合
	return c.Client.Database(db).Collection(collection)
}
