package matchlight

// A SearchService can interact with Matchlight detailed search.
type SearchService interface {
	Search(fingerprints []string) SearchResults
}

// SearchResults
type SearchResults struct {
	CWID      string `json:"cwid"`
	Score     uint   `json:"score"`
	TimeStamp uint   `json:"ts"`
	// TODO: Figure out how to represent this...
	//Urls
}
