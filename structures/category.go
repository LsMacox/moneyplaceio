package structures

type CategoriesParams struct {
	MP          string `url:"mp,omitempty"`
	Type        string `url:"type"`
	Period      string `url:"period"`
	Sort        string `url:"sort,omitempty"`
	Name        string `url:"q[name][like],omitempty"`
	ID          int    `url:"id,omitempty"`
	DateBetween string `url:"q[date][between],omitempty"`
}

type CategoryResponse struct {
	AverageBill   float64  `json:"average_bill"`
	BrandsCount   int      `json:"brands_count"`
	ProductsCount int      `json:"products_count"`
	SellersCount  int      `json:"sellers_count"`
	TotalSales    int      `json:"total_sales"`
	Tstc          int      `json:"tstc"`
	Tsts          float64  `json:"tsts"`
	Turnover      int      `json:"turnover"`
	Category      Category `json:"category"`
}

type Category struct {
	Childrens     string `json:"childrens"`
	Count         int    `json:"count"`
	ID            int    `json:"id"`
	IDMP          int    `json:"id_mp"`
	IDParent      int    `json:"id_parent"`
	MP            string `json:"mp"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	ProductsCount int    `json:"products_count"`
}

type CategoryDetailResponse struct {
	Childrens     []int  `json:"childrens"`
	Count         int    `json:"count"`
	ID            int    `json:"id"`
	IDMP          int    `json:"id_mp"`
	IDParent      int    `json:"id_parent"`
	Link          string `json:"link"`
	MP            string `json:"mp"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	ProductsCount int    `json:"products_count"`
}

type CategoryStatisticsParams struct {
	Type        string `url:"type"`
	Period      string `url:"period"`
	DateBetween string `url:"q[date][between],omitempty"`
}

type CategoryStatisticsResponse struct {
	AvgPrice    float64 `json:"avg_price"`
	Date        string  `json:"date"`
	TotalAmount string  `json:"total_amount"`
	TotalSale   string  `json:"total_sale"`
	TotalSum    string  `json:"total_sum"`
}

type CategoryProductsParams struct {
	Type         string `url:"type"`
	Period       string `url:"period"`
	Sort         string `url:"sort,omitempty"`
	DateBetween  string `url:"q[date][between],omitempty"`
	PriceBetween string `url:"q[price][between],omitempty"`
	PerPage      int    `url:"per-page,omitempty"`
	Page         int    `url:"page,omitempty"`
}

type CategoryProductResponse struct {
	AvgPosition       float64  `json:"avg_position"`
	AvgPrice          float64  `json:"avg_price"`
	DaysInStock       string   `json:"days_in_stock"`
	DaysWithSales     string   `json:"days_with_sales"`
	IDBrand           int      `json:"id_brand"`
	IDCategory        int      `json:"id_category"`
	IDSeller          int      `json:"id_seller"`
	Loses             int      `json:"loses"`
	MP                string   `json:"mp"`
	Rating            float64  `json:"rating"`
	Reviews           int      `json:"reviews"`
	SKU               int      `json:"sku"`
	TotalDaysInPeriod string   `json:"total_days_in_period"`
	TotalSales        string   `json:"total_sales"`
	Turnover          string   `json:"turnover"`
	DateGraph         []string `json:"date_graph"`
	SaleGraph         []int    `json:"sale_graph"`
	Brand             Brand    `json:"brand"`
	Category          Category `json:"category"`
	Product           Product  `json:"product"`
	Seller            Seller   `json:"seller"`
}

type Brand struct {
	ID        int    `json:"id"`
	IDMP      int    `json:"id_mp"`
	MP        string `json:"mp"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MP        string `json:"mp"`
	SKU       int    `json:"sku"`
	RealPrice int    `json:"real_price"`
}

type Seller struct {
	ID   int    `json:"id"`
	IDMP int    `json:"id_mp"`
	Name string `json:"name"`
	MP   string `json:"mp"`
}

type CategoryBrandsParams struct {
	MP           string `url:"mp"`
	Type         string `url:"type"`
	Period       string `url:"period"`
	Sort         string `url:"sort,omitempty"`
	DateBetween  string `url:"q[date][between],omitempty"`
	PriceBetween string `url:"q[price][between],omitempty"`
	CategoryIDMP int    `url:"q[id_mp_category][equal]"`
}

type CategoryBrandResponse struct {
	AverageBill   float64 `json:"average_bill"`
	IDBrand       int     `json:"id_brand"`
	MP            string  `json:"mp"`
	ProductsCount string  `json:"products_count"`
	Rating        float64 `json:"rating"`
	TotalSales    string  `json:"total_sales"`
	Turnover      string  `json:"turnover"`
	Brand         Brand   `json:"brand"`
}

type CategorySellerResponse struct {
	AverageBill   float64 `json:"average_bill"`
	BrandsCount   string  `json:"brands_count"`
	IDSeller      int     `json:"id_seller"`
	MP            string  `json:"mp"`
	ProductsCount string  `json:"products_count"`
	Rating        float64 `json:"rating"`
	TotalSales    string  `json:"total_sales"`
	Turnover      string  `json:"turnover"`
	Seller        Seller  `json:"seller"`
}

type CategorySellersParams struct {
	MP           string `url:"mp"`
	Type         string `url:"type"`
	Period       string `url:"period"`
	Sort         string `url:"sort,omitempty"`
	DateBetween  string `url:"q[date][between],omitempty"`
	PriceBetween string `url:"q[price][between],omitempty"`
	CategoryIDMP int    `url:"q[id_mp_category][equal]"`
}
