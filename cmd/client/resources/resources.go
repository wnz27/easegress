/*
 * Copyright (c) 2017, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package resources provides the resources utilities for the client.
package resources

import (
	"fmt"

	"github.com/megaease/easegress/v2/cmd/client/general"
)

// GetResourceKind returns the kind of the resource.
func GetResourceKind(arg string) (string, error) {
	if general.InAPIResource(arg, CustomData()) {
		return CustomData().Kind, nil
	}
	if general.InAPIResource(arg, CustomDataKind()) {
		return CustomDataKind().Kind, nil
	}
	if general.InAPIResource(arg, Member()) {
		return Member().Kind, nil
	}
	objects, err := ObjectAPIResources()
	if err != nil {
		return "", err
	}
	for _, object := range objects {
		if general.InAPIResource(arg, object) {
			return object.Kind, nil
		}
	}
	return "", fmt.Errorf("unknown resource: %s", arg)
}
