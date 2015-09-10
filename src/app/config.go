package app

/*
	Application Configuration structure
*/

type Config struct {
	Server struct {
		Port int
	}
	Db struct {
		CreateDB         bool
		ConnectionString string
		Dialect          string
	}
}
