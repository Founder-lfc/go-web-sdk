package config

import (
	"github.com/Founder-lfc/go-web-core/storage"
	"github.com/Founder-lfc/go-web-core/storage/queue"
)

type Queue struct {
	Memory *QueueMemory
}

type QueueMemory struct {
	PoolSize uint
}

var QueueConfig = new(Queue)

// Empty 空设置
func (e Queue) Empty() bool {
	return e.Memory == nil
}

// Setup 启用顺序  其他 > memory
func (e Queue) Setup() (storage.AdapterQueue, error) {
	return queue.NewMemory(e.Memory.PoolSize), nil
}
