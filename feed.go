package matchlight

// A FeedService can interact with Matchlight feeds.
type FeedService interface {
	List() ([]Feed, error)
	Prepare(req PrepareReq) (string, error)
	Link(name, feedResponseID string) (LinkStatus, error)
	ListHitCounts(name string) (FeedHits, error)
}

// A Feed does what...?
type Feed struct {
	Description      string `json:"description,omitempty"`
	Name             string `json:"name"`
	RecentAlertCount uint   `json:"recent_alert_count"`
	StartTimestamp   uint   `json:"start_timestamp"`
	StopTimestamp    *uint  `json:"stop_timestamp,omitempty"`
}

// A PrepareReq is a request struct for calling Prepare on the FeedService.
type PrepareReq struct {
	Name      string `json:"-"`
	StartDate uint   `json:"start_date"`
	EndDate   uint   `json:"end_date"`
}

// FeedHits is an alias for hit count per-day representation.
type FeedHits map[uint]uint

// LinkStatus is the result of calling Link on FeedService.
type LinkStatus struct {
	Progress string `json:"progress"`
	Status   string `json:"status"`
}
