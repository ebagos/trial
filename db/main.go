package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"
)

const MSEC = 1000

var lockPath string
var LOCK_PATH_DEFAULT = "locker"

func readHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Read World")
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	fd, err := lock(lockPath)
	if err != nil {
		fmt.Fprintf(w, "Error from Write World: %v", err)
		return
	}
	time.Sleep(time.Millisecond * MSEC)
	err = unlock(fd, lockPath)
	if err != nil {
		fmt.Fprintf(w, "Error from Write World: %v", err)
		return
	}
	fmt.Fprintf(w, "Hello Write World")
}

func main() {
	myPort := os.Getenv("PORT")
	if myPort == "" {
		myPort = ":8083"
	}
	if !strings.HasPrefix(myPort, ":") {
		myPort = ":" + myPort
	}
	lockPath = os.Getenv("LOCK_FILE")
	if lockPath == "" {
		lockPath = LOCK_PATH_DEFAULT
	}
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(myPort, nil)
}

func lock(path string) (int, error) {
	fd, fd_err := syscall.Open(path, syscall.O_CREAT, 0)
	if fd_err != nil {
		log.Println("syscall.Open :", fd_err)
		return -1, fd_err
	}
	flock_err := syscall.Flock(fd, syscall.LOCK_EX)
	if flock_err != nil {
		log.Println("syscall.Flock :", flock_err)
		return -1, flock_err
	}
	return fd, nil
}

func unlock(fd int, path string) error {
	funlock_err := syscall.Flock(fd, syscall.LOCK_UN)
	if funlock_err != nil {
		log.Println("syscall.Funlock :", funlock_err)
		unlink_err := syscall.Unlink(path)
		if unlink_err != nil {
			log.Println("syscall.Unlink :", unlink_err)
		}
		return funlock_err
	}
	close_err := syscall.Close(fd)
	if close_err != nil {
		log.Println("syscall.Close :", close_err)
		return close_err
	}
	unlink_err := syscall.Unlink(path)
	if unlink_err != nil {
		log.Println("syscall.Unlink :", unlink_err)
	}
	return nil
}
