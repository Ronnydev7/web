package lib

import (
	"api/config"
	"api/intl"
	"api/intl/intlgenerated"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type (
	MailIdentity struct {
		Name    string
		Address string
	}

	TemplateMail struct {
		Tos          []MailIdentity
		TemplateId   string
		TemplateData map[string]interface{}
	}

	SendgridMailer struct {
		Mailer
		mailerConfig   config.MailerConfig
		sendgridConfig config.SendgridConfig
	}

	SendgridMailerError struct {
		intl.IntlError
		inner error
	}

	Mailer interface {
		SendTemplateMail(TemplateMail) intl.IntlError
		SendConfirmSignupEmailEmail(receiver string, emailSignupUrl string) intl.IntlError
		SendResetPasswordEmail(receiver string, resetEmailUrl string) intl.IntlError
	}

	NewMailerFunc func(config.MailerConfig) Mailer
)

const MAILER_LOGGER_NAME = "lib.mailer"

var NewMailer NewMailerFunc = func(mailerConfig config.MailerConfig) Mailer {
	return SendgridMailer{
		mailerConfig:   mailerConfig,
		sendgridConfig: config.GetSendgridConfig(),
	}
}

func (mailer SendgridMailer) handleErrorFromSendgrid(err error) intl.IntlError {
	NewLogger(MAILER_LOGGER_NAME).LogError(err)
	return &SendgridMailerError{
		inner: err,
	}
}

func (mailer SendgridMailer) SendConfirmSignupEmailEmail(receiver string, confirmSignupUrl string) intl.IntlError {
	m := TemplateMail{
		Tos:        []MailIdentity{{Address: receiver}},
		TemplateId: mailer.sendgridConfig.GetEmailSignupConfirmationTemplateId(),
		TemplateData: map[string]interface{}{
			"confirm_email_signup_address": confirmSignupUrl,
		},
	}

	return mailer.SendTemplateMail(m)
}

func (mailer SendgridMailer) SendResetPasswordEmail(receiver string, resetPasswordUrl string) intl.IntlError {
	m := TemplateMail{
		Tos:        []MailIdentity{{Address: receiver}},
		TemplateId: mailer.sendgridConfig.GetResetPasswordTemplateId(),
		TemplateData: map[string]interface{}{
			"reset_password_url": resetPasswordUrl,
		},
	}

	return mailer.SendTemplateMail(m)
}

func (mailer SendgridMailer) SendTemplateMail(m TemplateMail) intl.IntlError {
	p := mail.NewPersonalization()
	emails := make([]*mail.Email, len(m.Tos))
	for i, identity := range m.Tos {
		emails[i] = mail.NewEmail(
			identity.Name,
			identity.Address,
		)
	}
	p.AddTos(emails...)
	for field, value := range m.TemplateData {
		p.SetDynamicTemplateData(field, value)
	}

	sendgridMail := mail.NewV3Mail()
	sendgridMail.SetFrom(
		mail.NewEmail(mailer.mailerConfig.GetNoResponseEmailName(), mailer.mailerConfig.GetNoResponseEmailAddress()),
	)
	sendgridMail.SetTemplateID(m.TemplateId)
	sendgridMail.AddPersonalizations(p)

	_, err := mailer.sendSendgridMail(sendgridMail)
	return err
}

func (mailer SendgridMailer) sendSendgridMail(sendgridMail *mail.SGMailV3) (*rest.Response, intl.IntlError) {
	apiKey := mailer.sendgridConfig.GetApiKey()
	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(sendgridMail)
	response, err := sendgrid.API(request)
	if err != nil {
		return nil, mailer.handleErrorFromSendgrid(err)
	}
	return response, nil
}

func (err SendgridMailerError) Error() string {
	return err.inner.Error()
}

func (SendgridMailerError) GetIntlKey() string {
	return intlgenerated.COMMON_STRINGS__UNKNOWN_SERVER_ERROR
}
