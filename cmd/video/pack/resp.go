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
	"errors"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/errno"
	"time"
)

// 创建基本返回，通过封装统一错误
func BuildBaseResp(err error) *video.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	//如果生成错误已经是封装的错误类型，则直接返回
	if errors.As(err, &e) {
		//调用生成BaseResp的接口，传入错误信息
		return baseResp(e)
	}
	//如果不是则直接调用默认服务错误类型配合对应的错误信息进行封装
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *video.BaseResp {
	//返回三个信息，错误代码，错误信息，服务执行时间戳
	return &video.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
