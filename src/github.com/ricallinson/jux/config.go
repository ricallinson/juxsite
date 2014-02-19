package jux

type Config struct {
	// Page level configuration keys.
	Page struct {
		BaseUrl     string
		Name        string
		Description string
		Lang        string
		Direction   string
	}

	// Application level configuration keys.
	Defaults struct {
		Env         string
		Debug       bool
		PublicTheme string
		AdminTheme  string
	}

	// Map of all usable layouts.
	Layouts map[string]map[string][]string
}
