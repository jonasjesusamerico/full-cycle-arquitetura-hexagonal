package application

type ProductService struct {
	Persistense ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistense: persistence}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistense.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProdut()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistense.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistense.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistense.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}
