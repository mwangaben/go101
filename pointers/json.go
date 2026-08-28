package pointers

type Config struct {
	Host    string `json:"host"`
	Port    *int   `json:"port"`
	Timeout *int   `json:"timeout"`
}
