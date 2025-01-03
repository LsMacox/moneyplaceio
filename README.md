

# How to use


```go
c = client.NewClient("https://api.moneyplace.io", "e23ce887f5759136eb8983eee02a9")
```

### Products

```go

/* Список товаров */
params := st.ProductsParams{
    MP:     "wildberries",
    Type:   "fbo",
    Period: "week",
    Sort:   "-turnover",
}

products, err := GetProducts(client, params)

/* Информация о товаре */
params := st.ProductDetailParams{
    Expand: "category,seller,brand",
}

productDetail, err := GetProductDetail(client, 6810650, params)


/* Статистика продаж товара */
// Example 1: Weekly statistics
params := st.ProductStatisticsParams{
    Type:   "fbo",
    Period: "week",
}

// Example 2: Custom date range statistics
params := st.ProductStatisticsParams{
    Type:        "fbo",
    Period:      "period",
    DateBetween: "1620345600,1621468800",
}

statistics, err := GetProductStatistics(client, 6810650, params)

/* Позиция товара в категориях */
// Example 1: Weekly statistics
params := st.ProductPositionParams{
    Period: "week",
}

// Example 2: Custom date range statistics
params := st.ProductPositionParams{
    Period:      "period",
    DateBetween: "1620345600,1621468800",
}

positions, err := GetProductPosition(client, 6810650, params)
```

### Categories
```go
/* Получение дочерних категорий/всех категорий */
params := st.CategoriesParams{
    ID:      3273, // Без id получим всех
    MP:      "wildberries",
    Type:    "fbo",
    Period:  "week",
    Sort:    "-turnover",
}

childCategories, err := GetCategories(client, params)


/* Информация о категории */
categoryDetail, err := GetCategoryDetail(client, 3273)

/* Статистика продаж категории */
// Example 1: Weekly statistics
params := st.CategoryStatisticsParams{
    Type:   "fbo",
    Period: "week",
}

// Example 2: Custom date range statistics
params := st.CategoryStatisticsParams{
    Type:        "fbo",
    Period:      "period",
    DateBetween: "1620345600,1621468800",
}

statistics, err := GetCategoryStatistics(client, 3273, params)

/* Статистика продаж товаров в категории */
// Example 1: Weekly statistics with price filter
params := st.CategoryProductsParams{
    Type:         "fbo",
    Period:       "week",
    Sort:         "-turnover",
    PriceBetween: "107,25480",
}

// Example 2: Custom date range with pagination
params := st.CategoryProductsParams{
    Type:         "fbo",
    Period:       "period",
    DateBetween:  "1620345600,1621468800",
    Sort:         "-turnover",
    PriceBetween: "107,25480",
    PerPage:      500,
    Page:         2,
}

products, err := GetCategoryProducts(client, 3273, params)

/* Статистика брендов в категории */
params := st.CategoryBrandsParams{
    MP:           "wildberries",
    Type:         "fbo",
    Period:       "period",
    Sort:         "-turnover",
    DateBetween:  "1620345600,1621468800",
    PriceBetween: "107,25480",
    CategoryIDMP: 3273,
}

brands, err := GetCategoryBrands(client, params)

/* Статистика продавцов в категории */
params := st.CategorySellersParams{
    MP:           "wildberries",
    Type:         "fbo",
    Period:       "period",
    Sort:         "-turnover",
    DateBetween:  "1620345600,1621468800",
    PriceBetween: "94,25480",
    CategoryIDMP: 3273,
}

sellers, err := GetCategorySellers(client, params)
``` 