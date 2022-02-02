package system

import (
	"os"
	"os/user"
	"path/filepath"
)

// radical System Params ...
var (
	Usr, _      = user.Current()
	RadiantHome = filepath.Join(Usr.HomeDir, "/.radiant")
	CurrentDir  = getCurrentDirectory()
	GoPath      = os.Getenv("GOPATH")
)

func getCurrentDirectory() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	return ""
}
