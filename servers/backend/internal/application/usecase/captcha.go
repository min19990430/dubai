package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type CaptchaUsecase struct {
	captcha irepository.ICaptchaRepository
}

func NewCaptchaUsecase(captcha irepository.ICaptchaRepository) *CaptchaUsecase {
	return &CaptchaUsecase{
		captcha: captcha,
	}
}

func (cu *CaptchaUsecase) NewDigit() (domain.Captcha, error) {
	return cu.captcha.NewDigitCaptcha()
}

func (cu *CaptchaUsecase) Verify(captcha domain.Captcha) bool {
	return cu.captcha.Verify(captcha)
}
