// +build !darwin

package api

import (
	"os"
	"syscall"
)

func checkDirectIOFlag(path string) bool {
	// check if fs where disk.img file is located or block device
	// support direct i/o
	f, err := os.OpenFile(path, syscall.O_RDONLY|syscall.O_DIRECT, 0)
	defer f.Close()
	if err != nil && !os.IsNotExist(err) {
		return false
	}
	return true
}
