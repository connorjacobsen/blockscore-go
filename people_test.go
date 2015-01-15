package blockscore

import "testing"

func init() {
	// In order to execute the Unit Test, you must set your BlockScore
	// API key as environment variable: BLOCKSCORE_API_KEY=xxxx
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

var peopleParams = PersonParams{
	NameFirst:          "John",
	NameMiddle:         "P",
	NameLast:           "Denver",
	DocumentType:       "ssn",
	DocumentValue:      "0000",
	BirthDay:           7,
	BirthMonth:         6,
	BirthYear:          1980,
	AddressStreet1:     "1234 Main Street",
	AddressStreet2:     "APT 12",
	AddressCity:        "Palo Alto",
	AddressSubdivision: "California",
	AddressPostalCode:  "94025",
	AddressCountryCode: "US",
	PhoneNumber:        "123-456-78910",
	IPAddress:          "127.0.0.1",
	Note:               "Hello, world",
}

var personID string

func TestCreatePeople(t *testing.T) {
	// Create the People.
	resp, err := People.Create(&peopleParams)

	if err != nil {
		t.Errorf("Expected successful People creation, got Error %s", err.Error())
	}

	personID = resp.Id

	if resp.NameFirst != peopleParams.NameFirst {
		t.Errorf("Expected NameFirst: %s, got: %s", peopleParams.NameFirst, resp.NameFirst)
	}

	if resp.NameMiddle != peopleParams.NameMiddle {
		t.Errorf("Expected NameMiddle: %s, got: %s", peopleParams.NameMiddle, resp.NameMiddle)
	}

	if resp.NameLast != peopleParams.NameLast {
		t.Errorf("Expected NameLast: %s, got: %s", peopleParams.NameLast, resp.NameLast)
	}

	if resp.DocumentType != peopleParams.DocumentType {
		t.Errorf("Expected DocumentType: %s, got: %s", peopleParams.DocumentType, resp.DocumentType)
	}

	if resp.DocumentValue != peopleParams.DocumentValue {
		t.Errorf("Expected DocumentValue: %s, got: %s", peopleParams.DocumentValue, resp.DocumentValue)
	}

	if resp.BirthDay != peopleParams.BirthDay {
		t.Errorf("Expected BirthDay: %s, got: %s", peopleParams.BirthDay, resp.BirthDay)
	}

	if resp.BirthMonth != peopleParams.BirthMonth {
		t.Errorf("Expected BirthMonth: %s, got: %s", peopleParams.BirthMonth, resp.BirthMonth)
	}

	if resp.BirthYear != peopleParams.BirthYear {
		t.Errorf("Expected BirthYear: %s, got: %s", peopleParams.BirthYear, resp.BirthYear)
	}

	if resp.AddressStreet1 != peopleParams.AddressStreet1 {
		t.Errorf("Expected AddressStreet1: %s, got: %s", peopleParams.AddressStreet1, resp.AddressStreet1)
	}

	if resp.AddressStreet2 != peopleParams.AddressStreet2 {
		t.Errorf("Expected AddressStreet2: %s, got: %s", peopleParams.AddressStreet2, resp.AddressStreet2)
	}

	if resp.AddressCity != peopleParams.AddressCity {
		t.Errorf("Expected AddressCity: %s, got: %s", peopleParams.AddressCity, resp.AddressCity)
	}

	if resp.AddressSubdivision != peopleParams.AddressSubdivision {
		t.Errorf("Expected AddressSubdivision: %s, got: %s", peopleParams.AddressSubdivision, resp.AddressSubdivision)
	}

	if resp.AddressPostalCode != peopleParams.AddressPostalCode {
		t.Errorf("Expected AddressPostalCode: %s, got: %s", peopleParams.AddressPostalCode, resp.AddressPostalCode)
	}

	if resp.AddressCountryCode != peopleParams.AddressCountryCode {
		t.Errorf("Expected AddressCountryCode: %s, got: %s", peopleParams.AddressCountryCode, resp.AddressCountryCode)
	}

	if resp.PhoneNumber != peopleParams.PhoneNumber {
		t.Errorf("Expected PhoneNumber: %s, got: %s", peopleParams.PhoneNumber, resp.PhoneNumber)
	}

	if resp.IPAddress != peopleParams.IPAddress {
		t.Errorf("Expected IPAddress: %s, got: %s", peopleParams.IPAddress, resp.IPAddress)
	}

	if resp.Note != peopleParams.Note {
		t.Errorf("Expected Note: %s, got: %s", peopleParams.Note, resp.Note)
	}
}

func TestRetrievePeople(t *testing.T) {
	resp, err := People.Retrieve(personID)

	if err != nil {
		t.Errorf("Expected successful People creation, got Error %s", err.Error())
	}

	if resp.NameFirst != peopleParams.NameFirst {
		t.Errorf("Expected NameFirst: %s, got: %s", peopleParams.NameFirst, resp.NameFirst)
	}

	if resp.NameMiddle != peopleParams.NameMiddle {
		t.Errorf("Expected NameMiddle: %s, got: %s", peopleParams.NameMiddle, resp.NameMiddle)
	}

	if resp.NameLast != peopleParams.NameLast {
		t.Errorf("Expected NameLast: %s, got: %s", peopleParams.NameLast, resp.NameLast)
	}

	if resp.DocumentType != peopleParams.DocumentType {
		t.Errorf("Expected DocumentType: %s, got: %s", peopleParams.DocumentType, resp.DocumentType)
	}

	if resp.DocumentValue != peopleParams.DocumentValue {
		t.Errorf("Expected DocumentValue: %s, got: %s", peopleParams.DocumentValue, resp.DocumentValue)
	}

	if resp.BirthDay != peopleParams.BirthDay {
		t.Errorf("Expected BirthDay: %s, got: %s", peopleParams.BirthDay, resp.BirthDay)
	}

	if resp.BirthMonth != peopleParams.BirthMonth {
		t.Errorf("Expected BirthMonth: %s, got: %s", peopleParams.BirthMonth, resp.BirthMonth)
	}

	if resp.BirthYear != peopleParams.BirthYear {
		t.Errorf("Expected BirthYear: %s, got: %s", peopleParams.BirthYear, resp.BirthYear)
	}

	if resp.AddressStreet1 != peopleParams.AddressStreet1 {
		t.Errorf("Expected AddressStreet1: %s, got: %s", peopleParams.AddressStreet1, resp.AddressStreet1)
	}

	if resp.AddressStreet2 != peopleParams.AddressStreet2 {
		t.Errorf("Expected AddressStreet2: %s, got: %s", peopleParams.AddressStreet2, resp.AddressStreet2)
	}

	if resp.AddressCity != peopleParams.AddressCity {
		t.Errorf("Expected AddressCity: %s, got: %s", peopleParams.AddressCity, resp.AddressCity)
	}

	if resp.AddressSubdivision != peopleParams.AddressSubdivision {
		t.Errorf("Expected AddressSubdivision: %s, got: %s", peopleParams.AddressSubdivision, resp.AddressSubdivision)
	}

	if resp.AddressPostalCode != peopleParams.AddressPostalCode {
		t.Errorf("Expected AddressPostalCode: %s, got: %s", peopleParams.AddressPostalCode, resp.AddressPostalCode)
	}

	if resp.AddressCountryCode != peopleParams.AddressCountryCode {
		t.Errorf("Expected AddressCountryCode: %s, got: %s", peopleParams.AddressCountryCode, resp.AddressCountryCode)
	}

	if resp.PhoneNumber != peopleParams.PhoneNumber {
		t.Errorf("Expected PhoneNumber: %s, got: %s", peopleParams.PhoneNumber, resp.PhoneNumber)
	}

	if resp.IPAddress != peopleParams.IPAddress {
		t.Errorf("Expected IPAddress: %s, got: %s", peopleParams.IPAddress, resp.IPAddress)
	}

	if resp.Note != peopleParams.Note {
		t.Errorf("Expected Note: %s, got: %s", peopleParams.Note, resp.Note)
	}
}

func TestListPeople(t *testing.T) {
	_, err := People.List()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestListNPeople(t *testing.T) {
	_, err := People.ListN(2, 5)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
