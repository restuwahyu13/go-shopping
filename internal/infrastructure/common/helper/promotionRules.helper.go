package helper

import (
	"fmt"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	"slices"
	"strings"
)

func promotionRulesCompare(src *hopt.PromotionRules) string {
	xBUY_ONE_GET_ONE := []string{cons.BEQF, cons.BGTF, cons.BGTEF, cons.BLTF, cons.BLTEF}
	xBUY_GET_OTHER_PRODUCT := []string{cons.BEQFP, cons.BGTFP, cons.BGTEFP, cons.BLTFP, cons.BLTEFP}
	xBUY_GET_DISCOUNT_PERCENTAGE := []string{cons.BEQDPC, cons.BGTDPC, cons.BGTEDPC, cons.BLTDPC, cons.BLTEDPC}
	xBUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT := []string{cons.BEQDPCP, cons.BGTDPCP, cons.BGTEDPCP, cons.BLTDPCP, cons.BLTEDPCP}
	xBUY_GET_DISCOUNT_FIXED := []string{cons.BEQDSUFX, cons.BGTDSUFX, cons.BGTEDSUFX, cons.BLTDSUFX, cons.BLTEDSUFX}
	xBUY_GET_DISCOUNT_FIXED_WITH_PRODUCT := []string{cons.BEQDSUFXP, cons.BGTDSUFXP, cons.BGTEDSUFXP, cons.BLTDSUFXP, cons.BLTEDSUFXP}
	xBUY_GET_HALF_PRICE := []string{cons.BEQPR, cons.BGTPR, cons.BGTEPR, cons.BLTPR, cons.BLTEPR}

	if slices.Contains(xBUY_ONE_GET_ONE, src.Key) {
		return cons.BUY_ONE_GET_ONE
	} else if slices.Contains(xBUY_GET_OTHER_PRODUCT, src.Key) {
		return cons.BUY_GET_OTHER_PRODUCT
	} else if slices.Contains(xBUY_GET_DISCOUNT_PERCENTAGE, src.Key) {
		return cons.BUY_GET_DISCOUNT_PERCENTAGE
	} else if slices.Contains(xBUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT, src.Key) {
		return cons.BUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT
	} else if slices.Contains(xBUY_GET_DISCOUNT_FIXED, src.Key) {
		return cons.BUY_GET_DISCOUNT_FIXED
	} else if slices.Contains(xBUY_GET_DISCOUNT_FIXED_WITH_PRODUCT, src.Key) {
		return cons.BUY_GET_DISCOUNT_FIXED_WITH_PRODUCT
	} else if slices.Contains(xBUY_GET_HALF_PRICE, src.Key) {
		return cons.BUY_GET_HALF_PRICE
	}

	return cons.EMPTY
}

func PromotionDiscount(total float64, discount int64, action string) (t, d float64) {
	parseFloat, _ := NewParser().ToFloat(discount)

	if action == cons.PERCENTAGE {
		d = total * (parseFloat / 100)
		t = total - d
		return
	} else if action == cons.FIXED {
		d = parseFloat
		t = total - parseFloat
		return
	}

	return
}

func PromotionRules(src []hdto.PromotionRulesDTO) *hopt.PromotionRules {
	var (
		value string               = ""
		res   *hopt.PromotionRules = new(hopt.PromotionRules)
	)

	for _, v := range src {
		res.Key += v.Key
		value += fmt.Sprintf("%s ", v.Value)
	}

	res.Key = strings.TrimSpace(res.Key)
	res.Type = promotionRulesCompare(res)

	result := strings.Split(strings.TrimSpace(value), " ")

	if len(result) == 3 && cons.PromotionRulesMapping[res.Key] {
		res.Action = result[0]
		res.Operator = result[1]
		res.OptionOne = result[2]

		return res
	}

	if len(result) == 4 && cons.PromotionRulesMapping[res.Key] {
		res.Action = result[0]
		res.Operator = result[1]
		res.OptionOne = result[2]
		res.OptionTwo = result[3]

		return res
	}

	if len(result) == 5 && cons.PromotionRulesMapping[res.Key] {
		res.Action = result[0]
		res.Operator = result[1]
		res.OptionOne = result[2]
		res.OptionTwo = result[3]
		res.OptionThree = result[4]

		return res
	}

	if len(result) == 6 && cons.PromotionRulesMapping[res.Key] {
		res.Action = result[0]
		res.Operator = result[1]
		res.OptionOne = result[2]
		res.OptionTwo = result[3]
		res.OptionThree = result[4]
		res.OptionFour = result[5]

		return res
	}

	return nil
}
