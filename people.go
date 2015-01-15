package blockscore

import (
	"net/url"
	"strconv"
)

type Person struct {
	Object             string        `json:"object"`
	Id                 string        `json:"id"`
	CreatedAt          int64         `json:"created_at"`
	UpdatedAt          int64         `json:"updated_at"`
	Status             string        `json:"status"`
	Livemode           bool          `json:"livemode"`
	PhoneNumber        string        `json:"phone_number"`
	IPAddress          string        `json:"ip_address"`
	BirthDay           int64         `json:"birth_day"`
	BirthMonth         int64         `json:"birth_month"`
	BirthYear          int64         `json:"birth_year"`
	NameFirst          string        `json:"name_first"`
	NameMiddle         string        `json:"name_middle"`
	NameLast           string        `json:"name_last"`
	DocumentType       string        `json:"document_type"`
	DocumentValue      string        `json:"document_value"`
	AddressStreet1     string        `json:"address_street1"`
	AddressStreet2     string        `json:"address_street2"`
	AddressCity        string        `json:"address_city"`
	AddressSubdivision string        `json:"address_subdivision"`
	AddressPostalCode  string        `json:"address_postal_code"`
	AddressCountryCode string        `json:"address_country_code"`
	Note               string        `json:"note"`
	Details            PersonDetails `json:"details"`

	// Have to figure out this part:
	// QuestionSets       []QuestionSet `json:"question_sets"`
}

type PersonDetails struct {
	Address        string `json:"address"`
	AddressRisk    string `json:"address_risk"`
	Identification string `json:"identification"`
	DateOfBirth    string `json:"date_of_birth"`
	Ofac           string `json:"ofac"`
	Pep            string `json:"pep"`
}

// Options for creating People.
type PersonParams struct {
	NameFirst          string `json:"name_first"`
	NameMiddle         string `json:"name_middle"`
	NameLast           string `json:"name_last"`
	DocumentType       string `json:"document_type"`
	DocumentValue      string `json:"document_value"`
	BirthDay           int64  `json:"birth_day"`
	BirthMonth         int64  `json:"birth_month"`
	BirthYear          int64  `json:"birth_year"`
	AddressStreet1     string `json:"address_street1"`
	AddressStreet2     string `json:"address_street2"`
	AddressCity        string `json:"address_city"`
	AddressSubdivision string `json:"address_subdivision"`
	AddressPostalCode  string `json:"address_postal_code"`
	AddressCountryCode string `json:"address_country_code"`
	PhoneNumber        string `json:"phone_number"`
	IPAddress          string `json:"ip_address"`
	Note               string `json:"note"`
}

type PersonClient struct{}

func (self *PersonClient) Create(params *PersonParams) (*Person, error) {
	person := Person{}
	values := url.Values{
		"name_first":           {params.NameFirst},
		"name_middle":          {params.NameMiddle},
		"name_last":            {params.NameLast},
		"document_type":        {params.DocumentType},
		"document_value":       {params.DocumentValue},
		"birth_day":            {strconv.FormatInt(params.BirthDay, 10)},
		"birth_month":          {strconv.FormatInt(params.BirthMonth, 10)},
		"birth_year":           {strconv.FormatInt(params.BirthYear, 10)},
		"address_street1":      {params.AddressStreet1},
		"address_street2":      {params.AddressStreet2},
		"address_city":         {params.AddressCity},
		"address_subdivision":  {params.AddressSubdivision},
		"address_postal_code":  {params.AddressPostalCode},
		"address_country_code": {params.AddressCountryCode},
		"phone_number":         {params.PhoneNumber},
		"ip_address":           {params.IPAddress},
		"note":                 {params.Note},
	}
	err := query("POST", "/people", values, &person)
	return &person, err
}

func (self *PersonClient) Retrieve(id string) (*Person, error) {
	person := Person{}
	path := "/people/" + url.QueryEscape(id)
	err := query("GET", path, nil, &person)
	return &person, err
}

func (self *PersonClient) List() ([]*Person, error) {
	return self.list(25, 0)
}

func (self *PersonClient) ListN(count, offset int) ([]*Person, error) {
	if count != 0 {
		return self.list(count, offset)
	} else {
		// returning count of 0 causes a bug!
		return self.list(25, offset)
	}
}

func (self *PersonClient) list(count, offset int) ([]*Person, error) {
	type listPersonResp struct{ Data []*Person }
	resp := listPersonResp{}

	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := query("GET", "/people", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
