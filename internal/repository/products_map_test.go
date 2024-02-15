package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductsMap_SearchProducts(t *testing.T) {
	t.Run("product found", func(t *testing.T) {
		// arrange
		product := internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "test_1",
				Price:       200.0,
				SellerId:    1,
			},
		}
		db := map[int]internal.Product{
			1: product,
		}
		rp := repository.NewProductsMap(db)
		// act
		result, err := rp.SearchProducts(internal.ProductQuery{Id: 1})
		// assert
		require.Nil(t, err)
		require.Equal(t, 1, len(result))
		require.Equal(t, product, result[1])
	})

	t.Run("product not found", func(t *testing.T) {
		// arrange
		product := internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "test_1",
				Price:       200.0,
				SellerId:    1,
			},
		}
		db := map[int]internal.Product{
			1: product,
		}
		rp := repository.NewProductsMap(db)
		// act
		result, err := rp.SearchProducts(internal.ProductQuery{Id: 2})
		// assert
		require.Nil(t, err)
		require.Equal(t, 0, len(result))
	})

	t.Run("return everything if not query passed", func(t *testing.T) {
		// arrange
		db := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "test_1",
					Price:       200.0,
					SellerId:    1,
				},
			},
			2: {
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "test_2",
					Price:       300.0,
					SellerId:    2,
				},
			},
		}
		rp := repository.NewProductsMap(db)
		// act
		result, err := rp.SearchProducts(internal.ProductQuery{})
		// assert
		require.Nil(t, err)
		require.Equal(t, len(db), len(result))
		require.Equal(t, db, result)
	})
}
