package cst

import (
	"github.com/google/uuid"
)

var (
	OaiDeviceId = uuid.NewString()
	ChatAskMap  = map[string]map[string]string{
		"backend-anon": map[string]string{
			"requirementsUrl": "https://chat.openai.com/backend-anon/sentinel/chat-requirements",
			"askUrl":          "https://chat.openai.com/backend-anon/conversation",
		},
		"backend-api": map[string]string{
			"requirementsUrl": "https://chat.openai.com/backend-api/sentinel/chat-requirements",
			"askUrl":          "https://chat.openai.com/backend-api/conversation",
		},
	}
)

const (
	OaiLanguage = "en-US"

	ChatOriginUrl       = "https://chat.openai.com"
	ChatRefererUrl      = "https://chat.openai.com/"
	ChatCsrfUrl         = "https://chat.openai.com/api/auth/csrf"
	ChatPromptLoginUrl  = "https://chat.openai.com/api/auth/signin/login-web?prompt=login&screen_hint=login"
	Auth0Url            = "https://auth0.openai.com"
	LoginUsernameUrl    = "https://auth0.openai.com/u/login/identifier?state="
	LoginPasswordUrl    = "https://auth0.openai.com/u/login/password?state="
	ChatAuthSessionUrl  = "https://chat.openai.com/api/auth/session"
	OauthTokenUrl       = "https://auth0.openai.com/oauth/token"
	OauthTokenRevokeUrl = "https://auth0.openai.com/oauth/revoke"

	PlatformAuthClientID     = "DRivsnm2Mu42T3KOpqdtwB3NYviHYzwD"
	PlatformAuth0Client      = "eyJuYW1lIjoiYXV0aDAtc3BhLWpzIiwidmVyc2lvbiI6IjEuMjEuMCJ9"
	PlatformAuthAudience     = "https://api.openai.com/v1"
	PlatformAuthRedirectURL  = "https://platform.openai.com/auth/callback"
	PlatformAuthScope        = "openid email profile offline_access model.request model.read organization.read organization.write"
	PlatformAuthResponseType = "code"
	PlatformAuthGrantType    = "authorization_code"
	PlatformAuth0Url         = "https://auth0.openai.com/authorize?"
	PlatformAuth0LogoutUrl   = "https://auth0.openai.com/v2/logout?returnTo=https%3A%2F%2Fplatform.openai.com%2Floggedout&client_id=DRivsnm2Mu42T3KOpqdtwB3NYviHYzwD&auth0Client=eyJuYW1lIjoiYXV0aDAtc3BhLWpzIiwidmVyc2lvbiI6IjEuMjEuMCJ9"
	DashboardLoginUrl        = "https://api.openai.com/dashboard/onboarding/login"
)
