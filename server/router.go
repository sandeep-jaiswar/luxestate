package server

func StartServer() {
    r := SetupRouter()
    r.Run()
}
