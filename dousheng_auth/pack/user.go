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

package pack

import (
	"dousheng/auth/dal/db"
	"dousheng/auth/kitex_gen/auth"
)

// 转换model数据到展示数据
func User(u *db.User) *auth.UserAuth {
	if u == nil {
		return nil
	}

	return &auth.UserAuth{UserId: int64(u.ID), Username: u.UserName}
}
