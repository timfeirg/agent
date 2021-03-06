package metric

import (
	"fmt"

	statsdlib "github.com/CMGS/statsd"
	log "github.com/Sirupsen/logrus"
)

func NewStatsdClient(addr string) *StatsDClient {
	return &StatsDClient{
		Addr: addr,
	}
}

type StatsDClient struct {
	Addr string
}

func (self *StatsDClient) Close() error {
	return nil
}

func (self *StatsDClient) Send(data map[string]float64, endpoint, tag string) error {
	remote, err := statsdlib.New(self.Addr)
	if err != nil {
		log.Errorf("[statsd] Connect statsd failed: %v", err)
		return err
	}
	defer remote.Close()
	defer remote.Flush()
	for k, v := range data {
		key := fmt.Sprintf("eru.%s.%s.%s", endpoint, tag, k)
		remote.Gauge(key, v)
	}
	return nil
}
