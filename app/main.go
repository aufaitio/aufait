package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/quantumew/data-access/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

// CLI manages the configuration of Au Fait
type CLI struct {
	listenerURL *url.URL
}

// Package repositories package file
type Package struct {
	Name            string            `json:"name"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

// LockFile represents a lock file
type LockFile struct {
	Dependencies map[string]LockDependency `json:"dependencies"`
}

// LockDependency dependency nested in lock file
type LockDependency struct {
	Version string `json:"version"`
}

// NewCLI creates CLI for configuring Au Fait
func NewCLI(listenerURLStr string) (CLI, error) {
	listenerURL, err := url.Parse(listenerURLStr)

	if err != nil {
		err = fmt.Errorf("Invalid listener URL provided %s", listenerURLStr)
	}

	return CLI{listenerURL}, err
}

// ConfigureLocalRepository config list of paths to locally cloned repositories
func (cli CLI) ConfigureLocalRepository(repoPath string, remote string, branch string, remoteName string) error {
	var (
		pkg     Package
		lock    LockFile
		rawLock []byte
	)

	// Load repo files and configure
	lockFilePath := path.Join(repoPath, "package-lock.json")
	pkgFilePath := path.Join(repoPath, "package.json")
	configFilePath := path.Join(repoPath, ".aufait.json")

	rawPkg, err := ioutil.ReadFile(pkgFilePath)

	if err != nil {
		return fmt.Errorf("Failed to load package.json for %s", repoPath)
	}

	pkgErr := json.Unmarshal(rawPkg, &pkg)

	if pkgErr != nil {
		return fmt.Errorf("Failed to marshal package.json, bad JSON, for %s", repoPath)
	}

	if _, err := os.Stat(lockFilePath); err == nil {
		rawLock, err = ioutil.ReadFile(lockFilePath)
		json.Unmarshal(rawLock, &lock)
	}

	if len(remote) < 1 {
		rawRemote, err := exec.Command("git", "-C", repoPath, "remote", "get-url", remoteName).Output()

		if err != nil {
			return fmt.Errorf("Failed to retrieve remote for git repository %s", repoPath)
		}

		remote = strings.Trim(string(rawRemote[:]), " \n")
	}

	repo, err := cli.buildRepositoryConfig(pkg, lock, branch, remote)

	if err != nil {
		return err
	}

	jsonPayload, err := json.Marshal(repo)

	if err != nil {
		return fmt.Errorf("Failed to marshal JSON for Au Fait config %s", repoPath)
	}

	err = ioutil.WriteFile(configFilePath, jsonPayload, 0644)

	if err != nil {
		return fmt.Errorf("Failed to write aufait config file for %s", repoPath)
	}

	return cli.ConfigureRepositories([]*models.Repository{repo})
}

// ConfigureRepositories sets up provided repositories for Au Fait
func (cli CLI) ConfigureRepositories(repositoryList []*models.Repository) error {
	if len(repositoryList) > 0 {
		client := &http.Client{}
		json, err := json.Marshal(repositoryList)

		if err != nil {
			return err
		}

		body := bytes.NewReader(json)
		relPath, _ := url.Parse("/repositories")
		req, err := http.NewRequest(http.MethodPatch, cli.listenerURL.ResolveReference(relPath).String(), body)

		if err != nil {
			return err
		}

		response, err := client.Do(req)

		if err != nil {
			return err
		}

		if response.StatusCode > 399 {
			bodyBytes := new(bytes.Buffer)
			bodyBytes.ReadFrom(response.Body)

			return fmt.Errorf(
				"Error configuring repositories, Status Code: %d, Response Body: %s",
				response.StatusCode,
				bodyBytes.String(),
			)
		}
	}

	return nil
}

func (cli CLI) buildRepositoryConfig(pkg Package, lockFile LockFile, branch string, remote string) (*models.Repository, error) {
	var repository models.Repository
	var dependencies []models.Dependency

	for depName, depVersion := range pkg.Dependencies {
		dependencies = append(dependencies, cli.processDependency(depName, depVersion, lockFile))
	}

	repository.Name = pkg.Name
	repository.Dependencies = dependencies
	repository.Config = models.Config{Branch: branch, Remote: remote}

	return &repository, nil
}

func (cli CLI) processDependency(depName string, depVersion string, lockFile LockFile) models.Dependency {
	installedVersion := depVersion

	for key, dep := range lockFile.Dependencies {
		if key == depName {
			installedVersion = dep.Version
			break
		}
	}
	return models.Dependency{Installed: installedVersion, Name: depName, Semver: depVersion}
}
