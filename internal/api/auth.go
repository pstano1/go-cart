package api

import (
	"errors"
	"strings"
	"time"

	"github.com/bcmills/unsafeslice"
	"github.com/gbrlsnchs/jwt"
	gojwt "github.com/golang-jwt/jwt"
	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	bearerToken = "Bearer"
)

type JWTConfig struct {
	User           *pkg.User `json:"user"`
	ExpirationTime int64     `json:"exp"`
	IssuedAt       int64     `json:"iat"`
}

func (c JWTConfig) Valid() error {
	if c.User.IsEmpty() {
		return errors.New("invalid token")
	}
	return nil
}

func (a *InstanceAPI) generateJWT(conf *JWTConfig) (string, error) {
	token := gojwt.New(gojwt.SigningMethodHS256)
	token.Claims = conf
	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		a.log.Error("Error in JWT token generation",
			zap.Error(err),
		)
		return "", err
	}
	return tokenString, nil
}

func (a *InstanceAPI) SignUserIn(credentials *pkg.Credentials) (*pkg.SignInResponse, error) {
	a.log.Debug("authenticating in progress",
		zap.String("username", credentials.Username),
	)
	users, err := a.GetUsers(&pkg.UserFilter{
		Username: credentials.Username,
	})
	if err != nil {
		a.log.Error(pkg.ErrRetrievingUsers.Error())
		return nil, err
	}
	if len(users) == 0 {
		a.log.Error(pkg.ErrUserNotFound.Error())
		return nil, pkg.ErrUserNotFound
	}
	user := users[0]
	providedPassword := unsafeslice.OfString(credentials.Password)
	userPassword := unsafeslice.OfString(user.Password)
	err = bcrypt.CompareHashAndPassword(userPassword, providedPassword)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	JWT, err := a.generateJWT(&JWTConfig{
		User:           &user,
		ExpirationTime: now.Unix() + 3600,
		IssuedAt:       now.Unix(),
	})
	if err != nil {
		return nil, err
	}
	return &pkg.SignInResponse{
		Username:    credentials.Username,
		Permissions: user.Permissions,
		Token:       JWT,
	}, nil
}

func (a *InstanceAPI) RefreshUserJWT(user *pkg.User) (*pkg.SignInResponse, error) {
	now := time.Now()
	jwtToken, err := a.generateJWT(&JWTConfig{
		User:           user,
		ExpirationTime: now.Unix() + 3600,
		IssuedAt:       now.Unix(),
	})
	if err != nil {
		return nil, err
	}
	return &pkg.SignInResponse{
		Username: user.Username,
		Token:    jwtToken,
	}, nil
}

func (a *InstanceAPI) GetUserFromRequest(ctx *fasthttp.RequestCtx) (*pkg.User, error) {
	header := string(ctx.Request.Header.Peek("Authorization"))
	if !strings.HasPrefix(header, bearerToken+" ") {
		a.log.Error("Invalid token in request")
		return nil, pkg.ErrInvalidToken
	}
	token := strings.TrimPrefix(header, bearerToken+" ")
	rawDecodeText, err := jwt.FromString(token)
	if err != nil {
		a.log.Error("Token invalid",
			zap.Error(err),
		)
		return nil, pkg.ErrUserUnauthorized
	}
	hs256 := jwt.HS256(string(a.secretKey))
	if err := rawDecodeText.Verify(hs256); err != nil {
		return nil, pkg.ErrInvalidToken
	}
	public := rawDecodeText.Public()
	users, err := a.GetUsers(&pkg.UserFilter{
		Username: public["username"].(string),
	})
	if err != nil || len(users) == 0 {
		return nil, pkg.ErrInvalidToken
	}
	return &users[0], nil
}
