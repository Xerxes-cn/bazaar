package balancer

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	factories[RandomBalancer] = NewRondom
}

type Rondom struct {
	sync.RWMutex
	hosts []string
	rnd   *rand.Rand
}

func NewRondom(hosts []string) Balancer {
	return &Rondom{
		hosts: hosts,
		rnd:   rand.New(rand.NewSource(time.Now().UnixMicro())),
	}
}

func (r *Rondom) Add(host string) {
	r.Lock()
	defer r.Unlock()
	for _, h := range r.hosts {
		if h == host {
			return
		}
	}
	r.hosts = append(r.hosts, host)
}

func (r *Rondom) Remove(host string) {
	r.Lock()
	defer r.Unlock()
	for i, h := range r.hosts {
		if h == host {
			r.hosts = append(r.hosts[:i], r.hosts[i+1:]...)
		}
	}
}
func (r *Rondom) Balance(_ string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.hosts) == 0 {
		return "", NoHostError
	}
	return r.hosts[r.rnd.Intn(len(r.hosts))], nil
}

func (r *Rondom) Inc(_ string)  {}
func (r *Rondom) Done(_ string) {}
