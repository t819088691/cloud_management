package setting

type Setting struct {
	DB DB
}
type DB struct {
	Engine       string
	Host         string
	Port         int
	User         string
	Password     string
	Database     string
	Prefix       string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
}

func Read() *Setting {
	var s Setting
	s.DB.Engine = "mysql"
	s.DB.Host = "127.0.0.1"
	s.DB.Port = 20000
	s.DB.User = "root"
	s.DB.Password = "000000"
	s.DB.Database = "aws_billing"
	s.DB.Prefix = ""
	s.DB.Charset = "utf8"
	s.DB.MaxIdleConns = 30
	s.DB.MaxOpenConns = 100
	return &s
}
