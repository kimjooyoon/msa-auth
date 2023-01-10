package api

func HandlerSetup() {
	memberHandler := InitializeMemberHandler()
	memberHandler.Mapping(Server)
}
