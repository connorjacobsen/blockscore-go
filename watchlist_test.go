package blockscore

import "testing"

func init() {
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

var watchlistParams = WatchlistParams{
	CandidateId: "",
	MatchType:   "person",
}

func TestWatchlistSearch(t *testing.T) {
	candidates, err := Candidates.List()
	watchlistParams.CandidateId = candidates[0].Id

	_, err = Watchlists.Search(&watchlistParams)

	if err != nil {
		t.Errorf("Expected successful Watchlist search, got Error %s", err.Error())
		return
	}
}
