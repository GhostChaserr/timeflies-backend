package story

type Story struct {
	Id      int64    `json: "id"`
	Title   string   `json: "title"`
	Summary string   `json: "summary"`
	Mood    string   `json: "mood"`
	Media   []string `json: "media"`
}
