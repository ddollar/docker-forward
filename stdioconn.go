package main

import (
	"net"
	"os"
	"time"
)

// matches the net.Conn interface using stdin/stdout
type StdioConn struct {
	Stdin  *os.File
	Stdout *os.File
}

func NewStdioConn(stdin, stdout *os.File) *StdioConn {
	return &StdioConn{Stdin: stdin, Stdout: stdout}
}

func (sc *StdioConn) Read(buf []byte) (int, error) {
	return sc.Stdin.Read(buf)
}

func (sc *StdioConn) Write(buf []byte) (int, error) {
	return sc.Stdout.Write(buf)
}

func (sc *StdioConn) Close() error {
	return nil
}

func (sc *StdioConn) LocalAddr() net.Addr {
	return nil
}

func (sc *StdioConn) RemoteAddr() net.Addr {
	return nil
}

func (sc *StdioConn) SetDeadline(time.Time) error {
	return nil
}

func (sc *StdioConn) SetReadDeadline(time.Time) error {
	return nil
}

func (sc *StdioConn) SetWriteDeadline(time.Time) error {
	return nil
}
