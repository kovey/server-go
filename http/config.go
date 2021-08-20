package http

type Config struct {
	Server ServConf `yaml:"server"`
	View   View     `yaml:"view"`
	Test   Test     `yaml:"test"`
}

type ServConf struct {
	Addr     Addr   `yaml:"addr"`
	LogDir   string `yaml:"log-dir"`
	LogLevel string `yaml:"log-level"`
}

type Addr struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type View struct {
	Path   string `yaml:"path"`
	Static string `yaml:"static"`
}

type Test struct {
	Open string `yaml:"open"`
}
