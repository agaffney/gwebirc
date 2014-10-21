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
	Nick      string
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
	c.readbuf = bufio.NewReader(c.conn)
	c.ChangeNick(c.user.nick)
	c.Send(fmt.Sprintf("USER %s %d * :%s\r\n", c.user.name, c.user.bitmask, c.user.real_name))
	c.Join("#gwebirc")
	// This should block until the connection is closed
	c.read_from_server()
}

func (c *Connection) Send(msg string) {
	c.conn.Write([]byte(msg))
	fmt.Printf("> %s", msg)
}

func (c *Connection) Join(channel string) {
	c.Send(fmt.Sprintf("JOIN %s\r\n", channel))
}

func (c *Connection) ChangeNick(nick string) {
	c.Send(fmt.Sprintf("NICK %s\r\n", nick))
	c.Nick = nick
}

func (c *Connection) read_from_server() {
	for {
		line, err := c.readbuf.ReadString('\n')
		if len(line) > 0 {
			cmd := parse_command(line)
			c.handle_server_command(cmd)
		}
		if err != nil {
			break
		}
	}
}
