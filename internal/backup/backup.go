package backup

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

type githubRepo struct {
	URL  string `json:"clone_url"`
	Name string `json:"name"`
}

var Version, Timestamp string

type VersionInfo struct {
	Version, Timestamp string
}

func ExecuteBackup(cmd *cobra.Command, args []string) {
	if cmd.Flags().Lookup("version").Value.String() == "true" {
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Build Time: %s\n", Timestamp)

		os.Exit(0)
	}

	if len(args) == 0 {
		fmt.Printf("invalid username. app must be called with the GitHub username to backup")
		os.Exit(1)
	}

	repos, err := getRepos(args[0])
	if err != nil {
		fmt.Printf("could not retrieve list of GitHub repos for given user: %s", err)
		os.Exit(1)
	}

	err = backupRepos(repos)
	if err != nil {
		fmt.Printf("could not clone all GitHub repos for given user: %s", err)
		os.Exit(1)
	}
}

func backupRepos(repos []githubRepo) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, repo := range repos {
		cmd := exec.Command("git", "clone", "--mirror", repo.URL, repo.Name) //nolint:gosec

		_, err := os.Stat(repo.Name)
		if !os.IsNotExist(err) {
			cmd = exec.Command("git", "remote", "update")
			cmd.Dir = path.Join(workingDir, repo.Name)
		}

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func getRepos(username string) ([]githubRepo, error) {
	var githubRepos []githubRepo

	requestURL := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, requestURL, http.NoBody)
	if err != nil {
		return githubRepos, nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return githubRepos, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return githubRepos, nil
	}

	err = json.Unmarshal(body, &githubRepos)

	return githubRepos, err
}
