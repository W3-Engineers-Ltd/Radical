package radiantpro

const radiantToml = `
	dsn = "root:123456@tcp(127.0.0.1:3306)/radiant"
	driver = "mysql"
	proType = "default"
	enableModule = []
	apiPrefix = "/"
	gitRemotePath = "https://github.com/W3-Engineers-Ltd/radiant-pro"
	format = true
	sourceGen = "text"
	gitPull = true
	[models.user]
		name = ["uid"]
		orm = ["auto"]
		comment = ["Uid"]
		
`
