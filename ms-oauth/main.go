package main

import (
	"context"
	dinerpb "foodSocialContact/ms-diner/proto/gen"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
)

func main() {
	manager := manage.NewDefaultManager()
	// token memory store
	//manager.MustTokenStorage(store.NewMemoryTokenStore())

	// use redis token store
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	}, "TOKEN:"))

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("appId", &models.Client{
		ID:     "appId",
		Secret: "appSecret",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect on port :8082")
		return
	}
	dinnerServiceClient := dinerpb.NewDinerServiceClient(conn)

	//通过username 鉴权 以返回token  调用diner.Service
	srv.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (userID string, err error) {
		d, err := dinnerServiceClient.GetUserByUsername(context.Background(), &dinerpb.GetDinerRequest{
			Username: username,
		})
		if err != nil {
			return "", err
		}
		return strconv.Itoa(int(d.Id)), nil
	})
	//todo 可以返回自定义的额外字段 默认只返回 access_token expires_in scope token_type
	srv.ExtensionFieldsHandler = func(ti oauth2.TokenInfo) (fieldsValue map[string]interface{}) {
		return nil
	}

	http.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))

}
