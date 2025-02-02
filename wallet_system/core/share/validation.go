package share

import (
	"go-wallet-system/pkg/utils"
	"strconv"
)

func UsernameValidation(username string) bool {
	return utils.Validate(username, `^[a-zA-Z0-9_]{3,20}$`)
}

func AmountValidationAndConversation(amount string) *float64 {
	vali := utils.Validate(amount, `^(0*[1-9]\d*(\.\d{1,6})?|0*\.\d{1,6})$`)
	if !vali {
		return nil
	}
	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil
	}

	return &f
}
