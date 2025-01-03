package raftgroup

import "github.com/WuKongIM/WuKongIM/pkg/raft/types"

type ITransport interface {
	// Send 发送事件
	Send(r IRaft, event types.Event)
}
