package domain

type Captcha struct {
	CaptchaID     string `json:"captcha_id"`
	CaptchaData   string `json:"captcha_data"`
	CaptchaLength int    `json:"captcha_length"`
	CaptchaAnswer string `json:"-"`
}
