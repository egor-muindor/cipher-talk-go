package storage

type Mutex interface {
	Lock()
	Unlock()
}
