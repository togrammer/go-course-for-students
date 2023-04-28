package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework9/internal/app"
)

func (s AdService) CreateAd(ctx context.Context, req *CreateAdRequest) (*AdResponse, error) {
	ad, err := s.a.CreateAd(ctx, req.Title, req.Text, req.UserId)
	if err != nil {
		return &AdResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &AdResponse{Id: ad.ID,
		Title:     ad.Title,
		Text:      ad.Text,
		AuthorId:  ad.AuthorID,
		Published: ad.Published,
	}, nil
}

func (s AdService) ChangeAdStatus(ctx context.Context, req *ChangeAdStatusRequest) (*AdResponse, error) {
	ad, err := s.a.ChangeAdStatus(ctx, req.AdId, req.UserId, req.Published)
	if err != nil {
		if errors.Is(err, app.ErrWrongUser) {
			return &AdResponse{}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &AdResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &AdResponse{Id: ad.ID,
		Title:     ad.Title,
		Text:      ad.Text,
		AuthorId:  ad.AuthorID,
		Published: ad.Published,
	}, nil
}

func (s AdService) UpdateAd(ctx context.Context, req *UpdateAdRequest) (*AdResponse, error) {
	ad, err := s.a.UpdateAd(ctx, req.AdId, req.UserId, req.Title, req.Text)
	if err != nil {
		if errors.Is(err, app.ErrWrongUser) {
			return &AdResponse{}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &AdResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &AdResponse{Id: ad.ID,
		Title:     ad.Title,
		Text:      ad.Text,
		AuthorId:  ad.AuthorID,
		Published: ad.Published}, nil
}

func (s AdService) ListAds(ctx context.Context, empty *emptypb.Empty) (*ListAdResponse, error) {
	ads := s.a.ListAds(ctx, -1, false, -1)
	resList := ListAdResponse{}
	for _, ad := range ads {
		resList.List = append(resList.List, &AdResponse{Id: ad.ID,
			Title:     ad.Title,
			Text:      ad.Text,
			AuthorId:  ad.AuthorID,
			Published: ad.Published,})
	}
	return &resList, nil
}

func (s AdService) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	usr := s.a.CreateUser(ctx, req.Name, string(""))
	return &UserResponse{Id: usr.ID, Name: usr.Nickname}, nil
}

func (s AdService) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
	usr, err := s.a.FindUser(ctx, req.Id)
	if err != nil {
		return &UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &UserResponse{
		Id:   usr.ID,
		Name: usr.Nickname,
	}, nil
}

func (s AdService) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s AdService) DeleteAd(context.Context, *DeleteAdRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
