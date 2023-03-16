package jwtmiddleware

import (
	"context"
	"dousheng/common/errno"
	"dousheng/router/biz/model/core"
	"dousheng/router/kitex_gen/auth"
	"dousheng/router/rpc"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

var (
	AuthMiddleware    *jwt.HertzJWTMiddleware
	successStatusCode int32 = 0
	successStatusMsg        = "success"
)

const (
	SecretKey      = "secret key"
	UserIDKey      = "user_id"
	UsernameKey    = "username"
	LoginedKey     = "logined"
	MaxUsernameLen = 32
	MaxPasswordLen = 32
)

func Init() {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		// Key:        []byte(constants.SecretKey),
		Key:        []byte(SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		// 将验证后返回的内容装入jwt
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*auth.UserAuth); ok && v != nil {
				return jwt.MapClaims{
					UserIDKey:   strconv.FormatInt(v.UserId, 10),
					UsernameKey: v.Username,
					LoginedKey:  true,
				}
			}
			return jwt.MapClaims{LoginedKey: false}
		},
		// 为了在错误时返回错误信息,是unauthorized的message参数
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			// c.JSON(consts.StatusOK, map[string]interface{}{
			// 	"code":   errno.SuccessCode,
			// 	"token":  token,
			// 	"expire": expire.Format(time.RFC3339),
			// })
			resp := new(core.DouyinUserLoginResponse)
			resp.Token = &token
			resp.StatusCode = &successStatusCode
			resp.StatusMsg = &successStatusMsg
			tk, _ := AuthMiddleware.ParseTokenString(token)
			claims := jwt.ExtractClaimsFromToken(tk)
			uid, _ := strconv.ParseInt(claims[UserIDKey].(string), 10, 64)
			resp.UserId = &uid
			c.JSON(consts.StatusOK, resp)
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			// c.JSON(code, map[string]interface{}{
			// 	"code":    errno.AuthorizationFailedErrCode,
			// 	"message": message,
			// })
			resp := new(core.DouyinUserLoginResponse)
			var errCode int32 = errno.AuthorizationFailedErrCode
			resp.StatusCode = &errCode
			resp.StatusMsg = &message
			c.JSON(code, resp)
		},
		// 用来进行登录验证并返回认证信息
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var loginVar core.DouyinUserLoginRequest
			err = c.BindAndValidate(&loginVar)
			var res *auth.UserAuth
			if err != nil {
				return res, jwt.ErrMissingLoginValues
			}
			ul, pl := len(*loginVar.Username), len(*loginVar.Password)
			if ul == 0 || pl == 0 || ul > MaxUsernameLen || pl > MaxPasswordLen {
				return res, jwt.ErrMissingLoginValues
			}

			res, err = rpc.Login(context.Background(), &auth.LoginRequest{Username: *loginVar.Username, Password: *loginVar.Password})
			return res, err
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		panic(err)
	}
}
