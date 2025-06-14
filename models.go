package main

type App struct {
	Name          string         `json:"name"`
	Policies      []BackupPolicy `json:"policies"`
	LastDiskUsage float32        `json:"last_disk_usage"`
	LastBackup    string         `json:"last_backup"`
	LastError     string         `json:"last_error"`
}

type BackupPolicy struct {
	Path      string          `json:"path"`
	Interval  string          `json:"interval"`
	Retention RetentionPolicy `json:"retention"`
	Enabled   bool            `json:"enabled"`
	Exclude   []string        `json:"exclude"`
}

type RetentionPolicy struct {
	PerDay   int `json:"per_day"`
	PerWeek  int `json:"per_week"`
	PerMonth int `json:"per_month"`
}
