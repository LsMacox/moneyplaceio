package structures

type ProductsParams struct {
	MP     string `url:"mp"`
	Type   string `url:"type"`
	Period string `url:"period"`
	Sort   string `url:"sort,omitempty"`
	Name   string `url:"q[name][like],omitempty"`
}

type ProductResponse struct {
	Sales      string        `json:"Sales"`
	IDBrand    int           `json:"id_brand"`
	IDCategory int           `json:"id_category"`
	IDSeller   int           `json:"id_seller"`
	IDType     int           `json:"id_type"`
	MP         string        `json:"mp"`
	SKU        int           `json:"sku"`
	Turnover   string        `json:"turnover"`
	Product    ProductDetail `json:"product"`
	Seller     Seller        `json:"seller"`
	Brand      Brand         `json:"brand"`
	Category   Category      `json:"category"`
}

type ProductDetail struct {
	Amount            int    `json:"amount"`
	CommentsCount     int    `json:"comments_count"`
	CreatedAt         int64  `json:"created_at"`
	DeliverySchema    int    `json:"delivery_schema"`
	Discount          int    `json:"discount"`
	ID                int    `json:"id"`
	IDBrand           int    `json:"id_brand"`
	IDCategory        int    `json:"id_category"`
	IDRoot            int    `json:"id_root"`
	IDSeller          int    `json:"id_seller"`
	IDType            int    `json:"id_type"`
	Image             string `json:"image"`
	Link              string `json:"link"`
	MP                string `json:"mp"`
	Name              string `json:"name"`
	Position          int    `json:"position"`
	PriceWithDiscount int    `json:"price_with_discount"`
	Rate              string `json:"rate"`
	RealPrice         int    `json:"real_price"`
	SKU               int    `json:"sku"`
	StoreAmount       int    `json:"store_amount"`
	UpdatedAt         int64  `json:"updated_at"`
}

type ProductDetailParams struct {
	Expand string `url:"expand,omitempty"` // comma-separated list: category,seller,brand
}

type ProductDetailResponse struct {
	ProductDetail
	Brand    *Brand    `json:"brand,omitempty"`
	Seller   *Seller   `json:"seller,omitempty"`
	Category *Category `json:"category,omitempty"`
}

type ProductStatisticsParams struct {
	Type        string `url:"type"`
	Period      string `url:"period"`
	DateBetween string `url:"q[date][between],omitempty"`
}

type ProductStatisticsResponse struct {
	AvgPrice      float64 `json:"avg_price"`
	CommentsCount string  `json:"comments_count"`
	CommentsToday string  `json:"comments_today"`
	Date          string  `json:"date"`
	Discount      int     `json:"discount"`
	Position      int     `json:"position"`
	Price         int     `json:"price"`
	Rating        int     `json:"rating"`
	TotalAmount   string  `json:"total_amount"`
	TotalSale     string  `json:"total_sale"`
	TotalSum      string  `json:"total_sum"`
}

type ProductPositionParams struct {
	Period      string `url:"period"`
	DateBetween string `url:"q[date][between],omitempty"`
}

type ProductPositionResponse struct {
	Category        Category `json:"category"`
	DateGraph       []string `json:"date_graph"`
	DifferenceGraph []int    `json:"difference_graph"`
	IDCategory      int      `json:"id_category"`
	PositionGraph   []int    `json:"position_graph"`
	VirtualMP       string   `json:"virtual_mp"`
}
