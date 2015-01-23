package blockscore

import (
	"net/url"
	"strconv"
)

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
	Id     int    `json:"id"`
	Answer string `json:"answer"`
}

// type QuestionSetParams struct {
// 	PersonId  string `json:"person_id"`
// 	TimeLimit int    `json:"time_limit"` // optional
// }

type ScoreParams struct {
	Answers []ScoreAnswer `json:"answers"`
}

type ScoreAnswer struct {
	QuestionId int `json:"question_id"`
	AnswerId   int `json:"answer_id"`
}

type QuestionSetClient struct{}

func (self *QuestionSetClient) Create(personId string) (*QuestionSet, error) {
	questionSet := QuestionSet{}
	values := url.Values{
		"person_id": {personId},
	}
	err := query("POST", "/question_sets", values, &questionSet)
	return &questionSet, err
}

func (self *QuestionSetClient) Retrieve(id string) (*QuestionSet, error) {
	questionSet := QuestionSet{}
	path := "/question_sets/" + url.QueryEscape(id)
	err := query("GET", path, nil, &questionSet)
	return &questionSet, err
}

func (self *QuestionSetClient) List() ([]*QuestionSet, error) {
	return self.list(25, 0)
}

func (self *QuestionSetClient) ListN(count, offset int) ([]*QuestionSet, error) {
	if count != 0 {
		return self.list(count, offset)
	} else {
		return self.list(25, offset)
	}
}

func (self *QuestionSetClient) list(count, offset int) ([]*QuestionSet, error) {
	type listQuestionSetResp struct{ Data []*QuestionSet }
	resp := listQuestionSetResp{}

	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := query("GET", "/question_sets", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (self *QuestionSetClient) Score(id string, params *ScoreParams) (*QuestionSet, error) {
	questionSet := QuestionSet{}
	path := "/question_sets/" + url.QueryEscape(id) + "/score"

	values := url.Values{}
	for _, answer := range params.Answers {
		// Create the correct array for answer scoring.
		values.Add("answers[][question_id]", strconv.Itoa(answer.QuestionId))
		values.Add("answers[][answer_id]", strconv.Itoa(answer.AnswerId))
	}

	err := query("POST", path, values, &questionSet)
	return &questionSet, err
}
