package v2_test

import (
	"testing"

	matchlight "github.com/eriktate/go-matchlight"
	"github.com/eriktate/go-matchlight/v2"
)

func Test_Projects(t *testing.T) {
	// SETUP
	key := "2295e73367f741a485664c2c401f69c3"
	secret := "a309adb1346f4e3c8cf00a3cc21c280f"
	client := v2.NewClient(key, secret)
	pClient := v2.NewProjectClient(client)

	proj1 := matchlight.AddProjectReq{
		Name: "TestProject1",
		Type: matchlight.ProjectTypePII,
	}

	proj2 := matchlight.AddProjectReq{
		Name: "TestProject2",
		Type: matchlight.ProjectTypePII,
	}

	// RUN
	// create some projects
	res1, err := pClient.AddProject(proj1)
	if err != nil {
		t.Fatal(err)
	}

	res2, err := pClient.AddProject(proj2)
	if err != nil {
		t.Fatal(err)
	}

	// see if we can list them
	projects, err := pClient.ListProjects(matchlight.ProjectTypePII)
	if err != nil {
		t.Fatal(err)
	}

	// see if we can fetch them individually
	getProj1, err := pClient.GetProject(res1.UploadToken)
	if err != nil {
		t.Fatal(err)
	}

	getProj2, err := pClient.GetProject(res2.UploadToken)
	if err != nil {
		t.Fatal(err)
	}

	// cleanup projects
	if err := pClient.DeleteProject(res1.UploadToken); err != nil {
		t.Fatal(err)
	}

	if err := pClient.DeleteProject(res2.UploadToken); err != nil {
		t.Fatal(err)
	}

	// check delete worked
	cleanedProjects, err := pClient.ListProjects(matchlight.ProjectTypePII)
	if err != nil {
		t.Fatal(err)
	}

	// ASSERT
	if len(projects) < 2 {
		t.Fatal("Less than 2 projects returned to list operation")
	}

	if getProj1.Name != proj1.Name {
		t.Fatal("Fetched project 1 doesn't match input data")
	}

	if getProj2.Name != proj2.Name {
		t.Fatal("Fetched project 2 doesn't match input data")
	}

	if len(projects)-len(cleanedProjects) != 2 {
		t.Fatal("Projects count after cleanup incorrect")
	}
}
