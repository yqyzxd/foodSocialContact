package api

import (
	"foodSocialContact/ms-diner/diner"
	oauthpb "foodSocialContact/ms-oauth/proto/gen"
	"foodSocialContact/shared/codes"
	"foodSocialContact/shared/resp"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {

	account := c.Query("account")
	password := c.Query("password")

	s := diner.Service{}
	result, err := s.SignIn(account, password)

	httpCode := http.StatusOK
	code := codes.SUCCESS
	msg := ""
	if err != nil {
		code = codes.ERROR
		msg = "oops"
	}
	res := resp.AppResponse{
		Err:  code,
		Msg:  msg,
		Data: result,
	}
	c.JSON(httpCode, res)

}

func Logout(c *gin.Context) {
	//获取access_token
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		//todo 处理错误
		return
	}
	oauthServiceClient := oauthpb.NewOAuthServiceClient(conn)

	accessToken := c.GetString("access_token")
	s := diner.Service{
		OAuthServiceClient: oauthServiceClient,
	}
	s.Logout(accessToken)
}

func Sign(c *gin.Context) {

	//todo 从token中获取用户dinerID
	dinerID := 1

	date := c.PostForm("date")
	s := diner.Service{}
	count, err := s.Sign(dinerID, date)

	//httpCode := http.StatusOK
	code := codes.SUCCESS
	msg := ""
	if err != nil {
		code = codes.ERROR
		msg = "oops"
	}
	res := resp.AppResponse{
		Err:  code,
		Msg:  msg,
		Data: count,
	}

	res.JSON(c, err)

	//c.JSON(httpCode, res)

}

func GetSignCount(c *gin.Context) {
	dinerID := 1
	date := c.PostForm("date")
	s := diner.Service{}
	s.GetSignCount(dinerID, date)
}

func GetSignInfoOfMonth(c *gin.Context) {
	dinerID := 1
	date := c.PostForm("date")
	s := diner.Service{}
	m, err := s.GetSignInfoOfMonth(dinerID, date)
	code := codes.SUCCESS
	msg := ""
	if err != nil {
		code = codes.ERROR
		msg = "oops"
	}
	res := resp.AppResponse{
		Err:  code,
		Msg:  msg,
		Data: m,
	}

	res.JSON(c, err)
}

func updateDinerLocation(c *gin.Context) {

	lon, err := strconv.ParseFloat(c.PostForm("lon"), 64)
	if err != nil {
		//error

		return
	}
	lat, err := strconv.ParseFloat(c.PostForm("lat"), 64)
	if err != nil {
		//error

		return
	}

	s := &diner.NearMeService{
		//todo 传入Redis

	}
	dinerID := 1
	s.UpdateDinerLocation(dinerID, lon, lat)

}
