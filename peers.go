package enamorecache

import pb "github.com/dmlvhh/enamorecahce/enamorecahcepb"

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}
