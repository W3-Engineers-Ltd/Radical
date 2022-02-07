package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
)

const confVer = 0

const (
	Version       = "1.0.10"
	GitRemotePath = "github.com/W3-Engineers-Ltd/Radical"
)

var Conf = struct {
	Version            int
	WatchExts          []string  `json:"watch_ext" yaml:"watch_ext"`
	WatchExtsStatic    []string  `json:"watch_ext_static" yaml:"watch_ext_static"`
	GoInstall          bool      `json:"go_install" yaml:"go_install"` // Indicates whether execute "go install" before "go build".
	DirStruct          dirStruct `json:"dir_structure" yaml:"dir_structure"`
	CmdArgs            []string  `json:"cmd_args" yaml:"cmd_args"`
	Envs               []string
	Bale               bale
	Database           database
	EnableReload       bool              `json:"enable_reload" yaml:"enable_reload"`
	EnableNotification bool              `json:"enable_notification" yaml:"enable_notification"`
	Scripts            map[string]string `json:"scripts" yaml:"scripts"`
}{
	WatchExts:       []string{".go"},
	WatchExtsStatic: []string{".html", ".tpl", ".js", ".css"},
	GoInstall:       true,
	DirStruct: dirStruct{
		Others: []string{},
	},
	CmdArgs: []string{},
	Envs:    []string{},
	Bale: bale{
		Dirs:   []string{},
		IngExt: []string{},
	},
	Database: database{
		Driver: "mysql",
	},
	EnableNotification: true,
	Scripts:            map[string]string{},
}

// dirStruct describes the application's directory structure
type dirStruct struct {
	WatchAll    bool `json:"watch_all" yaml:"watch_all"`
	Controllers string
	Models      string
	Others      []string // Other directories
}

// bale
type bale struct {
	Import string
	Dirs   []string
	IngExt []string `json:"ignore_ext" yaml:"ignore_ext"`
}

// database holds the database connection information
type database struct {
	Driver string
	Conn   string
	Dir    string
}

// LoadConfig loads the radical tool configuration.
// It looks for radicalfile or radical.json in the current path,
// and falls back to default configuration in case not found.
func LoadConfig() {
	currentPath, err := os.Getwd()
	if err != nil {
		radicalLogger.Log.Error(err.Error())
	}

	dir, err := os.Open(currentPath)
	if err != nil {
		radicalLogger.Log.Error(err.Error())
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		radicalLogger.Log.Error(err.Error())
	}

	for _, file := range files {
		switch file.Name() {
		case "radical.json":
			{
				err = parseJSON(filepath.Join(currentPath, file.Name()), &Conf)
				if err != nil {
					radicalLogger.Log.Errorf("Failed to parse JSON file: %s", err)
				}
				break
			}
		case "radicalfile":
			{
				err = parseYAML(filepath.Join(currentPath, file.Name()), &Conf)
				if err != nil {
					radicalLogger.Log.Errorf("Failed to parse YAML file: %s", err)
				}
				break
			}
		}
	}

	// Check format version
	if Conf.Version != confVer {
		radicalLogger.Log.Warn("Your configuration file is outdated. Please do consider updating it.")
		radicalLogger.Log.Hint("Check the latest version of radical's configuration file.")
	}

	// Set variables
	if len(Conf.DirStruct.Controllers) == 0 {
		Conf.DirStruct.Controllers = "controllers"
	}

	if len(Conf.DirStruct.Models) == 0 {
		Conf.DirStruct.Models = "models"
	}
}

func parseJSON(path string, v interface{}) error {
	var (
		data []byte
		err  error
	)
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	return err
}

func parseYAML(path string, v interface{}) error {
	var (
		data []byte
		err  error
	)
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, v)
	return err
}
