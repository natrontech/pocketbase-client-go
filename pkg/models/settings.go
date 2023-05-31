package models

type SettingsList struct {
	Meta                     Meta               `json:"meta"`
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

type Meta struct {
	AppName                    string        `json:"appName"`
	AppUrl                     string        `json:"appUrl"`
	HideControls               *bool         `json:"hideControls"`
	SenderName                 string        `json:"senderName"`
	SenderAddress              string        `json:"senderAddress"`
	VerificationTemplate       EmailTemplate `json:"verificationTemplate"`
	ResetPasswordTemplate      EmailTemplate `json:"resetPasswordTemplate"`
	ConfirmEmailChangeTemplate EmailTemplate `json:"confirmEmailChangeTemplate"`
}

type EmailTemplate struct {
	ActionUrl string `json:"actionUrl"`
	Body      string `json:"body"`
	Subject   string `json:"subject"`
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

type SettingsUpdateRequest struct {
	Meta                     *Meta               `json:"meta,omitempty"`
	Logs                     *Logs               `json:"logs,omitempty"`
	Backups                  *Backups            `json:"backups,omitempty"`
	Smtp                     *Smtp               `json:"smtp,omitempty"`
	S3                       *S3                 `json:"s3,omitempty"`
	AdminAuthToken           *AuthToken          `json:"adminAuthToken,omitempty"`
	AdminPasswordResetToken  *AuthToken          `json:"adminPasswordResetToken,omitempty"`
	RecordAuthToken          *AuthToken          `json:"recordAuthToken,omitempty"`
	RecordPasswordResetToken *AuthToken          `json:"recordPasswordResetToken,omitempty"`
	RecordEmailChangeToken   *AuthToken          `json:"recordEmailChangeToken,omitempty"`
	RecordVerificationToken  *AuthToken          `json:"recordVerificationToken,omitempty"`
	GoogleAuth               *OAuthConfiguration `json:"googleAuth,omitempty"`
	FacebookAuth             *FacebookAuth       `json:"facebookAuth,omitempty"`
	GithubAuth               *OAuthConfiguration `json:"githubAuth,omitempty"`
	GitlabAuth               *OAuthConfiguration `json:"gitlabAuth,omitempty"`
	DiscordAuth              *OAuthConfiguration `json:"discordAuth,omitempty"`
	TwitterAuth              *OAuthConfiguration `json:"twitterAuth,omitempty"`
	MicrosoftAuth            *OAuthConfiguration `json:"microsoftAuth,omitempty"`
	SpotifyAuth              *OAuthConfiguration `json:"spotifyAuth,omitempty"`
}

type SettingsS3StorageTestRequest struct {
	Filesystem string `json:"filesystem"`
}

type SettingsResponse struct {
	Code    *string `json:"code"`
	Message *string `json:"message"`
	Data    *any    `json:"data"`
}

type SettingsSendTestEmailRequest struct {
	Email    string `json:"email"`
	Template string `json:"template"`
}
