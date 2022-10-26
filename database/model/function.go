package model

type Function struct {
	FunctionName string   `bson:"function_name,omitempty"`
	ButtonType   string   `bson:"button_type,omitempty"`
	Subtopic     string   `bson:"subtopic,omitempty"`
	Options      []Option `bson:"options,omitempty"`
}
