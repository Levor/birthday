package db

func Providers() []interface{} {
	return []interface{}{
		//*gorm.DB
		NewConnection,
	}
}
