package repository

import (
	"github.com/mojocn/base64Captcha"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type CaptchaRepository struct{}

func NewCaptchaRepository() irepository.ICaptchaRepository {
	return &CaptchaRepository{}
}

func (*CaptchaRepository) NewDigitCaptcha() (domain.Captcha, error) {
	driver := base64Captcha.NewDriverDigit(75, 280, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	id, b64s, _, generateErr := cp.Generate()
	if generateErr != nil {
		return domain.Captcha{}, generateErr
	}
	return domain.Captcha{
			CaptchaID:     id,
			CaptchaData:   b64s,
			CaptchaLength: 6},
		nil
}

func (*CaptchaRepository) Verify(captcha domain.Captcha) bool {
	return base64Captcha.DefaultMemStore.Verify(captcha.CaptchaID, captcha.CaptchaAnswer, true)
}
