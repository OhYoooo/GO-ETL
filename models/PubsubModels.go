package models

// PubsubModel is the struct reading from pubsub
type PubsubModel struct {
	UserID    string `json:"userId"`
	Password string `json:"password"`
	Timestamp string `json:"timestamp"`
}
