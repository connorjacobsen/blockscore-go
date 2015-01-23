package blockscore

import "testing"

func init() {
	// In order to execute the Unit Test, you must set your BlockScore
	// API key as environment variable: BLOCKSCORE_API_KEY=xxxx
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

var questionSetId string

func TestQuestionSetCreate(t *testing.T) {
	people, err := People.List()
	personId := people[0].Id

	resp, err := QuestionSets.Create(personId)

	if err != nil {
		t.Errorf("Expected successful QuestionSet creation, got Error: %s", err.Error())
	}

	questionSetId = resp.Id
}

func TestQuestionSetRetrieve(t *testing.T) {
	_, err := QuestionSets.Retrieve(questionSetId)

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
