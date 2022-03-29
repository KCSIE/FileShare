package models

import "time"

type File struct {
	ID         string    `json:"id"`
	GroupId      string    `json:"group"`
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	MediaType  string    `json:"media_type"`
	UploadDate time.Time `json:"upload_date"`
	URL        string    `json:"url"`
}

type FileGroup struct {
	ID    string `json:"id"`
	Files []File `json:"files"`
	QR    string
}