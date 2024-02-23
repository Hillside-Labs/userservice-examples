package main

type Prefs struct {
	// Visual Preferences
	Theme       string `json:"theme"`
	FontSize    int    `json:"fontSize"`
	ColorScheme string `json:"colorScheme"`

	// Behavior Preferences
	Language          string `json:"language"`
	AutoSave          bool   `json:"autoSave"`
	ShowNotifications bool   `json:"showNotifications"`

	// Feature-specific Preferences
	CachePages CacheConfig `json:"cacheConfig"`
}

type CacheConfig struct {
	Enabled   bool `json:"enabled"`
	Threshold int  `json:"threshold"`
}
