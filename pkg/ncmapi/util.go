/*
Copyright 2023 Nokia

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ncmapi

import (
	"fmt"
	"net/url"
)

func GetPathFromCertHref(certHref string) (string, error) {
	parsedURL, err := url.Parse(certHref)
	if err != nil {
		return "", fmt.Errorf("cannot parsed given href: %s", certHref)
	}
	return parsedURL.Path, nil
}
