// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// {  "B": "3",  "EQ": '=', "F":  "1" }

var PromotionRulesMapping = map[string]bool{
	// Contoh: Beli 2 product gratis 1
	"BEQF":  true,
	"BGTF":  true,
	"BGTEF": true,
	"BLTF":  true,
	"BLTEF": true,

	//  Contoh: Beli 2 product gratis product lainnya
	"BEQFP":  true,
	"BGTFP":  true,
	"BGTEFP": true,
	"BLTFP":  true,
	"BLTEFP": true,

	//  Contoh: Beli 2 product dapat diskon 10%
	"BEQDPC":  true,
	"BGTDPC":  true,
	"BGTEDPC": true,
	"BLTDPC":  true,
	"BLTEDPC": true,

	"BEQDPCP":  true,
	"BGTDPCP":  true,
	"BGTEDPCP": true,
	"BLTDPCP":  true,
	"BLTEDPCP": true,

	// Contoh: Beli 2 product dapat diskon potongan Rp 100.000
	"BEQDSUFX":  true,
	"BGTDSUFX":  true,
	"BGTEDSUFX": true,
	"BLTDSUFX":  true,
	"BLTEDSUFX": true,

	"BEQDSUFXP":  true,
	"BGTDSUFXP":  true,
	"BGTEDSUFXP": true,
	"BLTDSUFXP":  true,
	"BLTEDSUFXP": true,

	// Contoh: Beli 3 produk bayar hanya 2 harga
	"BEQPR":  true,
	"BGTPR":  true,
	"BGTEPR": true,
	"BLTPR":  true,
	"BLTEPR": true,
}

const (
	BEQF  = "BEQF"
	BGTF  = "BGTF"
	BGTEF = "BGTEF"
	BLTF  = "BLTF"
	BLTEF = "BLTEF"

	BEQFP = "BEQFP"
	BGTFP = "BGTFP"
	BGTEP = "BGTEP"
	BLTP  = "BLTP"
	BLTEP = "BLTEP"

	BEQDPC  = "BEQDPC"
	BGTDPC  = "BGTDPC"
	BGTEDPC = "BGTEDPC"
	BLTDPC  = "BLTDPC"
	BLTEDPC = "BLTEDPC"

	BEQDPCP  = "BEQDPCP"
	BGTDPCP  = "BGTDPCP"
	BGTEDPCP = "BGTEDPCP"
	BLTDPCP  = "BLTDPCP"
	BLTEDPCP = "BLTEDPCP"

	BEQDSUFX  = "BEQDSUFX"
	BGTDSUFX  = "BGTDSUFX"
	BGTEDSUFX = "BGTEDSUFX"
	BLTDSUFX  = "BLTDSUFX"
	BLTEDSUFX = "BLTEDSUFX"

	BEQDSUFXP  = "BEQDSUFXP"
	BGTDSUFXP  = "BGTDSUFXP"
	BGTEDSUFXP = "BGTEDSUFXP"
	BLTDSUFXP  = "BLTDSUFXP"
	BLTEDSUFXP = "BLTEDSUFXP"

	BEQPR  = "BEQPR"
	BGTPR  = "BGTPR"
	BGTEPR = "BGTEPR"
	BLTPR  = "BLTPR"
	BLTEPR = "BLTEPR"
)

const (
	BUY_ONE_GET_ONE                          = "BUY_ONE_GET_ONE"
	BUY_GET_OTHER_PRODUCT                    = "BUY_GET_OTHER_PRODUCT"
	BUY_GET_DISCOUNT_PERCENTAGE              = "BUY_GET_DISCOUNT_PERCENTAGE"
	BUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT = "BUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT"
	BUY_GET_DISCOUNT_FIXED                   = "BUY_GET_DISCOUNT_FIXED"
	BUY_GET_DISCOUNT_FIXED_WITH_PRODUCT      = "BUY_GET_DISCOUNT_FIXED_WITH_PRODUCT"
	BUY_GET_HALF_PRICE                       = "BUY_GET_HALF_PRICE"
)

