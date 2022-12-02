package config

//config server
const (
	//Port is the port that the server will run on
	PortServer = "8080"
)

//config database
const (
	//Host is the host of the database
	DbHost = "localhost"
	//User is the user of the database
	DbUser = "root"
	//Password is the password of the database
	DbPassword = "root"
	//DBName is the name of the database
	DBName = "root"
	//Port is the port of the database
	PortDatabase = "5432"
	//SSLMode is the sslmode of the database
	SSLMode = "disable"
)

//config router config
const (
	//ApiVersion is the version of the api
	ApiVersion = "v1"
	//ApiName is the name of the api
	ApiName = "api"
)
