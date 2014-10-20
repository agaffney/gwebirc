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
	user      User
	channels  map[string]*Channel
}

func (c *Connection) Start() {
	// Initialize a few values
	c.host_port = fmt.Sprintf("%s:%d", c.Host, c.Port)
	c.channels = make(map[string]*Channel)
	c.user = User{name: "gwebirc", nick: "gwebirc", bitmask: 0, real_name: "gwebirc client"}
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
	c.Send(fmt.Sprintf("NICK %s\r\n", c.user.nick))
	c.Send(fmt.Sprintf("USER %s %d * :%s\r\n", c.user.name, c.user.bitmask, c.user.real_name))
	c.Send("JOIN #gwebirc\r\n")
	// This should block until the connection is closed
	c.read_from_server()
}

func (c *Connection) Send(msg string) {
	c.conn.Write([]byte(msg))
	fmt.Printf("> %s", msg)
}
