// Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"syscall"
)

func init() {
	// BOTCH: so we don't switch away from a diddled thread
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

func main() {
	var err error
	var path string
	var uname string
	var gname string
	var glist []string
	var u *user.User
	var g *user.Group

	log.SetFlags(0)

	if len(os.Args) <= 2 {
		log.Printf("Usage: runas <user[:group]> <cmd> [args...]")
		os.Exit(1)
	}

	uspec := strings.Split(os.Args[1], ":")
	switch len(uspec) {
	case 1:
		uname = uspec[0]
	case 2:
		uname = uspec[0]
		gname = uspec[1]
	default:
		log.Fatalf("error: invalid user: %q", os.Args[1])
	}

	if u, err = lookUser(uname); err != nil {
		log.Fatalf("error: missing user: %v", err)
	}

	if gname == "" {
		if glist, err = u.GroupIds(); err != nil {
			log.Fatalf("error: missing group ID: %v", err)
		}

		gname = glist[0]
	}

	if g, err = lookGroup(gname); err != nil {
		log.Fatalf("error: missing group: %v", err)
	}

	if err = setUser(u, g); err != nil {
		log.Fatalf("error: setuid/setgid failed: %v", err)
	}

	if path, err = exec.LookPath(os.Args[2]); err != nil {
		log.Fatalf("error: could not find binary: %v", err)
	}

	if err = syscall.Exec(path, os.Args[2:], os.Environ()); err != nil {
		log.Fatalf("error: exec failed: %v", err)
	}
}
