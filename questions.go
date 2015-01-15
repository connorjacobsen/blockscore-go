package blockscore

type QuestionSet struct {
	Object         string     `json:"object"`
	Id             string     `json:"id"`
	CreatedAt      int        `json:"created_at"`
	UpdatedAt      int        `json:"updated_at"`
	Livemode       bool       `json:"livemode"`
	VerificationId string     `json:"verification_id"`
	Score          float64    `json:"score"`
	Expired        bool       `json:"expired"`
	TimeLimit      int        `json:"time_limit"`
	Questions      []Question `json:"questions"`
}

type Question struct {
	Id       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	AnswerId int    `json:"answer_id"`
	Answer   string `json:"answer"`
}

type QuestionSetParams struct {
	PersonId  string `json:"person_id"`
	TimeLimit int    `json:"time_limit"` // optional
}
