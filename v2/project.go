package v2

import (
	"encoding/json"
	"errors"
	"fmt"

	matchlight "github.com/eriktate/go-matchlight"
)

// A ProjectClient knows how to make requests for Matchlight projects.
type ProjectClient struct {
	Client
}

// NewProjectClient returns a new ProjectClient given a Client instance.
func NewProjectClient(c Client) ProjectClient {
	return ProjectClient{Client: c}
}

// AddProject creates a new project within Matchlight.
func (c ProjectClient) AddProject(req matchlight.AddProjectReq) (matchlight.AddProjectRes, error) {
	var res matchlight.AddProjectRes
	if err := c.post("project/add", req, &res); err != nil {
		return res, err
	}

	return res, nil
}

// ListProjects returns a list of projects filtered by type.
func (c ProjectClient) ListProjects(projectType matchlight.ProjectType) ([]matchlight.Project, error) {
	var projects []matchlight.Project
	var res apiResult
	if err := c.get("projects", &res); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res.Data, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// DeleteProject deletes an existing project within Matchlight.
func (c ProjectClient) DeleteProject(uploadToken string) error {
	if _, err := c.postRaw(fmt.Sprintf("project/%s/delete", uploadToken), nil); err != nil {
		return err
	}

	return nil
}

// GetProject retrieves an existing project from within Matchlight.
func (c ProjectClient) GetProject(uploadToken string) (matchlight.Project, error) {
	var project matchlight.Project
	if err := c.get(fmt.Sprintf("project/%s", uploadToken), &project); err != nil {
		return project, err
	}

	return project, nil
}

// EditProject edits an existing project within Matchlight.
func (c ProjectClient) EditProject(uploadToken, name string) (matchlight.Project, error) {
	// TODO (erik): Implement
	var project matchlight.Project

	return project, errors.New("unimplemented")
}
