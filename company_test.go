package blockscore

import "testing"

func init() {
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

var companyParams = CompanyParams{
	EntityName:               "BlockScore",
	TaxID:                    "123410000",
	IncorporationState:       "DE",
	IncorporationCountryCode: "US",
	IncorporationType:        "corporation",
	IncorporationDay:         23,
	IncorporationMonth:       8,
	IncorporationYear:        1993,
	Dbas:                     "BitRemit",
	RegistrationNumber:       "123123123",
	Email:                    "test@example.com",
	URL:                      "https://blockscore.com",
	PhoneNumber:              "6505555555",
	IPAddress:                "67.160.8.182",
	Note:                     "Much wow",
	AddressStreet1:           "12345 Main Street",
	AddressStreet2:           "#12",
	AddressCity:              "Palo Alto",
	AddressSubdivision:       "CA",
	AddressPostalCode:        "94025",
	AddressCountryCode:       "US",
}

var companyID string

func TestCreateCompany(t *testing.T) {
	resp, err := Companies.Create(&companyParams)

	if err != nil {
		t.Errorf("Expected successful Company creation, got Error %s", err.Error())
		return
	}

	companyID = resp.ID

	if resp.EntityName != companyParams.EntityName {
		t.Errorf("Expected EntityName: %s, got: %s", companyParams.EntityName, resp.EntityName)
	}

	if resp.TaxID != companyParams.TaxID {
		t.Errorf("Expected TaxId: %s, got: %s", companyParams.TaxID, resp.TaxID)
	}

	if resp.IncorporationState != companyParams.IncorporationState {
		t.Errorf("Expected IncorporationState: %s, got: %s", companyParams.IncorporationState, resp.IncorporationState)
	}

	if resp.IncorporationCountryCode != companyParams.IncorporationCountryCode {
		t.Errorf("Expected IncorporationCountryCode: %s, got: %s", companyParams.IncorporationCountryCode, resp.IncorporationCountryCode)
	}

	if resp.IncorporationType != companyParams.IncorporationType {
		t.Errorf("Expected IncorporationType: %s, got: %s", companyParams.IncorporationType, resp.IncorporationType)
	}

	if resp.Dbas != companyParams.Dbas {
		t.Errorf("Expected Dbas: %s, got: %s", companyParams.Dbas, resp.Dbas)
	}

	if resp.RegistrationNumber != companyParams.RegistrationNumber {
		t.Errorf("Expected RegistrationNumber: %s, got: %s", companyParams.RegistrationNumber, resp.RegistrationNumber)
	}

	if resp.Email != companyParams.Email {
		t.Errorf("Expected Email: %s, got: %s", companyParams.Email, resp.Email)
	}

	if resp.URL != companyParams.URL {
		t.Errorf("Expected Url: %s, got: %s", companyParams.URL, resp.URL)
	}

	if resp.PhoneNumber != companyParams.PhoneNumber {
		t.Errorf("Expected PhoneNumber: %s, got: %s", companyParams.PhoneNumber, resp.PhoneNumber)
	}

	if resp.Note != companyParams.Note {
		t.Errorf("Expected Note: %s, got: %s", companyParams.Note, resp.Note)
	}

	if resp.AddressStreet1 != companyParams.AddressStreet1 {
		t.Errorf("Expected AddressStreet1: %s, got: %s", companyParams.AddressStreet1, resp.AddressStreet1)
	}

	if resp.AddressStreet2 != companyParams.AddressStreet2 {
		t.Errorf("Expected AddressStreet2: %s, got: %s", companyParams.AddressStreet2, resp.AddressStreet2)
	}

	if resp.AddressCity != companyParams.AddressCity {
		t.Errorf("Expected AddressCity: %s, got: %s", companyParams.AddressCity, resp.AddressCity)
	}

	if resp.AddressSubdivision != companyParams.AddressSubdivision {
		t.Errorf("Expected AddressSubdivision: %s, got: %s", companyParams.AddressSubdivision, resp.AddressSubdivision)
	}

	if resp.AddressPostalCode != companyParams.AddressPostalCode {
		t.Errorf("Expected AddressPostalCode: %s, got: %s", companyParams.AddressPostalCode, resp.AddressPostalCode)
	}

	if resp.AddressCountryCode != companyParams.AddressCountryCode {
		t.Errorf("Expected AddressCountryCode: %s, got: %s", companyParams.AddressCountryCode, resp.AddressCountryCode)
	}
}

func TestRetrieveCompany(t *testing.T) {
	resp, err := Companies.Retrieve(companyID)

	if err != nil {
		t.Errorf("Expected successful Company retrieval, got Error %s", err.Error())
	}

	if resp.ID != companyID {
		t.Errorf("Expected Company with Id: %s, got: %s", companyID, resp.ID)
	}
}

func TestListCompany(t *testing.T) {
	_, err := Companies.List()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestListNCompany(t *testing.T) {
	_, err := Companies.ListN(2, 5)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
