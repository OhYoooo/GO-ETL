package models

// DatastoreEntity is document formatted to save into db
type DatastoreEntity struct {
	UserID    string
	Password  string
	Timestamp int64
}
