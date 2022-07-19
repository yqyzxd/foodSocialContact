package main

import (
	"fmt"
	"foodSocialContact/ms-follow/dao"
	"foodSocialContact/ms-follow/follow"
	followpb "foodSocialContact/ms-follow/proto/gen"
	"foodSocialContact/ms-follow/router"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	go grpcServe(logger)
	httpServe(logger)

}

func httpServe(logger *zap.Logger) {
	//注册http服务
	engine := gin.Default()
	router.Register(engine)
	logger.Sugar().Fatal(engine.Run(":8080"))
}

func grpcServe(logger *zap.Logger) {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("cannot listen on port 8083:%v", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	dbName := "food"
	dsn := fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{},
	})
	if err != nil {
		log.Fatalf("cannot open mysql:%v", err)
	}

	s := grpc.NewServer()
	followpb.RegisterFollowServiceServer(s, &follow.Service{
		Dao: &dao.FollowDao{
			DB: db.Table("t_follow"),
		},
		Redis: rdb,
	})
	logger.Sugar().Fatal(s.Serve(lis))
}
