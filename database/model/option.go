package model

type Option struct {
	Text    string `bson:"text,omitempty"`
	Message string `bson:"message,omitempty"`
}
