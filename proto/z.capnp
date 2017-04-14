# Copyright (c) 2017 The Veltor Authors
#
# This file is part of Veltor.
#
# Veltor Network is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Veltor Network is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with Veltor Network.  If not, see <http://www.gnu.org/licenses/>.

using Go = import "/go.capnp";
$Go.package("proto");
$Go.import("proto");

using Ping = import "ping.capnp".Ping;
using Pong = import "pong.capnp".Pong;
using Discover = import "discover.capnp".Discover;
using Peers = import "peers.capnp".Peers;

@0x904d4f3f728c7f04;
struct Z {
	union {
		ping @0 :Ping;
		pong @1 :Pong;
		discover @2: Discover;
		peers @3: Peers;
		text @4: Text;
		data @5: Data;
	}
}
