package v1_test

import (
	"bytes"
	"encoding/json"
	"eulabs_challenger/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestProductController_CreateProduct(t *testing.T) {

	testCases := []struct {
		name                   string
		productRequest         *dto.CreateProductRequest
		urlRequested           string
		expectedHttpStatusCode int
		facadeProductResponse  *dto.CreateProductResponse
		facadeProductError     error
		expectedError          error
	}{
		{
			name: "OK",
			productRequest: &dto.CreateProductRequest{
				Name:        "Product 1",
				Description: "Product 1 description",
				Price:       10.0,
			},
			urlRequested:           "/eulabs/v1/product",
			expectedHttpStatusCode: 201,
			facadeProductResponse: &dto.CreateProductResponse{
				ID: 1,
			},
			facadeProductError: nil,
			expectedError:      nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ProductControllerMock.On("CreateProduct", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeProductResponse,
					tc.facadeProductError,
				)

			data, err := json.Marshal(tc.productRequest)
			assert.NoError(t, err)
			reader := bytes.NewReader(data)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tc.urlRequested, reader)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.CreateProductResponse{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeProductResponse.ID, getParsingJson.ID)
				return
			}

			errorResponse := &dto.Error{}
			err = json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}

}

func TestProductController_GetProductByID(t *testing.T) {
	testCases := []struct {
		name                   string
		productID              uint64
		urlRequested           string
		expectedHttpStatusCode int
		facadeProductResponse  *dto.GetByProductIDResponse
		facadeProductError     error
		expectedError          error
	}{
		{
			name:                   "OK",
			productID:              uint64(1),
			urlRequested:           "/eulabs/v1/product/1",
			expectedHttpStatusCode: 200,
			facadeProductResponse: &dto.GetByProductIDResponse{
				ID:          1,
				Name:        "Product 1",
				Description: "Product 1 description",
				Price:       10.0,
				CreateAt:    time.Now(),
				UpdateAt:    time.Now(),
			},
			facadeProductError: nil,
			expectedError:      nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ProductControllerMock.On("GetByIDProduct", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeProductResponse,
					tc.facadeProductError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.GetByProductIDResponse{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeProductResponse.ID, getParsingJson.ID)
				return
			}

			errorResponse := &dto.Error{}
			err := json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}

}

func TestProductController_GetAllProducts(t *testing.T) {
	testCases := []struct {
		name                   string
		urlRequested           string
		expectedHttpStatusCode int
		facadeProductResponse  *dto.GetAllProductsResponse
		facadeProductError     error
		expectedError          error
	}{
		{
			name:                   "OK",
			urlRequested:           "/eulabs/v1/product",
			expectedHttpStatusCode: 200,
			facadeProductResponse: &dto.GetAllProductsResponse{
				Products: []dto.Product{
					{
						ID:          uint64(1),
						Name:        "Product 1",
						Description: "Product 1 description",
						Price:       10.0,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
					{
						ID:          uint64(2),
						Name:        "Product 2",
						Description: "Product 2 description",
						Price:       9.99,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				},
				Pagination: dto.ProductPagination{
					Limit:  10,
					Offset: 0,
					Total:  2,
				},
			},
			facadeProductError: nil,
			expectedError:      nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ProductControllerMock.On("GetAllProducts", mock.AnythingOfType("*context.emptyCtx"), mock.Anything, mock.Anything).
				Return(
					tc.facadeProductResponse,
					tc.facadeProductError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.GetAllProductsResponse{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeProductResponse.Products[0].ID, getParsingJson.Products[0].ID)
				return
			}

			errorResponse := &dto.Error{}
			err := json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}
}

func TestProductController_UpdateProduct(t *testing.T) {
	testCases := []struct {
		name                   string
		productUpdateRequest   *dto.UpdateProductRequest
		productID              uint64
		urlRequested           string
		expectedHttpStatusCode int
		facadeProductError     error
		expectedError          error
	}{
		{
			name: "OK",
			productUpdateRequest: &dto.UpdateProductRequest{
				Name:        "Product 1 updated",
				Description: "Product 1 description updated",
				Price:       10.0,
			},
			productID:              1,
			urlRequested:           "/eulabs/v1/product/1",
			expectedHttpStatusCode: 200,
			facadeProductError:     nil,
			expectedError:          nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ProductControllerMock.On("UpdateProduct", mock.AnythingOfType("*context.emptyCtx"), mock.Anything, mock.Anything).
				Return(
					tc.facadeProductError,
				)

			data, err := json.Marshal(tc.productUpdateRequest)
			assert.NoError(t, err)
			reader := bytes.NewReader(data)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", tc.urlRequested, reader)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				assert.Equal(t, responseString, "\"Product updated successfully\"\n")
				return
			}

			errorResponse := &dto.Error{}
			err = json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}
}

func TestProductController_DeleteProduct(t *testing.T) {
	testCases := []struct {
		name                   string
		productID              uint64
		urlRequested           string
		expectedHttpStatusCode int
		facadeProductError     error
		expectedError          error
	}{
		{
			name:                   "OK",
			productID:              1,
			urlRequested:           "/eulabs/v1/product/1",
			expectedHttpStatusCode: 200,
			facadeProductError:     nil,
			expectedError:          nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ProductControllerMock.On("DeleteProduct", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeProductError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				assert.Equal(t, responseString, "\"Product deleted successfully\"\n")
				return
			}

			errorResponse := &dto.Error{}
			err := json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}
}
