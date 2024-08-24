package server

type Configuration struct {
	AppName string  `mapstructure:"AppName"`
	Server  *Server `mapstructure:"Server"`
	Mongo   *Mongo  `mapstructure:"Mongo"`
	Sql     *Sql    `mapstructure:"Sql"`
	Kafka   *Kafka  `mapstructure:"Kafka"`
}

type Server struct {
	Port   string   `mapstructure:"Port"`
	Origin []string `mapstructure:"Origin"`
	Header []string `mapstructure:"Header"`
}

type Mongo struct {
	Host        string       `mapstructure:"Host"`
	Port        string       `mapstructure:"Port"`
	User        string       `mapstructure:"User"`
	Password    string       `mapstructure:"Password"`
	Database    string       `mapstructure:"Database"`
	Collections *Collections `mapstructure:"Collections"`
}

type Collections struct {
	User string `mapstructure:"User"`
}

type Sql struct {
	Host         string `mapstructure:"Host"`
	Port         int    `mapstructure:"Port"`
	User         string `mapstructure:"User"`
	Password     string `mapstructure:"Password"`
	DatabaseName string `mapstructure:"DatabaseName"`
}

type Kafka struct {
	Broker string  `mapstructure:"Broker"`
	Group  string  `mapstructure:"Group"`
	Topics *Topics `mapstructure:"Topics"`
}

type Topics struct {
	CreatedUser string `mapstructure:"CreatedUser"`
	UpdatedUser string `mapstructure:"UpdatedUser"`
	DeletedUser string `mapstructure:"DeletedUser"`
}
