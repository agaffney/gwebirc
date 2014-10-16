package irc

import (
	"fmt"
	"net"
	"time"
)

type Connection struct {
	conn net.Conn
	host string
	port int
}

func (c Connection) Connect(host string, port int) error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Duration(5)*time.Second)
	if err != nil {
		fmt.Printf("connection failed: %v", err)
		return err
	}
	fmt.Println("connection successful!")
	c.conn = conn
	return nil
}
