// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package cluster

import (
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
)

func TestPeer(t *testing.T) {
	a := assertions.New(t)

	conn := new(grpc.ClientConn)

	p := &peer{
		name:   "name",
		roles:  []ttnpb.ClusterRole{ttnpb.ClusterRole_ACCESS},
		target: "target",
		conn:   conn,
	}

	a.So(p.HasRole(ttnpb.ClusterRole_APPLICATION_SERVER), should.BeFalse)
	a.So(p.HasRole(ttnpb.ClusterRole_ACCESS), should.BeTrue)

	a.So(p.Conn(), should.Equal, conn)
}
