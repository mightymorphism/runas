// Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

package main

import (
	"flag"
	"fmt"
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

var help = flag.Bool("h", false, "Print help message")
var re_exec = flag.Bool("r", true, "Re-exec as new user with bash -l")
var user_spec = flag.String("u", "", "Specify user[:group] for exec")

func main() {
	var err error
	var path string
	var uname string
	var gname string
	var glist []string
	var u *user.User
	var g *user.Group

	log.SetFlags(0)
	flag.CommandLine.SetOutput(os.Stderr)

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 || *help {
		fmt.Fprint(os.Stderr, "Usage: runas [options] program [args..]")
		fmt.Fprint(os.Stderr, "\n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *user_spec != "" {
		uspec := strings.Split(*user_spec, ":")
		switch len(uspec) {
		case 1:
			uname = uspec[0]
		case 2:
			uname = uspec[0]
			gname = uspec[1]
		default:
			log.Fatalf("error: invalid user spec: %q", *user_spec)
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
	}

	if *re_exec {
		if path, err = os.Readlink("/proc/self/exe"); err != nil {
			log.Fatalf("error: unable to read /proc/self/exe")
		}

		cmd := fmt.Sprintf("%s %s", path, ShellJoin(args...))
		bash_args := []string{"/bin/bash", "-l", "-c", cmd}

		if err = syscall.Exec("/bin/bash", bash_args, os.Environ()); err != nil {
			log.Fatalf("error: unable to re-exec self")
		}
	} else {
		if path, err = exec.LookPath(args[0]); err != nil {
			log.Fatalf("error: could not find binary: %v", err)
		}

		if err = syscall.Exec(path, args, os.Environ()); err != nil {
			log.Fatalf("error: exec failed: %v", err)
		}
	}
}
