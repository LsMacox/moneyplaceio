package client

import (
	st "analytic-tools/moneyplace/structures"
	"fmt"
)

func GetCategories(client *Client, params st.CategoriesParams) ([]st.CategoryResponse, error) {
	resp, err := client.PerformRequest("GET", "/statistic/category", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var categories []st.CategoryResponse
	if err := client.DecodeResponse(resp, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryDetail(client *Client, categoryID int) (*st.CategoryDetailResponse, error) {
	if categoryID == 0 {
		return nil, fmt.Errorf("category ID is required")
	}

	path := fmt.Sprintf("/v1/category/%d", categoryID)
	resp, err := client.PerformRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var category st.CategoryDetailResponse
	if err := client.DecodeResponse(resp, &category); err != nil {
		return nil, err
	}

	return &category, nil
}

func GetCategoryStatistics(client *Client, categoryID int, params st.CategoryStatisticsParams) ([]st.CategoryStatisticsResponse, error) {
	if categoryID == 0 {
		return nil, fmt.Errorf("category ID is required")
	}

	path := fmt.Sprintf("/statistic/category/charts/%d", categoryID)
	resp, err := client.PerformRequest("GET", path, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statistics []st.CategoryStatisticsResponse
	if err := client.DecodeResponse(resp, &statistics); err != nil {
		return nil, err
	}

	return statistics, nil
}

func GetCategoryProducts(client *Client, categoryID int, params st.CategoryProductsParams) ([]st.CategoryProductResponse, error) {
	if categoryID == 0 {
		return nil, fmt.Errorf("category ID is required")
	}

	path := fmt.Sprintf("/statistic/category/%d", categoryID)
	resp, err := client.PerformRequest("GET", path, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var products []st.CategoryProductResponse
	if err := client.DecodeResponse(resp, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func GetCategoryBrands(client *Client, params st.CategoryBrandsParams) ([]st.CategoryBrandResponse, error) {
	if params.CategoryIDMP == 0 {
		return nil, fmt.Errorf("category ID MP is required")
	}

	resp, err := client.PerformRequest("GET", "/statistic/seller", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var brands []st.CategoryBrandResponse
	if err := client.DecodeResponse(resp, &brands); err != nil {
		return nil, err
	}

	return brands, nil
}

func GetCategorySellers(client *Client, params st.CategorySellersParams) ([]st.CategorySellerResponse, error) {
	if params.CategoryIDMP == 0 {
		return nil, fmt.Errorf("category ID MP is required")
	}

	resp, err := client.PerformRequest("GET", "/statistic/seller", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sellers []st.CategorySellerResponse
	if err := client.DecodeResponse(resp, &sellers); err != nil {
		return nil, err
	}

	return sellers, nil
}
