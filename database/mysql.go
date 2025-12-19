package database

var connection string

func init() {
	connection = "MySQL Database Connection Established"
}

func GetConnection() string {
	return connection
}