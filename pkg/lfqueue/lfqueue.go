// Author: https://github.com/antigloss

// Package queue offers goroutine-safe Queue implementations such as LockfreeQueue(Lock free queue).
package lfqueue

import (
	"sync/atomic"
	"unsafe"
)

// NewLockfreeQueue is the only way to get a new, ready-to-use LockfreeQueue.
//
// Example:
//
//   lfq := queue.NewLockfreeQueue()
//   lfq.Push(100)
//   v := lfq.Pop()
func NewLockfreeQueue() *LockfreeQueue {
	var lfq LockfreeQueue
	lfq.head = unsafe.Pointer(&lfq.dummy)
	lfq.tail = lfq.head
	return &lfq
}

// LockfreeQueue is a goroutine-safe Queue implementation.
// The overall performance of LockfreeQueue is much better than List+Mutex(standard package).
type LockfreeQueue struct {
	head    unsafe.Pointer
	tail    unsafe.Pointer
	dummy   lfqNode
	queue   uint64
	dequeue uint64
}

// Pop returns (and removes) an element from the front of the queue, or nil if the queue is empty.
// It performs about 100% better than list.List.Front() and list.List.Remove() with sync.Mutex.
func (lfq *LockfreeQueue) Pop() interface{} {
	var dequeue uint64
	for {
		dequeue = atomic.LoadUint64(&lfq.dequeue)
		h := atomic.LoadPointer(&lfq.head)
		rh := (*lfqNode)(h)
		n := (*lfqNode)(atomic.LoadPointer(&rh.next))
		if n != nil {
			if atomic.CompareAndSwapPointer(&lfq.head, h, rh.next) {
				atomic.StoreUint64(&lfq.dequeue, dequeue+1)
				return n.val
			} else {
				continue
			}
		} else {
			return nil
		}
	}
}

// Push inserts an element to the back of the queue.
// It performs exactly the same as list.List.PushBack() with sync.Mutex.
func (lfq *LockfreeQueue) Push(val interface{}) {
	var queue uint64
	node := unsafe.Pointer(&lfqNode{val: val})
	for {
		queue = atomic.LoadUint64(&lfq.queue)
		t := atomic.LoadPointer(&lfq.tail)
		rt := (*lfqNode)(t)
		if atomic.CompareAndSwapPointer(&rt.next, nil, node) {
			// It'll be a dead loop if atomic.StorePointer() is used.
			// Don't know why.
			// atomic.StorePointer(&lfq.tail, node)
			atomic.CompareAndSwapPointer(&lfq.tail, t, node)
			atomic.StoreUint64(&lfq.queue, queue+1)
			return
		} else {
			continue
		}
	}
}

func (lfq *LockfreeQueue) Size() uint64 {
	return atomic.LoadUint64(&lfq.queue) - atomic.LoadUint64(&lfq.dequeue)
}

func (lfq *LockfreeQueue) IsEmpty() bool {
	return (atomic.LoadUint64(&lfq.queue) - atomic.LoadUint64(&lfq.dequeue)) == 0
}

type lfqNode struct {
	val  interface{}
	next unsafe.Pointer
}
