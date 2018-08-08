package matchlight

import (
	"image"
	"time"
)

// An IncidentService can interact with Matchlight incidents.
type IncidentService interface {
	GetScreenshot(id string) (image.Image, error)
	List() ([]Incident, error)
	Get(id string) (Incident, error)
}

// An Incident represents a Terbium incident.
// TODO: Investigate why so many of these are arrays.
type Incident struct {
	ID        string    `json:"id"`
	Impact    []string  `json:"impact"`
	Industry  []string  `json:"industry"`
	TimeStamp time.Time `json:"ts"`
	Actor     []string  `json:"actor"`
	Summary   string    `json:"summary"`
	Motive    []string  `json:"motive"`
	Type      []string  `json:"type"`
}
