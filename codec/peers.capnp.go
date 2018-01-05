// Code generated by capnpc-go. DO NOT EDIT.

package codec

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Peers struct{ capnp.Struct }

// Peers_TypeID is the unique identifier for the type Peers.
const Peers_TypeID = 0xbd8afa1dc55b521a

func NewPeers(s *capnp.Segment) (Peers, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Peers{st}, err
}

func NewRootPeers(s *capnp.Segment) (Peers, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Peers{st}, err
}

func ReadRootPeers(msg *capnp.Message) (Peers, error) {
	root, err := msg.RootPtr()
	return Peers{root.Struct()}, err
}

func (s Peers) String() string {
	str, _ := text.Marshal(0xbd8afa1dc55b521a, s.Struct)
	return str
}

func (s Peers) Addresses() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s Peers) HasAddresses() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Peers) SetAddresses(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewAddresses sets the addresses field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Peers) NewAddresses(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Peers_List is a list of Peers.
type Peers_List struct{ capnp.List }

// NewPeers creates a new list of Peers.
func NewPeers_List(s *capnp.Segment, sz int32) (Peers_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Peers_List{l}, err
}

func (s Peers_List) At(i int) Peers { return Peers{s.List.Struct(i)} }

func (s Peers_List) Set(i int, v Peers) error { return s.List.SetStruct(i, v.Struct) }

func (s Peers_List) String() string {
	str, _ := text.MarshalList(0xbd8afa1dc55b521a, s.List)
	return str
}

// Peers_Promise is a wrapper for a Peers promised by a client call.
type Peers_Promise struct{ *capnp.Pipeline }

func (p Peers_Promise) Struct() (Peers, error) {
	s, err := p.Pipeline.Struct()
	return Peers{s}, err
}

const schema_b8fb51aaf7fc2d2f = "x\xda\x12X\xe3\xc0d\xc8Z\xcf\xc2\xc0\x10h\xc2\xca" +
	"\xf6_*(\xfa\xa8\xec\xaf\xae\xbd\x0c\x82<\x8c\xff\xf5" +
	"u\xff|_\x15\xf8{\x07\x03+#;\x03\x83\xb0*" +
	"\xd3$a]&\x10K\x93\xc9\x9e\xa1\xf5\x7fAjj" +
	"Q\xb1^r\"cA^\x81U@jj\x11cq" +
	"\x00#c \x0b3\x0b\x03\x03\x0b#\x03\x83 o\x10" +
	"\x03C \x0f3c\xa0\x06\x13\xe3\xff\xc4\x94\x94\xa2\xd4" +
	"\xe2\xe2T\x06\xc6bF>\x06\xc6\x00fFF\x1e\x06" +
	"&\x10\x13\x10\x00\x00\xff\xff\x06\x1b\x1eI"

func init() {
	schemas.Register(schema_b8fb51aaf7fc2d2f,
		0xbd8afa1dc55b521a)
}