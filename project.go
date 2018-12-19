package matchlight

// A ProjectType represents the enumeration of potential types a Matchlight project
// can be.
type ProjectType string

// Possible values for ProjectType.
const (
	ProjectTypePII        = ProjectType("pii")
	ProjectTypeDocument   = ProjectType("document")
	ProjectTypeSourceCode = ProjectType("source_code")
	ProjectTypeAny        = ProjectType("")
)

// A ProjectService can interact with MatchLight projects.
type ProjectService interface {
	AddProject(req AddProjectReq) (AddProjectRes, error)
	ListProjects(projectType ProjectType) ([]Project, error)
	DeleteProject(uploadToken string) error
	EditProject(uploadToken, name string) (Project, error)
	GetProject(uploadToken string) (Project, error)
}

// AddProjectReq is the request for project.Add.
type AddProjectReq struct {
	Name           string      `json:"name"`
	Type           ProjectType `json:"type"`
	AlertThreshold *uint       `json:"alert_notification_threshold,omitempty"`
}

// AddProjectRes is the response from project.Add.
type AddProjectRes struct {
	ID          string      `json:"id"`
	Name        string      `json:"project_name"`
	Type        ProjectType `json:"project_type"`
	UploadToken string      `json:"project_upload_token"`
}

// A Project represents a Matchlight project.
type Project struct {
	UploadToken  string      `json:"project_upload_token"`
	Name         string      `json:"project_name"`
	Type         ProjectType `json:"project_type"`
	UnseenAlerts uint        `json:"number_of_unseen_alerts"`
	Records      uint        `json:"number_of_records"`
	LastModified uint        `json:"last_date_modified"`
}
