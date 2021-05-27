package main

import (
	"fmt"
	"sync"
)

type TcpConnection struct {
	Addr string
	Port int
}

type Pool struct {
	pool    *sync.Pool
	max     int
	idle    int
	count   int
	poolCnt int
	close   bool
	locker  *sync.Mutex
}

func NewPool(max, idle int) *Pool {
	pool := &Pool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &TcpConnection{"127.0.0.1", 8080}
			},
		},
		max:     max,
		idle:    idle,
		count:   idle,
		poolCnt: idle,
		locker:  &sync.Mutex{},
	}
	for idle > 0 {
		conn := pool.pool.Get()
		pool.pool.Put(conn)
		idle--
	}
	return pool
}

func (p *Pool) Get() (*TcpConnection, error) {
	p.locker.Lock()
	defer p.locker.Unlock()

	if p.close {
		return nil, fmt.Errorf("pool is closed!")
	}

	if p.poolCnt > 0 {
		p.poolCnt--
		conn := p.pool.Get()
		return conn.(*TcpConnection), nil
	}
	if p.count >= p.max {
		return nil, fmt.Errorf("max connection")
	}
	fmt.Println("创建新连接")
	conn := p.pool.Get()
	p.count++
	return conn.(*TcpConnection), nil
}

func (p *Pool) Put(conn *TcpConnection) {
	p.locker.Lock()
	defer p.locker.Unlock()

	if p.close {
		return
	}

	if p.poolCnt >= p.idle {
		p.count--
		fmt.Println("销毁连接")
		return
	}

	p.pool.Put(conn)
	p.poolCnt++
}

func (p *Pool) Close() {
	p.locker.Lock()
	defer p.locker.Unlock()

	for p.poolCnt > 0 {
		c := p.pool.Get()
		conn := c.(*TcpConnection)
		fmt.Printf("销毁连接: %s:%d\n", conn.Addr, conn.Port)
		p.poolCnt--
		p.count--
	}
	p.close = true
}

func main() {
	p := NewPool(5, 3)
	c1, _ := p.Get()
	c2, _ := p.Get()
	c3, _ := p.Get()
	c4, _ := p.Get()
	c5, _ := p.Get()
	c6, _ := p.Get()

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
	fmt.Println(c6)
	p.Put(c5)
	c7, _ := p.Get()
	fmt.Println(c7)
}
