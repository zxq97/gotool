package register

import (
	"context"
	"log"
	"time"

	"github.com/zxq97/gotool/constant"
	"go.etcd.io/etcd/client/v3"
)

type EtcdRegister struct {
	svcKey     string
	svcValue   string
	done       chan struct{}
	etcdClient *clientv3.Client
	leaseID    clientv3.LeaseID
	keepAlice  <-chan *clientv3.LeaseKeepAliveResponse
}

func NewEtcdRegister(etcdClient *clientv3.Client, key, value string) *EtcdRegister {
	return &EtcdRegister{
		etcdClient: etcdClient,
		svcKey:     key,
		svcValue:   value,
	}
}

func (er *EtcdRegister) Register() (chan<- struct{}, error) {
	err := er.add()
	if err != nil {
		return nil, err
	}

	er.done = make(chan struct{})

	go er.keepAlive()

	return er.done, nil
}

func (er *EtcdRegister) Stop() {
	close(er.done)
}

func (er *EtcdRegister) add() error {
	ctx, cancel := context.WithTimeout(context.Background(), constant.DefaultTimeout)
	defer cancel()

	res, err := er.etcdClient.Grant(ctx, constant.EtcdLeaseTTL)
	if err != nil {
		return err
	}
	er.leaseID = res.ID

	er.keepAlice, err = er.etcdClient.KeepAlive(context.Background(), res.ID)
	if err != nil {
		return err
	}

	_, err = er.etcdClient.Put(context.Background(), er.svcKey, er.svcValue, clientv3.WithLease(res.ID))
	return err
}

func (er *EtcdRegister) keepAlive() {
	var err error
	t := time.NewTicker(constant.DefaultTicker)
	for {
		select {
		case <-er.done:
			_, err = er.etcdClient.Delete(context.Background(), er.svcKey)
			if err != nil {
				log.Println(err)
			}
			_, err = er.etcdClient.Revoke(context.Background(), er.leaseID)
			if err != nil {
				log.Println(err)
			}
		case res := <-er.keepAlice:
			if res == nil {
				err = er.add()
				if err != nil {
					log.Println(err)
				}
			}
		case <-t.C:
			err = er.add()
			if err != nil {
				log.Println(err)
			}
		}
	}
}
