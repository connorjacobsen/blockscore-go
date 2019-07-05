package blockscore

import (
	"net/url"
)

// Watchlist takes a candidate token and perform a global watchlist search
type Watchlist struct {
	Object        string   `json:"object"`
	Livemode      bool     `json:"livemode"`
	SearchedLists []string `json:"searched_lists"`
	Matches       []Match  `json:"matches"`
}

// Match is a matching watchlistd
type Match struct {
	WatchlistName      string   `json:"watchlist_name"`
	MatchingInfo       []string `json:"matching_info"`
	NameFull           string   `json:"name_full"`
	AlternateNames     string   `json:"alternate_names"`
	DateOfBirth        string   `json:"date_of_birth"`
	Ssn                string   `json:"ssn"`
	Passport           string   `json:"passport"`
	AddressRaw         string   `json:"address_raw"`
	AddressStreet1     string   `json:"address_street1"`
	AddressCity        string   `json:"address_city"`
	AddressState       string   `json:"address_state"`
	AddressPostalCode  string   `json:"address_postal_code"`
	AddressCountryCode string   `json:"address_country_code"`
}

// WatchlistParams has the paramters for a watchlist
type WatchlistParams struct {
	// The ID of the Candidate you have created.
	CandidateID string `json:"candidate_id"`

	// Can either be `person` or `company` and will restrict the search to
	// only search for people or entities on watchlists respectively.
	// Optional param.
	MatchType string `json:"match_type"`
}

// WatchlistClient wraps watchlist related methods
type WatchlistClient struct{}

// Search executes a search for the given watchlist
func (watchlistClient *WatchlistClient) Search(params *WatchlistParams) (*Watchlist, error) {
	watchlist := Watchlist{}
	values := url.Values{
		"candidate_id": {params.CandidateID},
		"match_type":   {params.MatchType},
	}
	err := query("POST", "/watchlists", values, &watchlist)
	return &watchlist, err
}
