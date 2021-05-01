package createstory

type CreateStory struct {
	Title   string   `json: "title"`
	Summary string   `json: "summary"`
	Mood    string   `json: "mood"`
	Media   []string `json: "media"`
}
