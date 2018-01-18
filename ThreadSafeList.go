// Package ThreadSafeList provides a simple wrapper to go's container/list with the addition af a mutex
package ThreadSafeList

import (
	"container/list"
	"sync"
)

// ThreadSafeList represents a thread safe doubly linked list
type ThreadSafeList struct {
	list      *list.List
	lockMutex sync.Mutex
}

func New() *ThreadSafeList {
	return &ThreadSafeList{list.New(), sync.Mutex{}}
}

func (l *ThreadSafeList) Back() *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.list.Back()
}

func (l *ThreadSafeList) Front() *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.list.Front()
}

func (l *ThreadSafeList) Init() *ThreadSafeList {
	l.Lock()
	defer l.Unlock()
	l.list.Init()
	return l
}

func (l *ThreadSafeList) InsertAfter(v interface{}, mark *list.Element) *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.list.InsertAfter(v, mark)
}

func (l *ThreadSafeList) InsertBefore(v interface{}, mark *list.Element) *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.list.InsertBefore(v, mark)
}

func (l *ThreadSafeList) Len() int {
	l.Lock()
	defer l.Unlock()
	return l.list.Len()
}

func (l *ThreadSafeList) MoveAfter(e, mark *list.Element) {
	l.Lock()
	defer l.Unlock()
	l.list.MoveAfter(e, mark)
}

func (l *ThreadSafeList) MoveBefore(e, mark *list.Element) {
	l.Lock()
	defer l.Unlock()
	l.list.MoveBefore(e, mark)
}

func (l *ThreadSafeList) MoveToBack(e *list.Element) {
	l.Lock()
	defer l.Unlock()
	l.list.MoveToBack(e)
}

func (l *ThreadSafeList) MoveToFront(e *list.Element) {
	l.Lock()
	defer l.Unlock()
	l.list.MoveToFront(e)
}

func (l *ThreadSafeList) PushBack(v interface{}) *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.list.PushBack(v)
}

func (l *ThreadSafeList) PushBackList(other *list.List) {
	l.Lock()
	defer l.Unlock()
	l.list.PushBackList(other)
}

func (l *ThreadSafeList) PushFront(v interface{}) *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.list.PushFront(v)
}

func (l *ThreadSafeList) PushFrontList(other *list.List) {
	l.Lock()
	defer l.Unlock()
	l.list.PushFrontList(other)
}

// locks the List
func (l *ThreadSafeList) Lock() {
	l.lockMutex.Lock()
}

// unlocks the List
func (l *ThreadSafeList) Unlock() {
	l.lockMutex.Unlock()
}
