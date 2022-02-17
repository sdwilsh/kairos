package main

import (
	"github.com/c3os-io/c3os/installer/utils"
	"github.com/pterm/pterm"
)

func Reboot() {
	pterm.Info.Println("Rebooting node")
	utils.SH("reboot")
}

func PowerOFF() {
	pterm.Info.Println("Shutdown node")
	utils.SH("shutdown -h now")
}
