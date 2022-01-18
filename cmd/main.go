package cmd

import "flag"

var (
	server   *string
	port     *int
	endpoint *string

	pathToConfig *string
)

func init() {
	port = flag.Int("port", 80, "HTTP Port")
	server = flag.String("server", "localhost", "HTTP Host name or IP Address")
	endpoint = flag.String("endpoint", "api", "Endpoint `/api`, you can set param in config file")
	pathToConfig = flag.String("config", "", "Path to configuration file in format `json`")
}

func main() {
	// todo
	// web-service

}
