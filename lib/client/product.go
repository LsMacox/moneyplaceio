package client

import (
	"fmt"

	st "github.com/LsMacox/moneyplaceio/structures"
)

func GetProducts(client *Client, params st.ProductsParams) ([]st.ProductResponse, error) {
	if params.MP == "" {
		return nil, fmt.Errorf("marketplace parameter is required")
	}
	if params.Type == "" {
		return nil, fmt.Errorf("type parameter is required")
	}
	if params.Period == "" {
		return nil, fmt.Errorf("period parameter is required")
	}

	resp, err := client.PerformRequest("GET", "/statistic/product", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var products []st.ProductResponse
	if err := client.DecodeResponse(resp, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductDetail(client *Client, productID int, params st.ProductDetailParams) (*st.ProductDetailResponse, error) {
	if productID == 0 {
		return nil, fmt.Errorf("product ID is required")
	}

	path := fmt.Sprintf("/v1/product/%d", productID)
	resp, err := client.PerformRequest("GET", path, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var product st.ProductDetailResponse
	if err := client.DecodeResponse(resp, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func GetProductStatistics(client *Client, productID int, params st.ProductStatisticsParams) ([]st.ProductStatisticsResponse, error) {
	if productID == 0 {
		return nil, fmt.Errorf("product ID is required")
	}
	if params.Type == "" {
		return nil, fmt.Errorf("type parameter is required")
	}
	if params.Period == "" {
		return nil, fmt.Errorf("period parameter is required")
	}

	path := fmt.Sprintf("/statistic/product/charts/%d", productID)
	resp, err := client.PerformRequest("GET", path, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statistics []st.ProductStatisticsResponse
	if err := client.DecodeResponse(resp, &statistics); err != nil {
		return nil, err
	}

	return statistics, nil
}

func GetProductPosition(client *Client, productID int, params st.ProductPositionParams) ([]st.ProductPositionResponse, error) {
	if productID == 0 {
		return nil, fmt.Errorf("product ID is required")
	}
	if params.Period == "" {
		return nil, fmt.Errorf("period parameter is required")
	}

	path := fmt.Sprintf("/statistic/product-info/charts/%d", productID)
	resp, err := client.PerformRequest("GET", path, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var positions []st.ProductPositionResponse
	if err := client.DecodeResponse(resp, &positions); err != nil {
		return nil, err
	}

	return positions, nil
}
