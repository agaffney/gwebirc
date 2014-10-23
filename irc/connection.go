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
	Channels  map[string]*Channel
	host_port string
	readbuf   *bufio.Reader
	user      User
	handlers  map[string][]HandlerFunc
}

func (c *Connection) Init() {
	// Initialize a few values
	c.host_port = fmt.Sprintf("%s:%d", c.Host, c.Port)
	c.Channels = make(map[string]*Channel)
	c.user = User{name: "gwebirc", nick: "gwebirc", bitmask: 0, real_name: "gwebirc client"}
	c.setup_handlers()
}

func (c *Connection) Start() {
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
	msg += "\r\n"
	c.conn.Write([]byte(msg))
	fmt.Printf("> %s", msg)
}

func (c *Connection) Join(channel string) {
	c.Send("JOIN " + channel)
}

func (c *Connection) ChangeNick(nick string) {
	c.Send("NICK " + nick)
	c.Nick = nick
}

func (c *Connection) Privmsg(target string, msg string) {
	c.Send(fmt.Sprintf("PRIVMSG %s :%s", target, msg))
}

func (c *Connection) Notice(target string, msg string) {
	c.Send(fmt.Sprintf("NOTICE %s :%s", target, msg))
}

func (c *Connection) Ctcp(code string, target string, msg string) {
	if msg != "" {
		msg = " " + msg
	}
	c.Privmsg(target, fmt.Sprintf("\x01%s%s\x01", code, msg))
}

func (c *Connection) CtcpResponse(code string, target string, msg string) {
	c.Notice(target, fmt.Sprintf("\x01%s %s\x01", code, msg))
}

func (c *Connection) read_from_server() {
	for {
		line, err := c.readbuf.ReadString('\n')
		if len(line) > 0 {
			e := &Event{Raw: line}
			e.parse()
			c.handle_event(e)
		}
		if err != nil {
			break
		}
	}
}
