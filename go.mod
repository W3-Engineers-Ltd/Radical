module github.com/W3-Engineers-Ltd/Radical

go 1.13

require (
	github.com/W3-Engineers-Ltd/Radiant v0.0.0-20220202053408-9a1fc375156e
	github.com/davecgh/go-spew v1.1.1
	github.com/flosch/pongo2 v0.0.0-20200529170236-5abacdfa4915
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-delve/delve v1.5.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/websocket v1.4.2
	github.com/lib/pq v1.10.2
	github.com/pelletier/go-toml v1.9.2
	github.com/smartwalle/pongo2render v1.0.1
	github.com/spf13/viper v1.7.0
	gopkg.in/yaml.v2 v2.4.0
)

//replace github.com/W3-Engineers-Ltd/Radiant => ../radiant
