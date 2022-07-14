package services

func ServiceProviders() []interface{} {
	return []interface{}{
		NewUserService,
		NewToken,
	}
}