const (
	// Produk sejenis (kategori sama), type berbeda, brand harus sama
	SAME_BRAND_SIMILAR_CATEGORY = "SAME_BRAND_SIMILAR_CATEGORY"

	// Produk berbeda boleh apa saja, brand harus sama
	SAME_BRAND_ANY_CATEGORY = "SAME_BRAND_ANY_CATEGORY"

	// Produk dan brand boleh berbeda
	ANY_BRAND_ANY_CATEGORY = "ANY_BRAND_ANY_CATEGORY"

	// Produk harus sama persis, brand boleh berbeda
	ANY_BRAND_SAME_PRODUCT = "ANY_BRAND_SAME_PRODUCT"
)

type PromotionRules struct {
	Key         string
	Action      string
	Type        string
	Operator    string
	OptionOne   string
	OptionTwo   string
	OptionThree string
	OptionFour  string
}

type PromotionRulesDTO struct {
	Key   string
	Value string
}

func promotionRulesCompare(src *PromotionRules) string {
	xBUY_ONE_GET_ONE := []string{BEQF, BGTF, BGTEF, BLTF, BLTEF}
	xBUY_GET_OTHER_PRODUCT := []string{BEQFP, BGTFP, BGTEP, BLTP, BLTEP}
	xBUY_GET_DISCOUNT_PERCENTAGE := []string{BEQDPC, BGTDPC, BGTEDPC, BLTDPC, BLTEDPC}
	xBUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT := []string{BEQDPCP, BGTDPCP, BGTEDPCP, BLTDPCP, BLTEDPCP}
	xBUY_GET_DISCOUNT_FIXED := []string{BEQDSUFX, BGTDSUFX, BGTEDSUFX, BLTDSUFX, BLTEDSUFX}
	xBUY_GET_DISCOUNT_FIXED_WITH_PRODUCT := []string{BEQDSUFXP, BGTDSUFXP, BGTEDSUFXP, BLTDSUFXP, BLTEDSUFXP}
	xBUY_GET_HALF_PRICE := []string{BEQPR, BGTPR, BGTEPR, BLTPR, BLTEPR}

	if slices.Contains(xBUY_ONE_GET_ONE, src.Key) {
		return BUY_ONE_GET_ONE
	} else if slices.Contains(xBUY_GET_OTHER_PRODUCT, src.Key) {
		return BUY_GET_OTHER_PRODUCT
	} else if slices.Contains(xBUY_GET_DISCOUNT_PERCENTAGE, src.Key) {
		return BUY_GET_DISCOUNT_PERCENTAGE
	} else if slices.Contains(xBUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT, src.Key) {
		return BUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT
	} else if slices.Contains(xBUY_GET_DISCOUNT_FIXED, src.Key) {
		return BUY_GET_DISCOUNT_FIXED
	} else if slices.Contains(xBUY_GET_DISCOUNT_FIXED_WITH_PRODUCT, src.Key) {
		return BUY_GET_DISCOUNT_FIXED_WITH_PRODUCT
	} else if slices.Contains(xBUY_GET_HALF_PRICE, src.Key) {
		return BUY_GET_HALF_PRICE
	}

	return ""
}

