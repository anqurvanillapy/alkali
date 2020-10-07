package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

const defaultDirMode = 0750
const defaultFileMode = 0666

// NewListener creates new UDS listener.
func NewListener(path string) (net.Listener, error) {
	var err error

	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("Cannot remove unix://%s", path)
	}

	if err := os.MkdirAll(filepath.Dir(path), defaultDirMode); err != nil {
		log.Printf("WARN: Cannot create directory: path=%v, err=%v", path, err)
	}

	lis, err := net.Listen("unix", path)
	if err != nil {
		return nil, fmt.Errorf("Listen error: path=%q, err=%v", path, err)
	}

	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("UDS file not exists: path=%q", path)
	}
	if err := os.Chmod(path, defaultFileMode); err != nil {
		return nil, fmt.Errorf("Cannot update mode: path=%q", path)
	}

	return lis, nil
}

const echoServerPath = "/tmp/uds/echo"

// StartEchoServer starts a new UDS echo server.
func StartEchoServer() {
	lis, err := NewListener(echoServerPath)
	if err != nil {
		log.Fatalf("Cannot create UDS listener: err=%v", err)
	}
	defer lis.Close()

	log.Printf("INFO: UDS listening to %v", lis.Addr())

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("ERRO: Connection error: err=%v", err)
			continue
		}

		// Pass by value, not by reference!
		go func(c net.Conn) {
			defer c.Close()
			io.Copy(c, c)
		}(conn)
	}
}
