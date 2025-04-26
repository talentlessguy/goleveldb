//go:build wasip1
// +build wasip1

package storage

import (
	"os"
)

type wasiFileLock struct {
	f *os.File
}

func (fl *wasiFileLock) release() error {
	return fl.f.Close()
}

func newFileLock(path string, readOnly bool) (fl fileLock, err error) {
	var flag int
	if readOnly {
		flag = os.O_RDONLY
	} else {
		flag = os.O_RDWR
	}
	f, err := os.OpenFile(path, flag, 0)
	if os.IsNotExist(err) {
		f, err = os.OpenFile(path, flag|os.O_CREATE, 0644)
	}
	if err != nil {
		return
	}
	// WASI does not support file locks (yet), so we just proceed without.
	fl = &wasiFileLock{f: f}
	return
}

func rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func isErrInvalid(err error) bool {
	// WASI may behave slightly differently, but mimic Unix behavior here.
	if err == os.ErrInvalid {
		return true
	}
	if patherr, ok := err.(*os.PathError); ok && patherr.Err == os.ErrInvalid {
		return true
	}
	return false
}

func syncDir(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	// WASI supports fdatasync but fsync is not guaranteed.
	// Just attempt to Sync(), ignore if not supported.
	if err := f.Sync(); err != nil && !isErrInvalid(err) {
		return err
	}
	return nil
}
