package publishstorypayload

type CreateStory struct {
	UserId  int      `json: "userId"`
	Title   string   `json: "title"`
	Summary string   `json: "summary"`
	Mood    string   `json: "mood"`
	Media   []string `json: "media"`
}
