package models

type SettingsList struct {
	Meta                     interface{}        `json:"meta"`
	Logs                     Logs               `json:"logs"`
	Backups                  Backups            `json:"backups"`
	Smtp                     Smtp               `json:"smtp"`
	S3                       S3                 `json:"s3"`
	AdminAuthToken           AuthToken          `json:"adminAuthToken"`
	AdminPasswordResetToken  AuthToken          `json:"adminPasswordResetToken"`
	RecordAuthToken          AuthToken          `json:"recordAuthToken"`
	RecordPasswordResetToken AuthToken          `json:"recordPasswordResetToken"`
	RecordEmailChangeToken   AuthToken          `json:"recordEmailChangeToken"`
	RecordVerificationToken  AuthToken          `json:"recordVerificationToken"`
	GoogleAuth               OAuthConfiguration `json:"googleAuth"`
	FacebookAuth             FacebookAuth       `json:"facebookAuth"`
	GithubAuth               OAuthConfiguration `json:"githubAuth"`
	GitlabAuth               OAuthConfiguration `json:"gitlabAuth"`
	DiscordAuth              OAuthConfiguration `json:"discordAuth"`
	TwitterAuth              OAuthConfiguration `json:"twitterAuth"`
	MicrosoftAuth            OAuthConfiguration `json:"microsoftAuth"`
	SpotifyAuth              OAuthConfiguration `json:"spotifyAuth"`
}

type Logs struct {
	MaxDays int `json:"maxDays"`
}

type Backups struct {
	Cron        string `json:"cron"`
	CronMaxKeep int    `json:"cronMaxKeep"`
	S3          S3     `json:"s3"`
}

type Smtp struct {
	Enabled  bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Tls      bool   `json:"tls"`
}

type S3 struct {
	Enabled        bool   `json:"enabled"`
	Bucket         string `json:"bucket"`
	Region         string `json:"region"`
	Endpoint       string `json:"endpoint"`
	AccessKey      string `json:"accessKey"`
	Secret         string `json:"secret"`
	ForcePathStyle bool   `json:"forcePathStyle"`
}

type AuthToken struct {
	Secret   string `json:"secret"`
	Duration int    `json:"duration"`
}

type OAuthConfiguration struct {
	Enabled      bool   `json:"enabled"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type FacebookAuth struct {
	Enabled            bool `json:"enabled"`
	AllowRegistrations bool `json:"allowRegistrations"`
}
