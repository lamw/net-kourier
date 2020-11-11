/*
Copyright 2020 The Knative Authors

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

package envoy

import (
	"testing"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"
)

func TestNewWeightedCluster(t *testing.T) {
	got := NewWeightedCluster("test", 50, map[string]string{
		"foo": "bar",
	})
	want := &route.WeightedCluster_ClusterWeight{
		Name: "test",
		Weight: &wrappers.UInt32Value{
			Value: 50,
		},
		RequestHeadersToAdd: []*core.HeaderValueOption{{
			Header: &core.HeaderValue{
				Key:   "foo",
				Value: "bar",
			},
			Append: &wrappers.BoolValue{
				Value: false,
			},
		}},
	}

	assert.DeepEqual(t, got, want, cmpopts.IgnoreUnexported(wrappers.BoolValue{}, wrappers.UInt32Value{}))
}