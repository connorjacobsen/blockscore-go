package blockscore

import "testing"

func init() {
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

var watchlistParams = WatchlistParams{
	CandidateID: "",
	MatchType:   "person",
}

func TestWatchlistSearch(t *testing.T) {
	candidates, err := Candidates.List()
	watchlistParams.CandidateID = candidates[0].ID

	_, err = Watchlists.Search(&watchlistParams)

	if err != nil {
		t.Errorf("Expected successful Watchlist search, got Error %s", err.Error())
		return
	}
}
