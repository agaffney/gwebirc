package irc

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type Connection struct {
	conn      net.Conn
	Host      string
	Port      int
	Tls       bool
	host_port string
	readbuf   *bufio.Reader
}

func (c *Connection) Connect() error {
	c.host_port = fmt.Sprintf("%s:%d", c.Host, c.Port)
	dialer := &net.Dialer{Timeout: time.Duration(5) * time.Second}
	if c.Tls {
		conn, err := tls.DialWithDialer(dialer, "tcp", c.host_port, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			fmt.Printf("connection failed: %s\n", err)
			return err
		}
		c.conn = conn
	} else {
		conn, err := dialer.Dial("tcp", c.host_port)
		if err != nil {
			fmt.Printf("connection failed: %s\n", err)
			return err
		}
		c.conn = conn
	}
	fmt.Println("connection successful!")
	c.readbuf = bufio.NewReader(c.conn)
	c.conn.Write([]byte("NICK Guest868734\r\n"))
	c.conn.Write([]byte("USER guest868 0 * :Guest868734\r\n"))
	for {
		str, err := c.readbuf.ReadString('\n')
		if len(str) > 0 {
			fmt.Println(str)
		}
		if err != nil {
			break
		}
	}
	return nil
}
