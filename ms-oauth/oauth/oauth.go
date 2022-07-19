package oauth

import (
	"context"
	oauthpb "foodSocialContact/ms-oauth/proto/gen"
)

type Service struct {
	oauthpb.UnimplementedOAuthServiceServer
}

func (s *Service) Verify(context.Context, *oauthpb.OAuthTokenVerifyRequest) (*oauthpb.OAuthTokenVerifyResponse, error) {

	return nil, nil
}
func (s *Service) RemoveToken(context.Context, *oauthpb.OAuthRemoveTokenRequest) (*oauthpb.OAuthRemoveTokenResponse, error) {

	return nil, nil
}
