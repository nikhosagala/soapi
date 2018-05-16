package db

// Create function to insert data to database
func Create(create interface{}) interface{} {
	DB.NewRecord(create)
	DB.Create(create)
	return create
}

// Update function to update row to database
func Update(update interface{}) interface{} {
	DB.Model(update).Updates(update)
	return update
}

// Delete function to delete data from database
func Delete(delete interface{}) interface{} {
	DB.Delete(delete)
	return delete
}
