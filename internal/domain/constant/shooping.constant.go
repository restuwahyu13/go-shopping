package cons

const (
	CHECKOUT = "checkout"
	REMOVE   = "remove"
	ORDER    = "order"
)

const (
	BEQF  = "BEQF"
	BGTF  = "BGTF"
	BGTEF = "BGTEF"
	BLTF  = "BLTF"
	BLTEF = "BLTEF"

	BEQFP  = "BEQFP"
	BGTFP  = "BGTFP"
	BGTEFP = "BGTEFP"
	BLTFP  = "BLTFP"
	BLTEFP = "BLTEFP"

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
	// Produk sejenis (kategori sama), type berbeda, brand harus sama
	SAME_BRAND_SIMILAR_CATEGORY = "SAME_BRAND_SIMILAR_CATEGORY"

	// Produk berbeda boleh apa saja, brand harus sama
	SAME_BRAND_ANY_CATEGORY = "SAME_BRAND_ANY_CATEGORY"

	// Produk dan brand boleh berbeda
	ANY_BRAND_ANY_CATEGORY = "ANY_BRAND_ANY_CATEGORY"

	// Produk harus sama persis, brand boleh berbeda
	ANY_BRAND_SAME_PRODUCT = "ANY_BRAND_SAME_PRODUCT"
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

var PromotionRulesMapping = map[string]bool{
	// Contoh: Beli 2 product gratis 1
	"BEQF":  true,
	"BGTF":  true,
	"BGTEF": true,
	"BLTF":  true,
	"BLTEF": true,

	//  Contoh: Beli 2 product gratis product lainnya
	"BEQFP": true,
	"BGTFP": true,
	"BGTEP": true,
	"BLTP":  true,
	"BLTEP": true,

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
