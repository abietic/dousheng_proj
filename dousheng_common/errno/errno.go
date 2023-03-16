// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                       = 0
	ServiceErrCode                    = 10001
	ParamErrCode                      = 10002
	UserAlreadyExistErrCode           = 10003
	AuthorizationFailedErrCode        = 10004
	UserNotExistErrCode               = 10005
	RegisterErrCode                   = 10006
	GetUserStatsErrCode               = 10007
	GetUserFollowedErrCode            = 10008
	GetUserDataErrCode                = 10009
	GetVideoInfoListNotConsistErrCode = 10010
	GetVideoDataListErrCode           = 10011
	GetVideoStatsListErrCode          = 10012
	GetUserFavoriteVideoListErrCode   = 10013
	SaveVideoFileErrCode              = 10014
	PublishVideoErrCode               = 10015
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                       = NewErrNo(SuccessCode, "Success")
	ServiceErr                    = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr                      = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr           = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr        = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserNotExistErr               = NewErrNo(UserNotExistErrCode, "User not exist")
	RegisterErr                   = NewErrNo(RegisterErrCode, "User register failed")
	GetUserStatsErr               = NewErrNo(GetUserStatsErrCode, "Cannot get user stats")
	GetUserFollowedErr            = NewErrNo(GetUserFollowedErrCode, "Cannot get user followed user")
	GetUserDataErr                = NewErrNo(GetUserDataErrCode, "Cannot get user data")
	GetVideoInfoListNotConsistErr = NewErrNo(GetVideoInfoListNotConsistErrCode, "Data not consist on remote store")
	GetVideoDataListErr           = NewErrNo(GetVideoDataListErrCode, "Get user's published video data failed")
	GetVideoStatsListErr          = NewErrNo(GetVideoStatsListErrCode, "Get user's published video stats failed")
	GetUserFavoriteVideoListErr   = NewErrNo(GetUserFavoriteVideoListErrCode, "Get user's favorite videos failed")
	SaveVideoFileErr              = NewErrNo(SaveVideoFileErrCode, "Save video file failed")
	PublishVideoErr               = NewErrNo(PublishVideoErrCode, "Publish video failed")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
