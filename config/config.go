package config

type Config struct {
	path string

	Mode     string `json:"mode"`
	Database string `json:"database"`

	Session sessionConfig `json:"session"`

	User    userConfig    `json:"user"`
	Section sectionConfig `json:"section"`
}

type sessionConfig struct {
	IdentificatorLength uint16 `json:"identificator_length"`
	DaysTimeout         uint16 `json:"timeout"`
}

type userConfig struct {
	MinUsernameLength uint16 `json:"min_username_length"`
	MaxUsernameLength uint16 `json:"max_username_length"`

	MinPasswordLength uint16 `json:"min_password_length"`
	MaxPasswordLength uint16 `json:"max_password_length"`
}

type sectionConfig struct {
	MinTitleLength uint16 `json:"min_title_length"`
	MaxTitleLength uint16 `json:"max_title_length"`

	MinBodyLength uint16 `json:"min_body_length"`
	MaxBodyLength uint16 `json:"max_body_length"`
}
