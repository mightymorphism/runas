// Copyright (c) 2018 Trough Creek Holdings, LLC.  All Rights Reserved

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

	if _, err = net.DialTimeout(network, addr, duration); err != nil {
		fmt.Fprintf(os.Stderr, "dial failure: %q\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
