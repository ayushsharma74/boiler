package git

import (
    "fmt"
    "os"
    "os/exec"
)

// clones a git repository
func CloneRepository(repoURL string, projectPath string) error {

    cmd := exec.Command("git", "clone", repoURL, projectPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("error running git clone: %w", err)
    }
    return nil
}