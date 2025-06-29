package models

import (
	"time"

	"gorm.io/gorm"
)

type App struct {
	gorm.Model
	Name          string         `json:"name"`
	TmpPath       string         `json:"tmp_path" gorm:"default:'/tmp'"`
	Policies      []BackupPolicy `json:"policies" gorm:"foreignKey:AppID"`
	Source        AppSource      `json:"source" gorm:"serializer:json"`
	LastDiskUsage float32        `json:"last_disk_usage"`
	LastBackup    string         `json:"last_backup"`
	LastError     string         `json:"last_error"`
}

type BackupPolicy struct {
	gorm.Model
	AppID          uint            `json:"app_id"`
	Name           string          `json:"name"`
	Path           string          `json:"path"`
	Type           PolicyType      `json:"type"`
	Configuration  map[string]any  `gorm:"serializer:json" json:"configuration"`
	Interval       uint            `json:"interval"`
	Retention      RetentionPolicy `json:"retention" gorm:"serializer:json"`
	Enabled        bool            `json:"enabled"`
	Exclude        []string        `json:"exclude" gorm:"type:json"`
	Backups        []Backup        `json:"backups"`
	NextBackupDate time.Time       `json:"next_backup_date"`
	App            App             `json:"app" gorm:"foreignKey:AppID"`
}

type RetentionPolicy struct {
	PerDay   int `json:"per_day"`
	PerWeek  int `json:"per_week"`
	PerMonth int `json:"per_month"`
}
type AppSource struct {
	Type          SourceType     `json:"type"`
	Configuration map[string]any `json:"configuration"`
}
type Backup struct {
	gorm.Model
	AppID          uint `json:"app_id"`
	BackupPolicyID uint `json:"backup_policy_id"`
}

type PolicyType string
type SourceType string

const (
	PolicyTypeDatabase   PolicyType = "database"
	PolicyTypeFileSystem PolicyType = "filesystem"
)
const (
	SourceTypeLocal SourceType = "local"
	SourceTypeSSH   SourceType = "ssh"
)
