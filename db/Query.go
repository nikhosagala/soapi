package db

// Create ...
func Create(insert interface{}) interface{} {
	DB.NewRecord(insert)
	DB.Create(&insert)
	return &insert
}