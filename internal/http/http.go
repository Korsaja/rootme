package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"korsaj.io/rootme/pkg/models"
)

var (
	apiEndpoint = ""
	apiKey      = ""
	apiUID      = ""
)

func InitAPI(endpoint, key, uid string) {
	apiKey = key
	apiUID = uid
	apiEndpoint = endpoint
}

type profile struct {
	Name        string `json:"nom"`
	Score       string `json:"score"`
	Position    int    `json:"position"`
	Rank        string `json:"rang"`
	Validations []struct {
		IDChallenge string `json:"id_challenge"`
		IDRubrique  string `json:"id_rubrique"`
		Date        string `json:"date"`
	} `json:"validations"`
}

type validation struct {
	id   string
	date string
}

func (p *profile) validations() []validation {
	var validations = make([]validation, 0)
	for _, val := range p.Validations {
		v := validation{id: val.IDChallenge, date: val.Date}
		validations = append(validations, v)
	}
	return validations
}

type challenges []struct {
	Title      string `json:"titre,omitempty"`
	Rubric     string `json:"rubrique,omitempty"`
	SoursTitle string `json:"soustitre,omitempty"`
	Score      string `json:"score,omitempty"`
	Difficulty string `json:"difficulte,omitempty"`
}

func UserProfile(ctx context.Context) (*models.Profile, error) {
	var respProfile = &profile{}
	if err := request(ctx, userProfileURL(), respProfile); err != nil {
		return nil, fmt.Errorf("user profile %w", err)
	}

	prof := &models.Profile{
		NickName: respProfile.Name,
		Score:    respProfile.Score,
		Position: respProfile.Position,
		Rank:     respProfile.Rank,
		Solved:   make(map[string][]models.Task),
	}

	for _, val := range respProfile.validations() {
		var respChallenge = make(challenges, 0)
		if err := request(ctx, challengesURL(val.id), &respChallenge); err != nil {
			return nil, fmt.Errorf("challenge info error: %w", err)
		}
		if len(respChallenge) > 0 {
			c := respChallenge[0]
			prof.AddSolved(c.Rubric, models.NewTask(c.Title, c.Difficulty, val.date))
		}
	}

	return prof, nil
}

func userProfileURL() string {
	return fmt.Sprintf("%s/auteurs/%s&lang=en", apiEndpoint, apiUID)
}

func challengesURL(id string) string {
	return fmt.Sprintf("%s/challenges/%s", apiEndpoint, id)
}

func cookie() (string, string) {
	return "Cookie", fmt.Sprintf("api_key=%s", apiKey)
}

func request(ctx context.Context, url string, out interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set(cookie())
	return httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusTooManyRequests:
			return errors.New("many request")
		case http.StatusUnauthorized:
			return errors.New("invalid cookie")
		case http.StatusOK:
			if err = json.NewDecoder(resp.Body).Decode(out); err != nil {
				return fmt.Errorf("decode body %w", err)
			}
			return nil
		default:
			return fmt.Errorf("invalid request status %s", resp.Status)
		}
	})
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() { c <- f(http.DefaultClient.Do(req)) }()
	select {
	case <-ctx.Done():
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}
