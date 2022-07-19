package follow

import (
	"context"
	"fmt"
	"foodSocialContact/ms-follow/dao"
	followpb "foodSocialContact/ms-follow/proto/gen"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	followpb.UnimplementedFollowServiceServer
	Dao *dao.FollowDao

	Redis *redis.Client
}

// Follow 关注取关
func (s *Service) Follow(ctx context.Context, req *followpb.FollowRequest) (*followpb.FollowResponse, error) {
	//todo ctx中获取用户登录信息
	dinerID := int(req.DinerID)
	followDinerID := int(req.FollowDinerID)
	followStatus := int(req.Followed)

	rec, err := s.Dao.SelectFollow(dinerID, followDinerID)
	if err != nil {
		status := status.Convert(err)
		if status == nil || status.Code() != codes.NotFound {
			return nil, err
		}
	}

	//如果没有关注，而且要进行关注操作 --添加关注
	if rec == nil && req.Followed == 1 {
		//添加关注信息
		_, err := s.Dao.Save(dinerID, followDinerID, followStatus)
		if err != nil {
			return nil, status.Error(codes.Internal, "")
		}
		//添加关注列表到Redis
		s.addToRedisSet(dinerID, followDinerID)
	}
	//如果有关注记录，而且目前是关注状态，而且进行的是取关操作 --取消关注
	if rec != nil && rec.Valid == 1 && followStatus == 0 {

		err = s.Dao.Update(int(rec.ID), followStatus)
		if err != nil {
			return nil, status.Error(codes.Internal, "")
		}
		s.removeFromRedisSet(dinerID, followDinerID)

	}

	//如果有关注记录，而且目前是取关状态，进行的是关注操作 --重新关注
	if rec != nil && rec.Valid == 0 && followStatus == 1 {
		err = s.Dao.Update(int(rec.ID), followStatus)
		if err != nil {
			return nil, status.Error(codes.Internal, "")
		}
		s.addToRedisSet(dinerID, followDinerID)

	}

	return nil, nil
}

func (s *Service) removeFromRedisSet(dinerID int, followDinerID int) {
	followingKey := fmt.Sprintf("following:%d", dinerID)       //关注集合
	followersKey := fmt.Sprintf("followers:%d", followDinerID) //粉丝集合
	s.Redis.SAdd(context.Background(), followingKey, followDinerID)
	s.Redis.SAdd(context.Background(), followersKey, dinerID)
}

func (s *Service) addToRedisSet(dinerID int, followDinerID int) {
	followingKey := fmt.Sprintf("following:%d", dinerID)       //关注集合
	followersKey := fmt.Sprintf("followers:%d", followDinerID) //粉丝集合
	s.Redis.SRem(context.Background(), followingKey, followDinerID)
	s.Redis.SRem(context.Background(), followersKey, dinerID)
}

func (s *Service) IsFollow(ctx context.Context, in *followpb.IsFollowRequest) (*followpb.IsFollowResponse, error) {

	return nil, nil
}

func (s *Service) Following(ctx context.Context, in *followpb.FollowingRequest) (*followpb.FollowingResponse, error) {
	return nil, nil
}
func (s *Service) Followers(ctx context.Context, in *followpb.FollowersRequest) (*followpb.FollowersResponse, error) {
	return nil, nil
}
