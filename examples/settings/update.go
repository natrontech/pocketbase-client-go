package examples_settings

import (
	"github.com/natrontech/pocketbase-client-go/examples"
	"github.com/natrontech/pocketbase-client-go/pkg/models"
)

func UpdateSettings() {
	meata := models.Meta{
		AppName:       "test",
		AppUrl:        "https://test.com",
		SenderName:    "test",
		SenderAddress: "test@test.com",
		ConfirmEmailChangeTemplate: models.EmailTemplate{
			Subject:   "Confirm your {APP_NAME} new email address",
			Body:      "<p>Hello,</p>\n<p>Click on the button below to confirm your new email address.</p>\n<p>\n  <a class=\"btn\" href=\"{ACTION_URL}\" target=\"_blank\" rel=\"noopener\">Confirm new email</a>\n</p>\n<p><i>If you didn't ask to change your email address, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
			ActionUrl: "{APP_URL}/_/#/auth/confirm-email-change/{TOKEN}",
		},
		ResetPasswordTemplate: models.EmailTemplate{
			Subject:   "Reset your {APP_NAME} password",
			Body:      "<p>Hello,</p>\n<p>Click on the button below to reset your password.</p>\n<p>\n  <a class=\"btn\" href=\"{ACTION_URL}\" target=\"_blank\" rel=\"noopener\">Reset password</a>\n</p>\n<p><i>If you didn't ask to reset your password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
			ActionUrl: "{APP_URL}/_/#/auth/reset-password/{TOKEN}",
		},
		VerificationTemplate: models.EmailTemplate{
			Subject:   "Verify your {APP_NAME} email",
			Body:      "<p>Hello,</p>\n<p>Thank you for joining us at {APP_NAME}.</p>\n<p>Click on the button below to verify your email address.</p>\n<p>\n  <a class=\"btn\" href=\"{ACTION_URL}\" target=\"_blank\" rel=\"noopener\">Verify</a>\n</p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
			ActionUrl: "{APP_URL}/_/#/auth/verify/{TOKEN}",
		},
	}
	logs := models.Logs{
		MaxDays: 11,
	}
	settings := models.SettingsUpdateRequest{
		Meta: &meata,
		Logs: &logs,
	}

	resp, err := examples.Client.UpdateSettings(&settings)
	if err != nil {
		panic(err)
	}

	println(resp.Logs.MaxDays)
}
