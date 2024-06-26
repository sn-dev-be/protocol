// Copyright © 2023 OpenIM. All rights reserved.
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

package sdkws

import (
	"errors"
)

func (x *MsgData) Check() error {
	if x.SendID == "" {
		return errors.New("sendID is empty")
	}
	if x.Content == nil {
		return errors.New("content is empty")
	}
	return nil
}

func (x *SignalVoiceReq) Check() error {
	if x.InviteUserID == "" {
		return errors.New("inviteUserID empty")
	}
	return nil
}
