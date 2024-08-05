package api

import (
	"context"
	"crypto/rand"
<<<<<<< HEAD
=======
	"crypto/sha256"
	"encoding/json"
>>>>>>> main
	"fmt"
	"main/backend/api/auth"
	"main/backend/config"
	"main/backend/store"
	v1pb "main/proto/generated-go/api/v1"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type AuthService struct {
	v1pb.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, r *v1pb.LoginRequest) (*v1pb.LoginResponse, error) {
	// Make sure that the username and password match.
	user, err := store.GetUserByName(r.Name)
	if err != nil {
		return nil, err
	}

	if user.PasswdHash != auth.Sha256(r.Passwd, user.Salt) {
		return nil, errors.New("Wrong username or password")
	}

	token, err := auth.GenerateAccessToken(user.Id, config.PrivateKey)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to generate access token")
	}

	sessionId := uuid.NewString()
	if err := store.SetSession(ctx, strconv.Itoa(user.Id), []byte(sessionId)); err != nil {
		return nil, err
	}

	return &v1pb.LoginResponse{
		Token:     token,
		SessionId: sessionId,
	}, nil
}

func (s *AuthService) Register(ctx context.Context, r *v1pb.RegisterRequest) (*v1pb.RegisterResponse, error) {
	// check if the user exists.
	user, err := store.GetUserByName(r.Name)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("The username already exists")
	}

	salt := make([]byte, 16)
	_, _ = rand.Read(salt)
	saltString := fmt.Sprintf("%x", salt)
<<<<<<< HEAD
	hashedPasswd := auth.Sha256(r.Passwd, saltString)
	err = store.CreateUser(r.Name, hashedPasswd, saltString, r.Email, r.Phone, store.UserTypeRegular)
=======
	passwordHash := hashPassword(r.Passwd, saltString)
	err = store.CreateUser(r.Name, passwordHash, saltString, r.Email, r.Phone, store.UserTypeRegular)
>>>>>>> main
	if err != nil {
		return nil, errors.Wrap(err, "failed to register")
	}

	return &v1pb.RegisterResponse{}, nil
}

func (s *AuthService) Logout(ctx context.Context, r *v1pb.LogoutRequest) (*v1pb.LogoutResponse, error) {
	delSessionErr := store.DelSession(ctx, fmt.Sprint(r.Id))
	deactivateTokenErr := store.DeactivateAccessToken(ctx, r.Token)
	if err := multierr.Combine(delSessionErr, deactivateTokenErr); err != nil {
		return nil, err
	}
	return &v1pb.LogoutResponse{}, nil
}
