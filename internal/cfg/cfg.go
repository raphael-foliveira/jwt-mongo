package cfg

import "os"

var JwtSecret string

func Init() {
	JwtSecret = os.Getenv("JWT_SECRET")
}
