// Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

package main

import (
	"os"
	"os/user"
	"strconv"
	"syscall"

	"golang.org/x/sys/unix"
)

func lookUser(uname string) (u *user.User, err error) {
	if isInt(uname) {
		u, err = user.LookupId(uname)
	} else {
		u, err = user.Lookup(uname)
	}
	return
}

func lookGroup(gname string) (g *user.Group, err error) {
	if isInt(gname) {
		g, err = user.LookupGroupId(gname)
	} else {
		g, err = user.LookupGroup(gname)
	}
	return
}

func setUser(u *user.User, g *user.Group) (err error) {
	var glist []string
	var gid_list []int

	var uid, gid int

	if glist, err = u.GroupIds(); err != nil {
		return
	}

	if gid_list, err = arrAtoi(glist); err != nil {
		return
	}

	if err = syscall.Setgroups(gid_list); err != nil {
		return
	}

	if gid, err = strconv.Atoi(g.Gid); err != nil {
		return
	}

	if err = Setgid(gid); err != nil {
		return
	}

	if uid, err = strconv.Atoi(u.Uid); err != nil {
		return
	}

	if err = Setuid(uid); err != nil {
		return
	}

	if u.HomeDir != "" {
		if err = os.Setenv("HOME", u.HomeDir); err != nil {
			return
		}
	}

	return
}

func Setuid(uid int) (err error) {
	_, _, e1 := unix.RawSyscall(unix.SYS_SETUID, uintptr(uid), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setgid(gid int) (err error) {
	_, _, e1 := unix.RawSyscall(unix.SYS_SETGID, uintptr(gid), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
