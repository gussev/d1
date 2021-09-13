package server

// Server service to run
type Server interface{
	Init()
	Run(port string)
}
