package entity

type Audio struct {
	Title string `json:"title"`
	Description string `json:"description"`
	URL string `json:"url"`
}

type Folder struct {
	Audios []Audio
}