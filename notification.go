package parallel

import "sync"

type Notification struct {
	wg sync.WaitGroup
}

func (n *Notification) Done() { n.wg.Done() }
func (n *Notification) Wait() { n.wg.Wait() }

func NewNotification() *Notification {
	notification := &Notification{}
	notification.wg.Add(1)
	return notification
}
