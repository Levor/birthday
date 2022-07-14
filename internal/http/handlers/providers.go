package handlers

func HandlerProviders() []interface{} {
	return []interface{}{
		NewHealthCheckHandler,
		NewWorkersHandler,
		NewUserHandler,
	}
}
