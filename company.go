package blockscore

import (
	"net/url"
	"strconv"
)

type Company struct {
	Object                   string         `json:"object"`
	Id                       string         `json:"id"`
	CreatedAt                int64          `json:"created_at"`
	UpdatedAt                int64          `json:"updated_at"`
	Status                   string         `json:"status"`
	Livemode                 bool           `json:"livemode"`
	EntityName               string         `json:"entity_name"`
	TaxId                    string         `json:"tax_id"`
	IncorporationDate        string         `json:"incorporation_date"`
	IncorporationState       string         `json:"incorporation_state"`
	IncorporationCountryCode string         `json:"incorporation_country_code"`
	IncorporationType        string         `json:"incorporation_type"`
	Dbas                     string         `json:"dbas"`
	RegistrationNumber       string         `json:"registration_number"`
	Email                    string         `json:"email"`
	Url                      string         `json:"url"`
	PhoneNumber              string         `json:"phone_number"`
	IPAddress                string         `json:"ip_address"`
	Note                     string         `json:"note"`
	AddressStreet1           string         `json:"address_street1"`
	AddressStreet2           string         `json:"address_street2"`
	AddressCity              string         `json:"address_city"`
	AddressSubdivision       string         `json:"address_subdivision"`
	AddressPostalCode        string         `json:"address_postal_code"`
	AddressCountryCode       string         `json:"address_country_code"`
	Details                  CompanyDetails `json:"details"`
}

type CompanyDetails struct {
	EntityName string `json:"entity_name"`
	TaxId      string `json:"tax_id"`
	Ofac       string `json:"ofac"`
}

type CompanyParams struct {
	EntityName               string `json:"entity_name"`
	TaxId                    string `json:"tax_id"`
	IncorporationState       string `json:"incorporation_state"` // optional
	IncorporationCountryCode string `json:"incorporation_country_code"`
	IncorporationType        string `json:"incorporation_type"`
	IncorporationDay         int64  `json:"incorporation_day"`   // optional
	IncorporationMonth       int64  `json:"incorporation_month"` // optional
	IncorporationYear        int64  `json:"incorporation_year"`  // optional
	Dbas                     string `json:"dbas"`                // optional
	RegistrationNumber       string `json:"registration_number"` // optional
	Email                    string `json:"email"`               // optional
	Url                      string `json:"url"`                 // optional
	PhoneNumber              string `json:"phone_number"`        // optional
	IPAddress                string `json:"ip_address"`          // optional
	Note                     string `json:"note"`                // optional
	AddressStreet1           string `json:"address_street1"`
	AddressStreet2           string `json:"address_street2"`
	AddressCity              string `json:"address_city"`
	AddressSubdivision       string `json:"address_subdivision"`
	AddressPostalCode        string `json:"address_postal_code"`
	AddressCountryCode       string `json:"address_country_code"`
}

type CompanyClient struct{}

func (self *CompanyClient) Create(params *CompanyParams) (*Company, error) {
	company := Company{}
	values := url.Values{
		"entity_name":                {params.EntityName},
		"tax_id":                     {params.TaxId},
		"incorporation_state":        {params.IncorporationState},
		"incorporation_country_code": {params.IncorporationCountryCode},
		"incorporation_type":         {params.IncorporationType},
		"incorporation_day":          {params.IncorporationDay},
		"incorporation_month":        {params.IncorporationMonth},
		"incorporation_year":         {params.IncorporationYear},
		"dbas":                       {params.Dbas},
		"registration_number":        {params.RegistrationNumber},
		"email":                      {params.Email},
		"url":                        {params.Url},
		"phone_number":               {params.PhoneNumber},
		"ip_address":                 {params.IPAddress},
		"note":                       {params.Note},
		"address_street1":            {params.AddressStreet1},
		"address_street2":            {params.AddressStreet2},
		"address_city":               {params.AddressCity},
		"address_subdivision":        {params.AddressSubdivision},
		"address_postal_code":        {params.AddressPostalCode},
		"address_country_code":       {params.AddressCountryCode},
	}
	err := query("POST", "/companies", values, &company)
	return &company, err
}

func (self *CompanyClient) Retrieve(id string) (*Company, error) {
	company := Company{}
	path := "/companies/" + url.QueryEscape(id)
	err := query("GET", path, nil, &company)
	return &company, err
}

func (self *CompanyClient) List() ([]*Company, error) {
	return self.list(25, 0)
}

func (self *CompanyClient) ListN(count, offset int) ([]*Company, error) {
	if count != 0 {
		return self.list(count, offset)
	} else {
		return self.list(25, offset)
	}
}

func (self *CompanyClient) list(count, offset int) ([]*Company, error) {
	type listCompanyResp struct{ Data []*Company }
	resp := listCompanyResp{}

	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := query("GET", "/companies", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
