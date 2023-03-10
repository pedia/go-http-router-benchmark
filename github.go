package bench

// http://developer.github.com/v3/
var githubAPI = []Route{
	// OAuth Authorizations
	{"GET", "/authorizations", false},
	{"GET", "/authorizations/:id", false},
	{"POST", "/authorizations", true},
	// {"PUT", "/authorizations/clients/:client_id", false},
	// {"PATCH", "/authorizations/:id", false},
	{"DELETE", "/authorizations/:id", true},
	{"GET", "/applications/:client_id/tokens/:access_token", false},
	{"DELETE", "/applications/:client_id/tokens", false},
	{"DELETE", "/applications/:client_id/tokens/:access_token", true},

	// Activity
	{"GET", "/events", false},
	{"GET", "/repos/:owner/:repo/events", false},
	{"GET", "/networks/:owner/:repo/events", false},
	{"GET", "/orgs/:org/events", false},
	{"GET", "/users/:user/received_events", false},
	{"GET", "/users/:user/received_events/public", false},
	{"GET", "/users/:user/events", false},
	{"GET", "/users/:user/events/public", false},
	{"GET", "/users/:user/events/orgs/:org", false},
	{"GET", "/feeds", false},
	{"GET", "/notifications", false},
	{"GET", "/repos/:owner/:repo/notifications", false},
	{"PUT", "/notifications", true},
	{"PUT", "/repos/:owner/:repo/notifications", true},
	{"GET", "/notifications/threads/:id", false},
	// {"PATCH", "/notifications/threads/:id", false},
	{"GET", "/notifications/threads/:id/subscription", false},
	{"PUT", "/notifications/threads/:id/subscription", true},
	{"DELETE", "/notifications/threads/:id/subscription", true},
	{"GET", "/repos/:owner/:repo/stargazers", false},
	{"GET", "/users/:user/starred", false},
	{"GET", "/user/starred", false},
	{"GET", "/user/starred/:owner/:repo", false},
	{"PUT", "/user/starred/:owner/:repo", true},
	{"DELETE", "/user/starred/:owner/:repo", true},
	{"GET", "/repos/:owner/:repo/subscribers", false},
	{"GET", "/users/:user/subscriptions", false},
	{"GET", "/user/subscriptions", false},
	{"GET", "/repos/:owner/:repo/subscription", false},
	{"PUT", "/repos/:owner/:repo/subscription", true},
	{"DELETE", "/repos/:owner/:repo/subscription", true},
	{"GET", "/user/subscriptions/:owner/:repo", false},
	{"PUT", "/user/subscriptions/:owner/:repo", true},
	{"DELETE", "/user/subscriptions/:owner/:repo", true},

	// Gists
	{"GET", "/users/:user/gists", false},
	{"GET", "/gists", false},
	// {"GET", "/gists/public", false},
	// {"GET", "/gists/starred", false},
	{"GET", "/gists/:id", false},
	{"POST", "/gists", true},
	// {"PATCH", "/gists/:id", false},
	{"PUT", "/gists/:id/star", false},
	{"DELETE", "/gists/:id/star", true},
	{"GET", "/gists/:id/star", true},
	{"POST", "/gists/:id/forks", false},
	{"DELETE", "/gists/:id", true},

	// Git Data
	{"GET", "/repos/:owner/:repo/git/blobs/:sha", false},
	{"POST", "/repos/:owner/:repo/git/blobs", false},
	{"GET", "/repos/:owner/:repo/git/commits/:sha", false},
	{"POST", "/repos/:owner/:repo/git/commits", false},
	// {"GET", "/repos/:owner/:repo/git/refs/*ref", false},
	{"GET", "/repos/:owner/:repo/git/refs", false},
	{"POST", "/repos/:owner/:repo/git/refs", true},
	// {"PATCH", "/repos/:owner/:repo/git/refs/*ref", false},
	// {"DELETE", "/repos/:owner/:repo/git/refs/*ref", false},
	{"GET", "/repos/:owner/:repo/git/tags/:sha", false},
	{"POST", "/repos/:owner/:repo/git/tags", false},
	{"GET", "/repos/:owner/:repo/git/trees/:sha", false},
	{"POST", "/repos/:owner/:repo/git/trees", false},

	// Issues
	{"GET", "/issues", false},
	{"GET", "/user/issues", false},
	{"GET", "/orgs/:org/issues", false},
	{"GET", "/repos/:owner/:repo/issues", false},
	{"GET", "/repos/:owner/:repo/issues/:number", false},
	{"POST", "/repos/:owner/:repo/issues", true},
	// {"PATCH", "/repos/:owner/:repo/issues/:number", false},
	{"GET", "/repos/:owner/:repo/assignees", false},
	{"GET", "/repos/:owner/:repo/assignees/:assignee", false},
	{"GET", "/repos/:owner/:repo/issues/:number/comments", false},
	// {"GET", "/repos/:owner/:repo/issues/comments", false},
	// {"GET", "/repos/:owner/:repo/issues/comments/:id", false},
	{"POST", "/repos/:owner/:repo/issues/:number/comments", true},
	// {"PATCH", "/repos/:owner/:repo/issues/comments/:id", false},
	// {"DELETE", "/repos/:owner/:repo/issues/comments/:id", false},
	{"GET", "/repos/:owner/:repo/issues/:number/events", false},
	// {"GET", "/repos/:owner/:repo/issues/events", false},
	// {"GET", "/repos/:owner/:repo/issues/events/:id", false},
	{"GET", "/repos/:owner/:repo/labels", false},
	{"GET", "/repos/:owner/:repo/labels/:name", false},
	{"POST", "/repos/:owner/:repo/labels", true},
	// {"PATCH", "/repos/:owner/:repo/labels/:name", false},
	{"DELETE", "/repos/:owner/:repo/labels/:name", true},
	{"GET", "/repos/:owner/:repo/issues/:number/labels", false},
	{"POST", "/repos/:owner/:repo/issues/:number/labels", true},
	{"DELETE", "/repos/:owner/:repo/issues/:number/labels/:name", false},
	{"PUT", "/repos/:owner/:repo/issues/:number/labels", true},
	{"DELETE", "/repos/:owner/:repo/issues/:number/labels", true},
	{"GET", "/repos/:owner/:repo/milestones/:number/labels", false},
	{"GET", "/repos/:owner/:repo/milestones", false},
	{"GET", "/repos/:owner/:repo/milestones/:number", false},
	{"POST", "/repos/:owner/:repo/milestones", true},
	// {"PATCH", "/repos/:owner/:repo/milestones/:number", false},
	{"DELETE", "/repos/:owner/:repo/milestones/:number", true},

	// Miscellaneous
	{"GET", "/emojis", false},
	{"GET", "/gitignore/templates", false},
	{"GET", "/gitignore/templates/:name", false},
	{"POST", "/markdown", false},
	{"POST", "/markdown/raw", false},
	{"GET", "/meta", false},
	{"GET", "/rate_limit", false},

	// Organizations
	{"GET", "/users/:user/orgs", false},
	{"GET", "/user/orgs", false},
	{"GET", "/orgs/:org", false},
	// {"PATCH", "/orgs/:org", false},
	{"GET", "/orgs/:org/members", false},
	{"GET", "/orgs/:org/members/:user", false},
	{"DELETE", "/orgs/:org/members/:user", true},
	{"GET", "/orgs/:org/public_members", false},
	{"GET", "/orgs/:org/public_members/:user", false},
	{"PUT", "/orgs/:org/public_members/:user", true},
	{"DELETE", "/orgs/:org/public_members/:user", true},
	{"GET", "/orgs/:org/teams", false},
	{"GET", "/teams/:id", false},
	{"POST", "/orgs/:org/teams", true},
	// {"PATCH", "/teams/:id", false},
	{"DELETE", "/teams/:id", true},
	{"GET", "/teams/:id/members", false},
	{"GET", "/teams/:id/members/:user", false},
	{"PUT", "/teams/:id/members/:user", true},
	{"DELETE", "/teams/:id/members/:user", true},
	{"GET", "/teams/:id/repos", false},
	{"GET", "/teams/:id/repos/:owner/:repo", false},
	{"PUT", "/teams/:id/repos/:owner/:repo", true},
	{"DELETE", "/teams/:id/repos/:owner/:repo", true},
	{"GET", "/user/teams", false},

	// Pull Requests
	{"GET", "/repos/:owner/:repo/pulls", false},
	{"GET", "/repos/:owner/:repo/pulls/:number", false},
	{"POST", "/repos/:owner/:repo/pulls", true},
	// {"PATCH", "/repos/:owner/:repo/pulls/:number", false},
	{"GET", "/repos/:owner/:repo/pulls/:number/commits", false},
	{"GET", "/repos/:owner/:repo/pulls/:number/files", false},
	{"GET", "/repos/:owner/:repo/pulls/:number/merge", false},
	{"PUT", "/repos/:owner/:repo/pulls/:number/merge", true},
	{"GET", "/repos/:owner/:repo/pulls/:number/comments", false},
	// {"GET", "/repos/:owner/:repo/pulls/comments", false},
	// {"GET", "/repos/:owner/:repo/pulls/comments/:number", false},
	{"PUT", "/repos/:owner/:repo/pulls/:number/comments", true},
	// {"PATCH", "/repos/:owner/:repo/pulls/comments/:number", false},
	// {"DELETE", "/repos/:owner/:repo/pulls/comments/:number", false},

	// Repositories
	{"GET", "/user/repos", false},
	{"GET", "/users/:user/repos", false},
	{"GET", "/orgs/:org/repos", false},
	{"GET", "/repositories", false},
	{"POST", "/user/repos", true},
	{"POST", "/orgs/:org/repos", true},
	{"GET", "/repos/:owner/:repo", false},
	// {"PATCH", "/repos/:owner/:repo", false},
	{"GET", "/repos/:owner/:repo/contributors", false},
	{"GET", "/repos/:owner/:repo/languages", false},
	{"GET", "/repos/:owner/:repo/teams", false},
	{"GET", "/repos/:owner/:repo/tags", false},
	{"GET", "/repos/:owner/:repo/branches", false},
	{"GET", "/repos/:owner/:repo/branches/:branch", false},
	{"DELETE", "/repos/:owner/:repo", true},
	{"GET", "/repos/:owner/:repo/collaborators", false},
	{"GET", "/repos/:owner/:repo/collaborators/:user", false},
	{"PUT", "/repos/:owner/:repo/collaborators/:user", true},
	{"DELETE", "/repos/:owner/:repo/collaborators/:user", true},
	{"GET", "/repos/:owner/:repo/comments", false},
	{"GET", "/repos/:owner/:repo/commits/:sha/comments", false},
	{"POST", "/repos/:owner/:repo/commits/:sha/comments", true},
	{"GET", "/repos/:owner/:repo/comments/:id", false},
	// {"PATCH", "/repos/:owner/:repo/comments/:id", false},
	{"DELETE", "/repos/:owner/:repo/comments/:id", true},
	{"GET", "/repos/:owner/:repo/commits", false},
	{"GET", "/repos/:owner/:repo/commits/:sha", false},
	{"GET", "/repos/:owner/:repo/readme", false},
	// {"GET", "/repos/:owner/:repo/contents/*path", false},
	// {"PUT", "/repos/:owner/:repo/contents/*path", false},
	// {"DELETE", "/repos/:owner/:repo/contents/*path", false},
	// {"GET", "/repos/:owner/:repo/:archive_format/:ref", false},
	{"GET", "/repos/:owner/:repo/keys", false},
	{"GET", "/repos/:owner/:repo/keys/:id", false},
	{"POST", "/repos/:owner/:repo/keys", true},
	// {"PATCH", "/repos/:owner/:repo/keys/:id", false},
	{"DELETE", "/repos/:owner/:repo/keys/:id", true},
	{"GET", "/repos/:owner/:repo/downloads", false},
	{"GET", "/repos/:owner/:repo/downloads/:id", false},
	{"DELETE", "/repos/:owner/:repo/downloads/:id", true},
	{"GET", "/repos/:owner/:repo/forks", false},
	{"POST", "/repos/:owner/:repo/forks", true},
	{"GET", "/repos/:owner/:repo/hooks", false},
	{"GET", "/repos/:owner/:repo/hooks/:id", false},
	{"POST", "/repos/:owner/:repo/hooks", true},
	// {"PATCH", "/repos/:owner/:repo/hooks/:id", false},
	{"POST", "/repos/:owner/:repo/hooks/:id/tests", false},
	{"DELETE", "/repos/:owner/:repo/hooks/:id", true},
	{"POST", "/repos/:owner/:repo/merges", false},
	{"GET", "/repos/:owner/:repo/releases", false},
	{"GET", "/repos/:owner/:repo/releases/:id", false},
	{"POST", "/repos/:owner/:repo/releases", true},
	// {"PATCH", "/repos/:owner/:repo/releases/:id", false},
	{"DELETE", "/repos/:owner/:repo/releases/:id", true},
	{"GET", "/repos/:owner/:repo/releases/:id/assets", false},
	{"GET", "/repos/:owner/:repo/stats/contributors", false},
	{"GET", "/repos/:owner/:repo/stats/commit_activity", false},
	{"GET", "/repos/:owner/:repo/stats/code_frequency", false},
	{"GET", "/repos/:owner/:repo/stats/participation", false},
	{"GET", "/repos/:owner/:repo/stats/punch_card", false},
	{"GET", "/repos/:owner/:repo/statuses/:ref", false},
	{"POST", "/repos/:owner/:repo/statuses/:ref", true},

	// Search
	{"GET", "/search/repositories", false},
	{"GET", "/search/code", false},
	{"GET", "/search/issues", false},
	{"GET", "/search/users", false},
	{"GET", "/legacy/issues/search/:owner/:repository/:state/:keyword", false},
	{"GET", "/legacy/repos/search/:keyword", false},
	{"GET", "/legacy/user/search/:keyword", false},
	{"GET", "/legacy/user/email/:email", false},

	// Users
	{"GET", "/users/:user", false},
	{"GET", "/user", false},
	// {"PATCH", "/user", false},
	{"GET", "/users", false},
	{"GET", "/user/emails", false},
	{"POST", "/user/emails", true},
	{"DELETE", "/user/emails", true},
	{"GET", "/users/:user/followers", false},
	{"GET", "/user/followers", false},
	{"GET", "/users/:user/following", false},
	{"GET", "/user/following", false},
	{"GET", "/user/following/:user", false},
	{"GET", "/users/:user/following/:target_user", false},
	{"PUT", "/user/following/:user", true},
	{"DELETE", "/user/following/:user", true},
	{"GET", "/users/:user/keys", false},
	{"GET", "/user/keys", false},
	{"GET", "/user/keys/:id", false},
	{"POST", "/user/keys", true},
	// {"PATCH", "/user/keys/:id", false},
	{"DELETE", "/user/keys/:id", true},
}
