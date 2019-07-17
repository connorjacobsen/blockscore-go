package blockscore

import "testing"

func init() {
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

var candidateParams = CandidateParams{
	NameFirst:          "John",
	NameLast:           "Brendekamp",
	Note:               "12341234",
	Ssn:                "0001",
	DateOfBirth:        "1940-08-11",
	AddressStreet1:     "1 Infinite Loop",
	AddressCity:        "Harare",
	AddressCountryCode: "ZW",
}

var updateParams = CandidateParams{
	Passport: "123456789",
}

var candidateID string

func TestCandidateCreate(t *testing.T) {
	resp, err := Candidates.Create(&candidateParams)

	if err != nil {
		t.Errorf("Expected successful Candidate creation, got Error %s", err.Error())
		return
	}

	candidateID = resp.ID

	if resp.NameFirst != candidateParams.NameFirst {
		t.Errorf("Expected NameFirst: %s, got: %s", candidateParams.NameFirst, resp.NameFirst)
	}

	if resp.NameLast != candidateParams.NameLast {
		t.Errorf("Expected NameFirst: %s, got: %s", candidateParams.NameLast, resp.NameLast)
	}

	if resp.Note != candidateParams.Note {
		t.Errorf("Expected Note: %s, got: %s", candidateParams.Note, resp.Note)
	}

	if resp.Ssn != candidateParams.Ssn {
		t.Errorf("Expected Ssn: %s, got: %s", candidateParams.Ssn, resp.Ssn)
	}

	if resp.DateOfBirth != candidateParams.DateOfBirth {
		t.Errorf("Expected DateOfBirth: %s, got: %s", candidateParams.DateOfBirth, resp.DateOfBirth)
	}

	if resp.AddressStreet1 != candidateParams.AddressStreet1 {
		t.Errorf("Expected AddressStreet1: %s, got: %s", candidateParams.AddressStreet1, resp.AddressStreet1)
	}

	if resp.AddressCity != candidateParams.AddressCity {
		t.Errorf("Expected AddressCity: %s, got: %s", candidateParams.AddressCity, resp.AddressCity)
	}

	if resp.AddressCountryCode != candidateParams.AddressCountryCode {
		t.Errorf("Expected AddressCountryCode: %s, got: %s", candidateParams.AddressCountryCode, resp.AddressCountryCode)
	}
}

func TestRetrieveCandidate(t *testing.T) {
	resp, err := Candidates.Retrieve(candidateID)

	if err != nil {
		t.Errorf("Expected successful Candidate retrieval, got Error: %s", err.Error())
	}

	if resp.ID != candidateID {
		t.Errorf("Expected Candidate with Id: %s, got: %s", candidateID, resp.ID)
	}
}

func TestUpdateCandidate(t *testing.T) {
	resp, err := Candidates.Update(candidateID, &updateParams)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if resp.Passport != "123456789" {
		t.Errorf("Expected Passport: %s, got: %s", updateParams.Passport, resp.Passport)
	}
}

func TestListCandidate(t *testing.T) {
	_, err := Candidates.List()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestListNCandidate(t *testing.T) {
	_, err := Candidates.ListN(2, 5)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestHistoryCandidate(t *testing.T) {
	_, err := Candidates.History(candidateID)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestHitsCandidate(t *testing.T) {
	_, err := Candidates.Hits(candidateID)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestDeleteCandidate(t *testing.T) {
	_, err := Candidates.Delete(candidateID)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
