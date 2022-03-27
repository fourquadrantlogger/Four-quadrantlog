package model

type Blob struct {
	Log
	Blob []byte `json:"blob"`
}
