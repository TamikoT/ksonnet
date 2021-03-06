// Copyright 2017 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package kubecfg

import (
	"os"

	"github.com/ksonnet/ksonnet/metadata"
)

func manager() (metadata.Manager, error) {
	appDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	appRoot := metadata.AbsPath(appDir)

	return metadata.Find(appRoot)
}
