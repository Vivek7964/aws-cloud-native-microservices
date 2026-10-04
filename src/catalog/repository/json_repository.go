package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/Vivek7964/microservice-demo/catalog/model"
	"github.com/prometheus/client_golang/prometheus"
)

type jsonProduct struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	Price       int      `json:"price"`
	Count       int      `json:"count"`
	Tags        []string `json:"tags"`
}

type jsonCatalogData struct {
	Products []jsonProduct `json:"products"`
	Tags     []model.Tag   `json:"tags"`
}

type jsonRepository struct {
	products []model.Product
	tags     []model.Tag
}

func newJSONRepository() (Repository, error) {
	data, err := ioutil.ReadFile("products.json")
	if err != nil {
		return nil, err
	}

	var catalog jsonCatalogData

	err = json.Unmarshal(data, &catalog)
	if err != nil {
		return nil, err
	}

	products := make([]model.Product, 0, len(catalog.Products))

	for _, product := range catalog.Products {
		products = append(products, model.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			ImageURL:    product.ImageURL,
			Price:       product.Price,
			Count:       product.Count,
			Tags:        product.Tags,
			TagString:   strings.Join(product.Tags, ","),
		})
	}

	return &jsonRepository{
		products: products,
		tags:     catalog.Tags,
	}, nil
}

func (r *jsonRepository) List(
	tags []string,
	order string,
	pageNum int,
	pageSize int,
) ([]model.Product, error) {

	var products []model.Product

	for _, product := range r.products {

		// No tag filter.
		if len(tags) == 0 {
			products = append(products, product)
			continue
		}

		// Match products containing any requested tag.
		for _, requestedTag := range tags {
			if contains(product.Tags, requestedTag) {
				products = append(products, product)
				break
			}
		}
	}

	// Apply ordering.
	switch strings.ToLower(order) {
	case "name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})

	case "price":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})

	case "count":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Count < products[j].Count
		})

	case "id":
		sort.Slice(products, func(i, j int) bool {
			return products[i].ID < products[j].ID
		})
	}

	products = paginate(products, pageNum, pageSize)

	return products, nil
}

func (r *jsonRepository) Count(tags []string) (int, error) {

	count := 0

	for _, product := range r.products {

		if len(tags) == 0 {
			count++
			continue
		}

		for _, requestedTag := range tags {
			if contains(product.Tags, requestedTag) {
				count++
				break
			}
		}
	}

	return count, nil
}

func (r *jsonRepository) Get(id string) (*model.Product, error) {

	for _, product := range r.products {
		if product.ID == id {
			result := product
			return &result, nil
		}
	}

	return nil, errors.New("not found")
}

func (r *jsonRepository) Tags() ([]model.Tag, error) {
	return r.tags, nil
}

// JSON/in-memory storage does not use database metrics.
// Return a no-op collector to satisfy the Repository interface.
func (r *jsonRepository) Collector() prometheus.Collector {
	return noopCollector{}
}

func (r *jsonRepository) ReaderCollector() prometheus.Collector {
	return noopCollector{}
}

type noopCollector struct{}

func (noopCollector) Describe(ch chan<- *prometheus.Desc) {
}

func (noopCollector) Collect(ch chan<- prometheus.Metric) {
}

func paginate(products []model.Product, pageNum, pageSize int) []model.Product {

	if pageNum == 0 || pageSize == 0 {
		return []model.Product{}
	}

	start := (pageNum - 1) * pageSize

	if start >= len(products) {
		return []model.Product{}
	}

	end := start + pageSize

	if end > len(products) {
		end = len(products)
	}

	return products[start:end]
}
