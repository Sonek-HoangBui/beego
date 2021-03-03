// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package session

import (
	"github.com/Sonek-HoangBui/beego/v2/server/web/session"
)

// EncodeGob encode the obj to gob
func EncodeGob(obj map[interface{}]interface{}) ([]byte, error) {
	return session.EncodeGob(obj)
}

// DecodeGob decode data to map
func DecodeGob(encoded []byte) (map[interface{}]interface{}, error) {
	return session.DecodeGob(encoded)
}
