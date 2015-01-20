package blockscore

import (
	"net/url"
	"strconv"
)

type Candidate struct {
	Object             string `json:"object"`
	Id                 string `json:"id"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
	Livemode           bool   `json:"livemode"`
	Note               string `json:"note"`
	Ssn                string `json:"ssn"`
	Passport           string `json:"passport"`
	DateOfBirth        string `json:"date_of_birth"`
	NameFirst          string `json:"name_first"`
	NameMiddle         string `json:"name_middle"`
	NameLast           string `json:"name_last"`
	AddressStreet1     string `json:"address_street1"`
	AddressStreet2     string `json:"address_street2"`
	AddressCity        string `json:"address_city"`
	AddressSubdivison  string `json:"address_subdivision"`
	AddressPostalCode  string `json:"address_postal_code"`
	AddressCountryCode string `json:"address_country_code"`
}

type CandidateParams struct {
	NameFirst          string `json:"name_first,omitempty"`
	NameMiddle         string `json:"name_middle,omitempty"`
	NameLast           string `json:"name_last,omitempty"`
	Note               string `json:"note,omitempty"`
	Ssn                string `json:"ssn,omitempty"`
	Passport           string `json:"passport,omitempty"`
	DateOfBirth        string `json:"date_of_birth,omitempty"`
	AddressStreet1     string `json:"address_street1,omitempty"`
	AddressStreet2     string `json:"address_street2,omitempty"`
	AddressCity        string `json:"address_city,omitempty"`
	AddressSubdivison  string `json:"address_subdivision,omitempty"`
	AddressPostalCode  string `json:"address_postal_code,omitempty"`
	AddressCountryCode string `json:"address_country_code,omitempty"`
}

type hitsCandidateResp struct {
	Object string       `json:"object"`
	Data   []*Candidate `json:"data"`
}

type CandidateClient struct{}

func (self *CandidateClient) Create(params *CandidateParams) (*Candidate, error) {
	candidate := Candidate{}
	values := url.Values{
		"name_first":           {params.NameFirst},
		"name_middle":          {params.NameMiddle},
		"name_last":            {params.NameLast},
		"note":                 {params.Note},
		"ssn":                  {params.Ssn},
		"passport":             {params.Passport},
		"date_of_birth":        {params.DateOfBirth},
		"address_street1":      {params.AddressStreet1},
		"address_street2":      {params.AddressStreet2},
		"address_city":         {params.AddressCity},
		"address_subdivision":  {params.AddressSubdivison},
		"address_postal_code":  {params.AddressPostalCode},
		"address_country_code": {params.AddressCountryCode},
	}
	err := query("POST", "/candidates", values, &candidate)
	return &candidate, err
}

func (self *CandidateClient) Retrieve(id string) (*Candidate, error) {
	candidate := Candidate{}
	path := "/candidates/" + url.QueryEscape(id)
	err := query("GET", path, nil, &candidate)
	return &candidate, err
}

func (self *CandidateClient) Update(id string, params *CandidateParams) (*Candidate, error) {
	candidate := Candidate{}
	values := url.Values{
		"name_first":           {params.NameFirst},
		"name_middle":          {params.NameMiddle},
		"name_last":            {params.NameLast},
		"note":                 {params.Note},
		"ssn":                  {params.Ssn},
		"passport":             {params.Passport},
		"date_of_birth":        {params.DateOfBirth},
		"address_street1":      {params.AddressStreet1},
		"address_street2":      {params.AddressStreet2},
		"address_city":         {params.AddressCity},
		"address_subdivision":  {params.AddressSubdivison},
		"address_postal_code":  {params.AddressPostalCode},
		"address_country_code": {params.AddressCountryCode},
	}

	// Remove all blank elements from the map.
	// Note: Maps are passed by reference, which is why we dont need
	// to do anything with a return value here.
	clean(values)

	path := "/candidates/" + url.QueryEscape(id)
	err := query("PATCH", path, values, &candidate)
	return &candidate, err
}

func (self *CandidateClient) Delete(id string) (*Candidate, error) {
	candidate := Candidate{}
	path := "/candidates/" + url.QueryEscape(id)
	err := query("DELETE", path, nil, &candidate)
	return &candidate, err
}

func (self *CandidateClient) List() ([]*Candidate, error) {
	return self.list(25, 0)
}

func (self *CandidateClient) ListN(count, offset int) ([]*Candidate, error) {
	if count != 0 {
		return self.list(count, offset)
	} else {
		return self.list(25, offset)
	}
}

func (self *CandidateClient) list(count, offset int) ([]*Candidate, error) {
	type listCandidateResp struct{ Data []*Candidate }
	resp := listCandidateResp{}

	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := query("GET", "/candidates", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (self *CandidateClient) History(id string) ([]*Candidate, error) {
	type historyCandidateResp struct{ Data []*Candidate }
	resp := historyCandidateResp{}
	path := "/candidates/" + url.QueryEscape(id) + "/history"
	err := query("GET", path, nil, &resp)
	return resp.Data, err
}

func (self *CandidateClient) Hits(id string) ([]*Candidate, error) {
	resp := hitsCandidateResp{}
	path := "/candidates/" + url.QueryEscape(id) + "/hits"
	err := query("GET", path, nil, &resp)
	return resp.Data, err
}

// Ugly, temporary work-around.
func clean(v url.Values) url.Values {
	if v.Get("name_first") == "" {
		v.Del("name_first")
	}

	if v.Get("name_middle") == "" {
		v.Del("name_middle")
	}

	if v.Get("name_last") == "" {
		v.Del("name_last")
	}

	if v.Get("note") == "" {
		v.Del("note")
	}

	if v.Get("ssn") == "" {
		v.Del("ssn")
	}

	if v.Get("passport") == "" {
		v.Del("passport")
	}

	if v.Get("date_of_birth") == "" {
		v.Del("date_of_birth")
	}

	if v.Get("address_street1") == "" {
		v.Del("address_street1")
	}

	if v.Get("address_street2") == "" {
		v.Del("address_street2")
	}

	if v.Get("address_city") == "" {
		v.Del("address_city")
	}

	if v.Get("address_subdivision") == "" {
		v.Del("address_subdivision")
	}

	if v.Get("address_postal_code") == "" {
		v.Del("address_postal_code")
	}

	if v.Get("address_country_code") == "" {
		v.Del("address_country_code")
	}

	return v
}
