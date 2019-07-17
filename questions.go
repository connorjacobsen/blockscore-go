package blockscore

import (
	"net/url"
	"strconv"
)

// QuestionSet helps you authenticate customers to see if they are who they say they are
type QuestionSet struct {
	Object         string     `json:"object"`
	ID             string     `json:"id"`
	CreatedAt      int        `json:"created_at"`
	UpdatedAt      int        `json:"updated_at"`
	Livemode       bool       `json:"livemode"`
	VerificationID string     `json:"verification_id"`
	Score          float64    `json:"score"`
	Expired        bool       `json:"expired"`
	TimeLimit      int        `json:"time_limit"`
	Questions      []Question `json:"questions"`
}

// Question has a question and a list of answers
type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

// Answer has an id and the actual answer
type Answer struct {
	ID     int    `json:"id"`
	Answer string `json:"answer"`
}

// type QuestionSetParams struct {
// 	PersonId  string `json:"person_id"`
// 	TimeLimit int    `json:"time_limit"` // optional
// }

// ScoreParams has a list of answers
type ScoreParams struct {
	Answers []ScoreAnswer `json:"answers"`
}

// ScoreAnswer stores the questionID , answerID
type ScoreAnswer struct {
	QuestionID int `json:"question_id"`
	AnswerID   int `json:"answer_id"`
}

// QuestionSetClient wraps questionSet related methods
type QuestionSetClient struct{}

// Create creates a new question set
func (questionSetClient *QuestionSetClient) Create(personID string) (*QuestionSet, error) {
	questionSet := QuestionSet{}
	values := url.Values{
		"person_id": {personID},
	}
	err := query("POST", "/question_sets", values, &questionSet)
	return &questionSet, err
}

// Retrieve return a historical record of all question sets that you have created.
func (questionSetClient *QuestionSetClient) Retrieve(id string) (*QuestionSet, error) {
	questionSet := QuestionSet{}
	path := "/question_sets/" + url.QueryEscape(id)
	err := query("GET", path, nil, &questionSet)
	return &questionSet, err
}

// List returns the last 25 question lists
func (questionSetClient *QuestionSetClient) List() ([]*QuestionSet, error) {
	return questionSetClient.list(25, 0)
}

// ListN returns the given count of question sets
func (questionSetClient *QuestionSetClient) ListN(count, offset int) ([]*QuestionSet, error) {
	if count != 0 {
		return questionSetClient.list(count, offset)
	}
	return questionSetClient.list(25, offset)
}

func (questionSetClient *QuestionSetClient) list(count, offset int) ([]*QuestionSet, error) {
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

// Score sends answers to blockscore to save it to the questionlist
func (questionSetClient *QuestionSetClient) Score(id string, params *ScoreParams) (*QuestionSet, error) {
	questionSet := QuestionSet{}
	path := "/question_sets/" + url.QueryEscape(id) + "/score"

	values := url.Values{}
	for _, answer := range params.Answers {
		// Create the correct array for answer scoring.
		values.Add("answers[][question_id]", strconv.Itoa(answer.QuestionID))
		values.Add("answers[][answer_id]", strconv.Itoa(answer.AnswerID))
	}

	err := query("POST", path, values, &questionSet)
	return &questionSet, err
}
