package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var defaultKey = "default"

// Secret :
type Secret struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Action :
func Action(c *cli.Context) error {
	var home string
	if home = os.Getenv("HOME"); len(home) == 0 {
		return errors.New("env $HOME has not been set")
	}

	var pwd string
	if pwd = os.Getenv("PWD"); len(pwd) == 0 {
		return errors.New("env $PWD has not been set")
	}

	var gitOrgName string
	gitURL, _ := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if len(gitURL) == 0 {
		gitOrgName = path.Base(pwd)
	} else {
		gitOrgName = path.Base(path.Dir(string(gitURL)))
	}

	gitSecretReader, err := os.Open(filepath.Join(home, "$HOME/.git-secret.json"))
	if err != nil {
		return errors.New(".git-secret.json not found")
	}
	defer gitSecretReader.Close()

	gitSecret := map[string]Secret{}
	if err = json.NewDecoder(gitSecretReader).Decode(&gitSecret); err != nil {
		return err
	}

	var secret Secret
	var ok bool
	if secret, ok = gitSecret[gitOrgName]; ok {
		fmt.Printf("username=%s\n", secret.Username)
		fmt.Printf("password=%s\n", secret.Password)
		return nil
	}

	if secret, ok = gitSecret[defaultKey]; ok {
		fmt.Printf("username=%s\n", secret.Username)
		fmt.Printf("password=%s\n", secret.Password)
		return nil
	}
	return errors.New("Failed to find the default credential. see https://github.com/sadayuki-matsuno/git-credential-substitute")
}
