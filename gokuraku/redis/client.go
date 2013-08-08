package redis

import (
	"github.com/ToQoz/Gokuraku/gokuraku"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	protocol                     = "tcp"
	connectTimeout time.Duration = 0
	idleTimeout                  = 240 * time.Second
	maxIdle                      = 3
	pool           redis.Pool
)

func init() {
	pool = redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(protocol, gokuraku.Config.RedisAddr)
			if err != nil {
				return nil, err
			}

			if gokuraku.Config.RedisPassword != "" {
				if _, err := conn.Do("AUTH", gokuraku.Config.RedisPassword); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

func Get() redis.Conn {
	return pool.Get()
}

// Deleate to redis. (github.com/garyburd/redigo/redis)
var (
	ScanStruct = redis.ScanStruct
	Values     = redis.Values
	Strings    = redis.Strings
	Int        = redis.Int
	String     = redis.String
	Bool       = redis.Bool
	ErrNil     = redis.ErrNil
)
