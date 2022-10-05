package transactions

type Service interface{
	GetAll()([]Transaction, error) // supongo que tambien funciona poniendo transaction en vez de []Transaction
	Store(id int, code, coin string, amount float64, company, receiver, date string)
	Update() (id int, code, coin string, amount float64, company, receiver, date string)
}

type service struct{
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll()([]Transaction, error){
	transactions, err:= s.repository.GetAll()
	if err != nil{
		return nil, err
	}
	return transactions, nil
}

func(s service) Store(code, coin string, amount float64, company, receiver, date string) (Transaction, error){
	lastID, err := s.repository.LastID()
	if err != nil {
		return Transaction{}, err
	}
	lastID ++

	transaction, err:= s.repository.Store(lastID, code, coin, amount, company, receiver, date)
	if err != nil{
		return Transaction{}, err
	}
	return transaction, nil
}

func (s service) Update(id int, code, coin string, amount float64, company, receiver, date string) (Transaction, error){
	return s.repository.Update(id, code, coin, amount, company, receiver, date)
}

