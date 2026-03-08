package main

import (
	"os"
	"os/exec"
	"syscall"
	"time"
)

const (
	FRPC_EXE  = `D:\...\frpc.exe`
	FRPC_INI  = `D:\...\frpc.ini`
	CHECK_INT = 5 * time.Second
)
const CREATE_NO_WINDOW = 0x08000000

func main() {
	if !pathExists(FRPC_EXE) {
		return
	}
	_ = exec.Command("taskkill", "/F", "/IM", "frpc.exe", "/T").Run()
	var cmd *exec.Cmd
	for {
		if !isAlive(cmd) {
			cmd = startSilentFrpc()
		}
		time.Sleep(CHECK_INT)
	}
}

func startSilentFrpc() *exec.Cmd {
	c := exec.Command(FRPC_EXE, "-c", FRPC_INI)
	c.Stdout = nil
	c.Stderr = nil
	c.Stdin = nil
	c.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: CREATE_NO_WINDOW,
		HideWindow:    true,
	}
	if err := c.Start(); err != nil {
		return nil
	}
	return c
}

func isAlive(cmd *exec.Cmd) bool {
	if cmd == nil || cmd.Process == nil {
		return false
	}
	p, err := os.FindProcess(cmd.Process.Pid)
	if err != nil {
		return false
	}
	err = p.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
