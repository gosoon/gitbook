```
package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

const (
	LeaderMutexName = "elastic-resource-flow-mutex"
)

// Options contains cloud platform all config
type Options struct {
	MasterName    string
	SentinelAddrs []string
}

// redisManager
type redisManager struct {
	opts        *Options
	client      *redis.Client
	leaderMutex *redsync.Mutex
}

func NewRedisManager(opts *Options) RedisManager {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    opts.MasterName,
		SentinelAddrs: opts.SentinelAddrs,
	})

	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	leaderMutex := rs.NewMutex(LeaderMutexName, redsync.WithExpiry(3*time.Hour))

	// return rm.resourceDeliveryGlobalMutex.Lock()
	return &redisManager{
		client:      client,
		opts:        opts,
		leaderMutex: leaderMutex,
	}
}
```

```
package redis

import "fmt"

func (rm *redisManager) GetResourceDeliveryLock() error {
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	return rm.leaderMutex.Lock()
}

func (rm *redisManager) ReleaseResourceDeliveryLock() error {
	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := rm.leaderMutex.Unlock(); !ok || err != nil {
		return fmt.Errorf("unlock failed")
	}

	return nil
}
```

```
package redis

type RedisManager interface {
	// leader election
	GetLeaderLock() error
	ReleaseLeaderLock() error
}
```

