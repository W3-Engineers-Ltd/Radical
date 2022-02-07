module github.com/W3-Engineers-Ltd/Radical

go 1.17

require (
	github.com/W3-Engineers-Ltd/Radiant v1.0.5
	github.com/davecgh/go-spew v1.1.1
	github.com/flosch/pongo2 v0.0.0-20200913210552-0d938eb266f3
	github.com/fsnotify/fsnotify v1.5.1
	github.com/go-delve/delve v1.8.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/websocket v1.4.2
	github.com/lib/pq v1.10.4
	github.com/pelletier/go-toml v1.9.4
	github.com/smartwalle/pongo2render v1.0.1
	github.com/spf13/viper v1.10.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/cilium/ebpf v0.8.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/peterh/liner v1.2.2 // indirect
	github.com/spf13/afero v1.8.0 // indirect
	go.starlark.net v0.0.0-20220203230714-bb14e151c28f // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	gopkg.in/ini.v1 v1.66.3 // indirect
)

//replace github.com/W3-Engineers-Ltd/Radiant => ../radiant
