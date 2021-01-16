package git

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GithubHook is webhook for Github.com.
type GithubHook struct{}

type ghRelease struct {
	Action  string `json:"action"`
	Release struct {
		TagName string      `json:"tag_name"`
		Name    interface{} `json:"name"`
	} `json:"release"`
}

type ghPush struct {
	Ref string `json:"ref"`
}

// DoesHandle satisfies hookHandler.
func (g GithubHook) DoesHandle(h http.Header) bool {
	userAgent := h.Get("User-Agent")

	// GitHub always uses a user-agent like "GitHub-Hookshot/<id>"
	if userAgent != "" && strings.HasPrefix(userAgent, "GitHub-Hookshot") {
		return true
	}
	return false
}

// Handle satisfies hookHandler.
func (g GithubHook) Handle(w http.ResponseWriter, r *http.Request, repo *Repo) (int, error) {
	if r.Method != "POST" {
		return http.StatusMethodNotAllowed, errors.New("the request had an invalid method")
	}

	// read full body - required for signature
	body, err := ioutil.ReadAll(r.Body)

	err = g.handleSignature(r, body, repo.Hook.Secret)
	if err != nil {
		return http.StatusBadRequest, err
	}

	event := r.Header.Get("X-Github-Event")
	if event == "" {
		return http.StatusBadRequest, errors.New("the 'X-Github-Event' header is required but was missing")
	}

	switch event {
	case "ping":
		w.Write([]byte("pong"))
	case "push":
		err = g.handlePush(body, repo)
		if !hookIgnored(err) && err != nil {
			return http.StatusBadRequest, err
		}
	case "release":
		err = g.handleRelease(body, repo)
		if err != nil {
			return http.StatusBadRequest, err
		}

	// return 400 if we do not handle the event type.
	// This is to visually show the user a configuration error in the GH ui.
	default:
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, err
}

// Check for an optional signature in the request
// if it is signed, verify the signature.
func (g GithubHook) handleSignature(r *http.Request, body []byte, secret string) error {
	signature := r.Header.Get("X-Hub-Signature")
	if signature != "" {
		if secret == "" {
			Logger().Print("Unable to verify request signature. Secret not set in caddyfile!\n")
		} else {
			mac := hmac.New(sha1.New, []byte(secret))
			mac.Write(body)
			expectedMac := hex.EncodeToString(mac.Sum(nil))

			if signature[5:] != expectedMac {
				return errors.New("could not verify request signature. The signature is invalid")
			}
		}
	}

	return nil
}

func (g GithubHook) handlePush(body []byte, repo *Repo) error {
	var push ghPush

	err := json.Unmarshal(body, &push)
	if err != nil {
		return err
	}

	// extract the branch being pushed from the ref string
	// and if it matches with our locally tracked one, pull.
	refSlice := strings.Split(push.Ref, "/")
	if len(refSlice) != 3 {
		return errors.New("the push request contained an invalid reference string")
	}

	branch := refSlice[2]

	if branch != repo.Branch {
		return hookIgnoredError{hookType: hookName(g), err: fmt.Errorf("found different branch %v", branch)}
	}

	Logger().Println("Received pull notification for the tracking branch, updating...")
	repo.Pull()
	return nil
}

func (g GithubHook) handleRelease(body []byte, repo *Repo) error {
	var release ghRelease

	err := json.Unmarshal(body, &release)
	if err != nil {
		return err
	}

	if release.Release.TagName == "" {
		return errors.New("the release request contained an invalid TagName")
	}

	Logger().Printf("Received new release '%s'. -> Updating local repository to this release.\n", release.Release.Name)

	// Update the local branch to the release tag name
	// this will pull the release tag.
	repo.Branch = release.Release.TagName
	repo.Pull()

	return nil
}
