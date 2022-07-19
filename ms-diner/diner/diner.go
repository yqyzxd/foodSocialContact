package diner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"foodSocialContact/ms-diner/dao"
	"foodSocialContact/ms-diner/domain"
	dinerpb "foodSocialContact/ms-diner/proto/gen"
	oauthpb "foodSocialContact/ms-oauth/proto/gen"
	appcodes "foodSocialContact/shared/codes"
	"foodSocialContact/shared/times"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type OAuth2Request struct {
	account  string
	password string

	clientID     string //client_id
	clientSecret string //client_secret
	grantType    string //表示授权类型，密码模式此处的值固定为"password"，必选项。
	scope        string //表示权限范围，可选项。 api / all

}

type Service struct {
	dinerpb.UnimplementedDinerServiceServer
	OAuthServiceClient oauthpb.OAuthServiceClient
	Dao                *dao.DinerDao
	Redis              *redis.Client
}

func (s *Service) GetUserByUsername(c context.Context, req *dinerpb.GetDinerRequest) (*dinerpb.DinerEntity, error) {

	diner, err := s.Dao.FindUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	return &dinerpb.DinerEntity{
		Id:    diner.ID,
		Diner: diner.Diner,
	}, nil

}

// SignIn 登录
func (s *Service) SignIn(account, password string) (*domain.OAuthDiner, error) {
	//向oauth2 发起登录http请求
	req := OAuth2Request{
		clientID:     "appId",
		clientSecret: "appSecret",
		grantType:    "password",
		scope:        "api",
		account:      account,
		password:     password,
	}

	srcUrl := "http://localhost:9096/oauth/token"
	url := fmt.Sprintf(srcUrl+"?grant_type=%s&client_id=%s&client_secret=%s&scope=%s&username=%s&password=%s", req.grantType,
		req.clientID, req.clientSecret, req.scope, req.account, req.password)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body := resp.Body
	defer body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	b, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}
	var diner domain.OAuthDiner
	err = json.Unmarshal(b, &diner)
	if err != nil {
		return nil, err
	}
	return &diner, nil
}

// Logout 向oauth2发起请求 清除token
func (s *Service) Logout(token string) error {

	res, err := s.OAuthServiceClient.RemoveToken(context.Background(), &oauthpb.OAuthRemoveTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return err
	}
	if res.Code != int32(appcodes.SUCCESS.Int()) {
		return errors.New(res.Msg)
	}

	return nil
}

func (s *Service) Sign(dinerID int, dateStr string) (int, error) {
	//获取日期 不传默认是当前日期
	date, err := getDate(dateStr)
	if err != nil {
		return 0, status.Error(codes.InvalidArgument, "")
	}
	//获取日期对应的天数，多少号
	offset := date.Day() - 1
	//签到key的格式user:sign:5:yyyyMM
	key := buildSignKey(dinerID, &date)

	//查看是否已经签到
	cmd := s.Redis.GetBit(context.Background(), key, int64(offset))
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	signed := cmd.Val() == 1
	if signed {
		return 0, status.Error(codes.AlreadyExists, "")
	}
	//签到
	s.Redis.SetBit(context.Background(), key, int64(offset), 1)

	//统计连续签到的次数
	count, err := s.getContinuousSignCount(dinerID, &date)
	if err != nil {
		return 0, nil
	}
	return count, nil
}

func getDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Now(), nil
	}
	return times.Parse(dateStr)

}

func buildSignKey(uid int, date *time.Time) string {
	dateStr := times.Format(date, "")
	yyyyMmDd := strings.Split(dateStr, " ")[0]
	parts := strings.Split(yyyyMmDd, "-")
	yyyyMM := fmt.Sprintf("%s%s", parts[0], parts[1])

	return fmt.Sprintf("user:sign:%d:%s", uid, yyyyMM)
}

//连续签到次数
func (s *Service) getContinuousSignCount(uid int, date *time.Time) (int, error) {
	//获取日期对应的天数，多少号

	dayOfMonth := date.Day()
	key := buildSignKey(uid, date)
	args := fmt.Sprintf("get u%d 0", dayOfMonth)
	cmd := s.Redis.BitField(context.Background(), key, args)
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	v := cmd.Val()[0]
	signedCount := 0
	for i := dayOfMonth; i > 0; i-- {
		//右移再左移 如果等于自己说明最低位是0，表示未签到，否则表示标签
		if v>>1<<1 == v {
			//低位0而且非当天说明连续签到中断了
			if i != dayOfMonth {
				break
			}
		} else {
			signedCount++
		}
		//右移以为并重新赋值，相当于把最低位丢弃一位
		v = v >> 1

	}
	return signedCount, nil
}

//获取用户本月签到次数
func (s *Service) GetSignCount(dinerID int, dateStr string) (int, error) {

	date, err := getDate(dateStr)
	if err != nil {
		return 0, status.Error(codes.InvalidArgument, "")
	}
	key := buildSignKey(dinerID, &date)

	//bitcount user:sign:5:202207
	cmd := s.Redis.BitCount(context.Background(), key, nil)
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}

	return int(cmd.Val()), nil
}

// GetSignInfoOfMonth 获得当月签到情况
func (s *Service) GetSignInfoOfMonth(dinerID int, dateStr string) (map[string]bool, error) {
	date, err := getDate(dateStr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}
	key := buildSignKey(dinerID, &date)
	//获取某月有多少天

	year, _ := strconv.Atoi(times.Format(&date, times.YYYY_LAYOUT))
	month, _ := strconv.Atoi(times.Format(&date, times.MM_LAYOUT))
	dayOfMonth := times.LengthOfMonth(year, month)
	args := fmt.Sprintf("get u%d 0", dayOfMonth)
	cmd := s.Redis.BitField(context.Background(), key, args)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	v := cmd.Val()[0]

	m := make(map[string]bool)
	yyyy_MM := times.Format(&date, times.YYYY_MM_LAYOUT)
	//从低位到高位进行遍历，为0表示未签到，为1表示已签到
	for i := dayOfMonth; i > 0; i-- {
		//签到   yyyy-MM-dd true
		//未签到   yyyy-MM-dd false
		signed := v>>1<<1 != v
		var dd string
		if dayOfMonth < 10 {
			dd = fmt.Sprintf("0%d", dayOfMonth)
		} else {
			dd = strconv.Itoa(dayOfMonth)
		}
		key := fmt.Sprintf("%s-%s", yyyy_MM, dd)
		m[key] = signed
		v = v >> 1
	}

	return m, nil
}
