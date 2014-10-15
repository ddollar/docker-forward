package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"strings"

	"github.com/inconshreveable/muxado"
	"github.com/nitrous-io/tug/pipe"
)

func die(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func main() {
	conn := NewStdioConn(os.Stdin, os.Stdout)

	session := muxado.Server(conn)

	handleSession(session)
}

func handleSession(session muxado.Session) {
	for {
		stream, err := session.Accept()

		if err != nil && err != io.EOF {
			panic(err)
		}

		go handleStream(session, stream)
	}
}

func handleStream(session muxado.Session, stream muxado.Stream) {
	reader := bufio.NewReader(stream)
	command := readStreamLine(reader)

	switch command {
	case "connect":
		go func(r *bufio.Reader, s muxado.Stream) {
			handleConnect(r, s)
			s.Close()
		}(reader, stream)
	}
}

func handleConnect(reader *bufio.Reader, stream muxado.Stream) {
	dest := readStreamLine(reader)
	conn, err := net.Dial("tcp", dest)
	if err != nil {
		panic(err)
	}
	stream.Write([]byte("ok"))
	pipe.PipeStreams(stream, conn)
	conn.Close()
}

func readStreamLine(reader *bufio.Reader) string {
	raw, err := reader.ReadBytes('\n')
	if err != nil {
		return ""
	} else {
		return strings.TrimRight(string(raw), "\n")
	}
}

func sessionMessage(session muxado.Session, message string) error {
	stream, err := session.Open()
	if err != nil {
		return err
	}
	stream.Write([]byte("message\n"))
	stream.Write([]byte(message))
	stream.Close()
	return nil
}
