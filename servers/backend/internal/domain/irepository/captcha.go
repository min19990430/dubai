package irepository

import (
	"auto-monitoring/internal/domain"
)

type ICaptchaRepository interface {
	NewDigitCaptcha() (domain.Captcha, error)
	Verify(domain.Captcha) bool
}
