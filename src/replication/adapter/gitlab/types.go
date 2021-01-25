package gitlab

// TokenResp is response of login.
type TokenResp struct {
	Token string `json:"token"`
}

// Project describes a project in GitLab
type Project struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FullPath        string `json:"path_with_namespace"`
	Visibility      string `json:"visibility"`
	RegistryEnabled bool   `json:"container_registry_enabled"`
}

// Repository describes a repository in GitLab
type Repository struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Location string `json:"location"`
}

// Tag describes a tag in GitLab
type Tag struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Location string `json:"location"`
}
