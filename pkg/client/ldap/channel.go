package ldap

import (
	"context"
	"errors"
	"sync"

	"github.com/MicroOps-cn/fuck/log"
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-ldap/ldap/v3"
)

// channelPool implements the Pool interface based on buffered channels.
type channelPool struct {
	// storage for our net.Conn connections
	mu    sync.Mutex
	conns chan ldap.Client

	name        string
	aliveChecks bool

	// net.Conn generator
	factory PoolFactory
	closeAt []uint16
	logger  kitlog.Logger
}

// PoolFactory is a function to create new connections.
type PoolFactory func(string) (ldap.Client, error)

// NewChannelPool returns a new pool based on buffered channels with an initial
// capacity and maximum capacity. Factory is used when initial capacity is
// greater than zero to fill the pool. A zero initialCap doesn't fill the Pool
// until a new Get() is called. During a Get(), If there is no new connection
// available in the pool, a new connection will be created via the Factory()
// method.
//
// closeAt will automagically mark the connection as unusable if the return code
// of the call is one of those passed, most likely you want to set this to something
// like []uint8{ldap.LDAPResultTimeLimitExceeded, ldap.ErrorNetwork}
func NewChannelPool(ctx context.Context, initialCap, maxCap int, name string, factory PoolFactory, closeAt []uint16) (Pool, error) {
	logger := log.GetContextLogger(ctx)
	if initialCap < 0 || maxCap <= 0 || initialCap > maxCap {
		return nil, errors.New("invalid capacity settings")
	}
	c := &channelPool{
		conns:       make(chan ldap.Client, maxCap),
		name:        name,
		factory:     factory,
		closeAt:     closeAt,
		aliveChecks: true,
		logger:      logger,
	}

	// create initial connections, if something goes wrong,
	// just close the pool error out.
	for i := 0; i < initialCap; i++ {
		conn, err := factory(c.name)
		if err != nil {
			c.Close()
			return nil, errors.New("factory is not able to fill the pool: " + err.Error())
		}
		c.conns <- conn
	}

	return c, nil
}

func (c *channelPool) AliveChecks(on bool) {
	c.mu.Lock()
	c.aliveChecks = on
	c.mu.Unlock()
}

func (c *channelPool) getConns() chan ldap.Client {
	c.mu.Lock()
	conns := c.conns
	c.mu.Unlock()
	return conns
}

// Get implements the Pool interfaces Get() method. If there is no new
// connection available in the pool, a new connection will be created via the
// Factory() method.
func (c *channelPool) Get() (*PoolConn, error) {
	conns := c.getConns()
	if conns == nil {
		return nil, ErrClosed
	}

	// wrap our connections with our ldap.Client implementation (wrapConn
	// method) that puts the connection back to the pool if it's closed.
	select {
	case conn := <-conns:
		if conn == nil {
			return nil, ErrClosed
		}
		if !c.aliveChecks || isAlive(conn) {
			return c.wrapConn(conn, c.closeAt), nil
		}
		conn.Close()
		return c.NewConn()
	default:
		return c.NewConn()
	}
}

func isAlive(conn ldap.Client) bool {
	_, err := conn.Search(&ldap.SearchRequest{BaseDN: "", Scope: ldap.ScopeBaseObject, Filter: "(&)", Attributes: []string{"1.1"}})
	return err == nil
}

func (c *channelPool) NewConn() (*PoolConn, error) {
	conn, err := c.factory(c.name)
	if err != nil {
		return nil, err
	}
	return c.wrapConn(conn, c.closeAt), nil
}

// put puts the connection back to the pool. If the pool is full or closed,
// conn is simply closed. A nil conn will be rejected.
func (c *channelPool) put(conn ldap.Client) {
	if conn == nil {
		level.Warn(c.logger).Log("msg", "conction is nil. rejecting")
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conns == nil {
		// pool is closed, close passed connection
		conn.Close()
		return
	}

	// put the resource back into the pool. If the pool is full, this will
	// block and the default case will be executed.
	select {
	case c.conns <- conn:
		return
	default:
		// pool is full, close passed connection
		conn.Close()
		return
	}
}

func (c *channelPool) Close() {
	defer func() {
		if err := recover(); err != nil {
			level.Info(c.logger).Log("msg", "failed to close ldap pool", "err", err)
		}
	}()
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.factory = nil
	c.mu.Unlock()

	if conns == nil {
		return
	}
	if len(conns) != 0 {
		close(conns)
		for conn := range conns {
			conn.Close()
		}
	}
}

func (c *channelPool) Len() int { return len(c.getConns()) }

func (c *channelPool) wrapConn(conn ldap.Client, closeAt []uint16) *PoolConn {
	p := &PoolConn{c: c, closeAt: closeAt, logger: c.logger}
	p.Conn = conn
	return p
}
