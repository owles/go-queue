package go_queue

import "sync"

type Driver string

const (
	DriverSync  Driver = "sync"
	DriverRedis        = "redis"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Database int
}

type Connection struct {
	Driver Driver
	Redis  *RedisConfig
}

type Connections struct {
	defaultConnection string

	list map[string]*Connection
	mu   sync.Mutex
}

func NewConnections() *Connections {
	return &Connections{
		list: make(map[string]*Connection),
	}
}

func (c *Connections) Default() *Connection {
	c.mu.Lock()
	defer c.mu.Unlock()

	if conn, ok := c.list[c.defaultConnection]; ok {
		return conn
	}

	return nil
}

func (c *Connections) GetDefault() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.defaultConnection
}

func (c *Connections) SetDefault(connection string) *Connections {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.defaultConnection = connection

	return c
}

func (c *Connections) Add(connectionName string, connection *Connection) *Connections {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.list[connectionName] = connection

	if c.defaultConnection == "" {
		c.defaultConnection = connectionName
	}

	return c
}

func (c *Connections) Get(connection string) *Connection {
	c.mu.Lock()
	defer c.mu.Unlock()

	connName := connection
	if connName == "" {
		connName = c.defaultConnection
	}

	if conn, ok := c.list[connection]; ok {
		return conn
	}

	return nil
}
