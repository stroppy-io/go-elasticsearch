// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package gentests

import "testing"

func TestExpand_StashedValues(t *testing.T) {
	for in, want := range map[string]string{
		// A stashed value as the first segment is the root of the path.
		"$profile.uid":                 `getKey(stash["$profile"], "uid")`,
		"$profile.user.username":       `getKey(getKey(stash["$profile"], "user"), "username")`,
		"$errors.details.missing.type": `getKey(getKey(getKey(stash["$errors"], "details"), "missing"), "type")`,
		// A stashed value later in the path names a key.
		"nodes.$master.modules": `getKey(getKey(getKey(mapi, "nodes"), fmt.Sprintf("%v", stash["$master"])), "modules")`,
		"user.username":         `getKey(getKey(mapi, "user"), "username")`,
	} {
		if got := expand(in); got != want {
			t.Errorf("expand(%q) = %s, want %s", in, got, want)
		}
	}
}
