package model

type File struct {
	FileName  string `json:"file_name"`
	MimeType  string `json:"mime_type"`
	Directory string `json:"directory"`
	FileUrl   string `json:"file_url"`
}
