package blockscore

import "testing"

func init() {
	// In order to execute the Unit Test, you must set your BlockScore
	// API key as environment variable: BLOCKSCORE_API_KEY=xxxx
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

var questionSetID string

func TestQuestionSetCreate(t *testing.T) {
	people, err := People.List()
	personID := people[0].ID

	resp, err := QuestionSets.Create(personID)

	if err != nil {
		t.Errorf("Expected successful QuestionSet creation, got Error: %s", err.Error())
	}

	questionSetID = resp.ID
}

func TestQuestionSetRetrieve(t *testing.T) {
	_, err := QuestionSets.Retrieve(questionSetID)

	if err != nil {
		t.Errorf("Expected successful QuestionSet retrieval, got Error: %s", err.Error())
	}
}

func TestQuestionSetList(t *testing.T) {
	_, err := QuestionSets.List()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestQuestionSetListN(t *testing.T) {
	_, err := QuestionSets.ListN(2, 5)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestQuestionSetScore(t *testing.T) {
	params := ScoreParams{
		Answers: []ScoreAnswer{
			ScoreAnswer{QuestionID: 1, AnswerID: 1},
			ScoreAnswer{QuestionID: 2, AnswerID: 1},
			ScoreAnswer{QuestionID: 3, AnswerID: 1},
			ScoreAnswer{QuestionID: 4, AnswerID: 1},
			ScoreAnswer{QuestionID: 5, AnswerID: 1},
		},
	}

	_, err := QuestionSets.Score(questionSetID, &params)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
