package components

import "github.com/anyappinc/actions-discord/app/module/githubActions"

var GithubActionsClient *githubActions.GithubActionsClient

type githubActionsComponent struct{}

var GithubActions = &githubActionsComponent{}

func (githubActionsComponent) GetGithubActionsInfo() *githubActions.GithubActionsClient {
	return GithubActionsClient
}
