package etcdc

import (
	"context"

	"go.etcd.io/etcd/client/v3/concurrency"
)

const (
	maxTries = 5
)

type NeedDistLockingJob func() (interface{}, error)

// 需要double check
func RunJobWithDistLock(ctx context.Context, lockKey string, job NeedDistLockingJob) (interface{}, error) {
	session, err := concurrency.NewSession(EtcdC)
	if err != nil {
		return nil, err
	}
	mutex := concurrency.NewMutex(session, lockKey)
	err = mutex.Lock(ctx)
	if err != nil {
		return nil, err
	}
	res, err := job()
	if err != nil {
		mutex.Unlock(ctx)
		return nil, err
	}
	mutex.Unlock(ctx)
	return res, nil
}
