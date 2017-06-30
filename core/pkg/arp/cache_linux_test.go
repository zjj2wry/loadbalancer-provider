/*
Copyright 2017 Caicloud authors. All rights reserved.

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

package arp

import "testing"

func Test_loadCache(t *testing.T) {
	tests := []struct {
		name    string
		want    Caches
		wantErr bool
	}{
		{
			"loadcache", nil, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := loadCache()
			if (err != nil) != tt.wantErr {
				t.Errorf("loadCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
