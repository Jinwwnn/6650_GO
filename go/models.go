package swagger

// AlbumInfo represents the album information
type AlbumInfo struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Year   string `json:"year"`
}

// ImageMetadata represents the response metadata
type ImageMetadata struct {
	AlbumID   string `json:"albumID"`
	ImageSize string `json:"imageSize"`
}

// ErrorMessage represents an error response
type ErrorMessage struct {
	Msg string `json:"message"`
}
