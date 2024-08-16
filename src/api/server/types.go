package server

type Configuration struct {
	AppName string
	Server  *Server
	Mongo   *Mongo
	Sql     *Sql
	Kafka   *Kafka
}

type Server struct {
	Port   string
	Origin string
	Header string
}

type Mongo struct {
	Host        string
	Port        string
	User        string
	Password    string
	Database    string
	Collections *Collections
}

type Collections struct {
	User string
}

type Sql struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

type Kafka struct {
	Broker string
	Group  string
	Topics *Topics
}

type Topics struct {
	CreatedUser string
	UpdatedUser string
	DeletedUser string
}