func PromoRules(src []PromotionRulesDTO) *PromotionRules {
	var (
		value string          = ""
		res   *PromotionRules = new(PromotionRules)
	)

	for _, v := range src {
		res.Key += v.Key
		value += fmt.Sprintf("%s ", v.Value)
	}

	res.Key = strings.TrimSpace(res.Key)
	res.Type = promotionRulesCompare(res)

	result := strings.Split(strings.TrimSpace(value), " ")
	fmt.Println("RESULT", result)

	if len(result) == 3 && PromotionRulesMapping[res.Key] {
		res.Action = result[0]
		res.Operator = result[1]
		res.OptionOne = result[2]

		return res
	}

	if len(result) == 4 && PromotionRulesMapping[res.Key] {
		res.Action = result[0]
		res.Operator = result[1]
		res.OptionOne = result[2]
		res.OptionTwo = result[3]

		return res
	}

	if len(result) == 5 && PromotionRulesMapping[res.Key] {
		res.Action = result[0]
		res.Operator = result[1]
		res.OptionOne = result[2]
		res.OptionTwo = result[3]
		res.OptionThree = result[4]

		return res
	}

	if len(result) == 6 && PromotionRulesMapping[res.Key] {
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

func DecimalToFloat(n int64) float64 {
	return float64(n) / 100
}

func promotionDiscount(total int, discount int64, action string) (t, d float64) {
	parseFloat, _ := strconv.ParseFloat(strconv.Itoa(int(discount)), 64)

	if action == "percentage" {
		d = float64(total) * (parseFloat / 100)
		t = float64(total) - d
		return
	} else if action == "fixed" {
		d = parseFloat
		t = float64(total) - parseFloat
		return
	}

	return
}

// func main() {
//     total := 100000
//     diskonPersen := 10.0

//     diskon, totalBayar := hitungDiskon(total, diskonPersen)

//     fmt.Printf("Total Belanja: Rp %d\n", total)
//     fmt.Printf("Diskon %.0f%%: Rp %.0f\n", diskonPersen, diskon)
//     fmt.Printf("Total Bayar: Rp %.0f\n", totalBayar)
// }

func main() {
	// BUY_ONE_GET_ONE := []PromotionRulesDTO{
	// 	{Key: "B", Value: "3"},
	// 	{Key: "EQ", Value: "="},
	// 	{Key: "F", Value: "1"},
	// }

	BUY_GET_OTHER_PRODUCT := []PromotionRulesDTO{
		{Key: "B", Value: "3"},
		{Key: "EQ", Value: "="},
		{Key: "F", Value: "7e672597-5ea9-4b75-978c-1787630a12b0"},
		{Key: "P", Value: "PRODUCT_MATCH_SAME_BRAND"},
	}

	// BUY_GET_DISCOUNT_PERCENTAGE := []PromotionRulesDTO{
	// 	{Key: "B", Value: "3"},
	// 	{Key: "GT", Value: ">"},
	// 	{Key: "D", Value: "10"},
	// 	{Key: "PC", Value: "%"},
	// }

	// BUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT := []PromotionRulesDTO{
	// 	{Key: "B", Value: "3"},
	// 	{Key: "GT", Value: ">"},
	// 	{Key: "D", Value: "10"},
	// 	{Key: "PC", Value: "%"},
	// 	{Key: "P", Value: "SAME_BRAND_SIMILAR_CATEGORY"},
	// }

	// BUY_GET_DISCOUNT_FIXED := []PromotionRulesDTO{
	// 	{Key: "B", Value: "3"},
	// 	{Key: "GT", Value: ">"},
	// 	{Key: "D", Value: "10"},
	// 	{Key: "SU", Value: "-"},
	// 	{Key: "FX", Value: "fx"},
	// }

	// BUY_GET_DISCOUNT_FIXED_WITH_PRODUCT := []PromotionRulesDTO{
	// 	{Key: "B", Value: "3"},
	// 	{Key: "GT", Value: ">"},
	// 	{Key: "D", Value: "10"},
	// 	{Key: "SU", Value: "-"},
	// 	{Key: "FX", Value: "fx"},
	// 	{Key: "P", Value: SAME_BRAND_SIMILAR_CATEGORY},
	// }

	// BUY_GET_HALF_PRICE := []PromotionRulesDTO{
	// 	{Key: "B", Value: "3"},
	// 	{Key: "GT", Value: ">"},
	// 	{Key: "PR", Value: "2"},
	// }

	result := PromoRules(BUY_GET_OTHER_PRODUCT)
	if result == nil {
		fmt.Println("Tidak valid")
	}

	fmt.Println("KEY: ", result.Key)
	fmt.Println("OPERATOR: ", result.Operator)
	fmt.Println("ACTION: ", result.Action)
	fmt.Println("OPTIONS 1: ", result.OptionOne)
	fmt.Println("OPTIONS 2: ", result.OptionTwo)
	fmt.Println("OPTIONS 3: ", result.OptionThree)
	fmt.Println("OPTIONS 4: ", result.OptionFour)
	fmt.Println("\n")
	fmt.Println("MASUK TYPE: ", result.Type)

	// totalPrice, discount := promotionDiscount(100000, 50, "percentage")

	// fmt.Println(totalPrice)
	// fmt.Println(discount)
}
