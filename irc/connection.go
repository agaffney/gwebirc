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
	mode      string
	channels  map[string]Channel
}

func (c *Connection) Start() {
	// Initialize a few values
	c.host_port = fmt.Sprintf("%s:%d", c.Host, c.Port)
	c.channels = make(map[string]Channel)
	dialer := &net.Dialer{Timeout: time.Duration(5) * time.Second}
	if c.Tls {
		conn, err := tls.DialWithDialer(dialer, "tcp", c.host_port, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			fmt.Printf("connection failed: %s\n", err)
		}
		c.conn = conn
	} else {
		conn, err := dialer.Dial("tcp", c.host_port)
		if err != nil {
			fmt.Printf("connection failed: %s\n", err)
		}
		c.conn = conn
	}
	fmt.Println("connection successful!")
	c.readbuf = bufio.NewReader(c.conn)
	c.conn.Write([]byte("NICK gwebirc\r\n"))
	c.conn.Write([]byte("USER gwebirc 0 * :gwebirc client\r\n"))
	c.conn.Write([]byte("JOIN #gwebirc\r\n"))
	// This should block until the connection is closed
	c.read_from_server()
}

func (c *Connection) Send(msg string) {
	c.conn.Write([]byte(msg))
	fmt.Printf("> %s", msg)
}
