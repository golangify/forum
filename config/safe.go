package config

func (c *Config) safe() {
	if c.Mode == "" {
		c.Mode = "debug"
	}
	if c.Database == "" {
		c.Database = "database.db"
	}
	if c.Session.IdentificatorLength == 0 {
		c.Session.IdentificatorLength = 255
	}
	if c.Session.DaysTimeout == 0 {
		c.Session.DaysTimeout = 365
	}

	// user
	if c.User.MinUsernameLength == 0 {
		c.User.MinUsernameLength = 1
	}
	if c.User.MaxUsernameLength == 0 || c.User.MaxUsernameLength < c.User.MinUsernameLength {
		c.User.MaxUsernameLength = c.User.MinUsernameLength + 20
	}
	if c.User.MinPasswordLength == 0 {
		c.User.MinPasswordLength = 5
	}
	if c.User.MaxPasswordLength == 0 || c.User.MaxPasswordLength < c.User.MinPasswordLength {
		c.User.MaxPasswordLength += c.User.MinPasswordLength + 50
	}
	if c.User.MaxPasswordLength > 72 {
		c.User.MaxPasswordLength = 72 // bcrypt limit
	}

	// section
	if c.Section.MinTitleLength == 0 {
		c.Section.MinTitleLength = 1
	}
	if c.Section.MaxTitleLength == 0 || c.Section.MaxTitleLength < c.Section.MinTitleLength {
		c.Section.MaxTitleLength = c.Section.MinTitleLength + 100
	}
	if c.Section.MinBodyLength == 0 {
		c.Section.MinBodyLength = 1
	}
	if c.Section.MaxBodyLength == 0 || c.Section.MaxBodyLength < c.Section.MinBodyLength {
		c.Section.MaxBodyLength = c.Section.MinBodyLength + 4000
	}
}
