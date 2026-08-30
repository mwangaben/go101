package structs

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"` // Omit if empty
	Password string `json:"-"`               // ignore in JSON
	Age      int    `json:"age,string"`      // Serialize as a string
}

type Config struct {
	Host    string `json:"host" xml:"host" yaml:"host"`
	Port    int    `json:"port" xml:"port" yaml:"port"`
	Timeout int    `json:"timeout" xml:"timeout" yaml:"timeout"`
}
