// Copyright (c) 2018 Trough Creek Holdings, LLC.  All Rights Reserved

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"syscall"
	"time"
)

var help = flag.Bool("h", false, "Print help message")
var timeout = flag.Int64("t", 90, "Timeout in seconds")

func main() {
	var err error

	log.SetFlags(0)
	flag.CommandLine.SetOutput(os.Stderr)

	flag.Parse()

	args := flag.Args()

	if len(args) != 1 || *help {
		fmt.Fprint(os.Stderr, "Usage: waitfor [-h] [-t timeout] <dial-string>")
		fmt.Fprint(os.Stderr, "\n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	network := "tcp"
	addr := "127.0.0.1:80"
	duration := time.Duration(*timeout) * time.Second

	s := strings.Split(args[0], "!")
	switch len(s) {
	case 1:
		addr = s[0]
	case 2:
		network = s[0]
		addr = s[1]
	default:
		fmt.Fprintf(os.Stderr, "malformed dial string: %q\n", args[0])
		os.Exit(1)
	}

	for retry := true; retry && duration > 0; {
		err, duration, retry = waitfor(network, addr, duration)
		if err == nil {
			os.Exit(0)
		} else if !retry {
			fmt.Fprintf(os.Stderr, "dial failure: %q\n", err)
			os.Exit(1)
		}
		time.Sleep(time.Second)

	}
	os.Exit(1)
}

func waitfor(network, addr string, duration time.Duration) (err error, remaining time.Duration, retry bool) {
	start := time.Now().Unix()

	if _, err = net.DialTimeout(network, addr, duration); err == nil {
		return
	}

	elapsed := time.Now().Unix() - start
	if elapsed < 1 {
		elapsed = 1
	}
	remaining = duration - time.Duration(elapsed) * time.Second

	switch t := err.(type) {
	case *net.OpError:
		if t.Temporary() || t.Timeout() {
			if remaining > 0 {
				retry = true
			}
			return
		}

		// dial: unknown host, connection refused
		// read: connection refused
		if t.Op == "dial" || t.Op == "read" {
			if remaining > 0 {
fmt.Fprintf(os.Stderr, "here 2 %q %d\n", t.Op, remaining)
				retry = true
			}
		}
	case syscall.Errno:
		if t == syscall.ECONNREFUSED {
			if remaining > 0 {
				retry = true
			}
		}
	}
	return
}
