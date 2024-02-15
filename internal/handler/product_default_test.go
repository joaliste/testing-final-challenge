package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductsDefault_Get(t *testing.T) {

	t.Run("success search products - no query - full", func(t *testing.T) {
		// arrange
		rp := repository.NewProductsMock()
		rp.SearchProductsFunc = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			p = map[int]internal.Product{
				1: {
					Id: 1,
					ProductAttributes: internal.ProductAttributes{
						Description: "item_1",
						Price:       1,
						SellerId:    1,
					},
				},
				2: {
					Id: 2,
					ProductAttributes: internal.ProductAttributes{
						Description: "item_2",
						Price:       2,
						SellerId:    2,
					},
				},
			}
			return p, nil
		}
		hd := handler.NewProductsDefault(rp)
		hdFunc := hd.Get()
		req := httptest.NewRequest(http.MethodGet, "/product", nil)
		res := httptest.NewRecorder()

		expectedBodyOutput := `
					{"message": "success", "data": {
						"1": {"id": 1, "description": "item_1", "price": 1, "seller_id": 1},
						"2": {"id": 2, "description": "item_2", "price": 2, "seller_id": 2}
					}}
				`
		expectedHeaderOutput := http.Header{
			"Content-Type": []string{"application/json"},
		}
		// act
		hdFunc(res, req)
		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedBodyOutput, res.Body.String())
		require.Equal(t, expectedHeaderOutput, res.Header())
	})

	t.Run("success search products - query one", func(t *testing.T) {
		// arrange
		rp := repository.NewProductsMock()
		rp.SearchProductsFunc = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			p = map[int]internal.Product{
				1: {
					Id: 1,
					ProductAttributes: internal.ProductAttributes{
						Description: "item_1",
						Price:       1,
						SellerId:    1,
					},
				},
			}
			return p, nil
		}
		hd := handler.NewProductsDefault(rp)
		hdFunc := hd.Get()
		req := httptest.NewRequest(http.MethodGet, "/product", nil)
		res := httptest.NewRecorder()
		q := req.URL.Query()
		q.Add("id", "1")
		req.URL.RawQuery = q.Encode()

		expectedBodyOutput := `
					{"message": "success", "data": {
						"1": {"id": 1, "description": "item_1", "price": 1, "seller_id": 1}
					}}
				`
		expectedHeaderOutput := http.Header{
			"Content-Type": []string{"application/json"},
		}
		// act
		hdFunc(res, req)
		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedBodyOutput, res.Body.String())
		require.Equal(t, expectedHeaderOutput, res.Header())
	})

}
