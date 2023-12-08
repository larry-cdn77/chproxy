//go:build linux
package main

import "github.com/coreos/go-systemd/v22/daemon"

func sdNotifyReady() {
	daemon.SdNotify(false, daemon.SdNotifyReady)
}
