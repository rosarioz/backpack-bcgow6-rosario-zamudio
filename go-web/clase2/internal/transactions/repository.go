package transactions

import "fmt"

//Repositorio, debe tener el acceso a la variable guardada en memoria.
//Se debe crear el archivo repository.go
//Se debe crear la estructura de la entidad
//Se deben crear las variables globales donde guardar las entidades
//Se debe generar la interface Repository con todos sus métodos
//Se debe generar la estructura repository
//Se debe generar una función que devuelva el Repositorio
//Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)


type Transaction struct {
	ID 				int		`json:"id"`
	TransactionCode string	`json:"code"`
	Coin 			string	`json:"coin"`
	Amount 			float64	`json:"amount"`
	IssuingCompany 	string	`json:"company"`
	Receiver 		string	`json:"receiver"`
	TransactionDate string	`json:"date"`
}

var transactions []Transaction
var lastID int

type Repository interface{
	GetAll()([]Transaction, error) // supongo que tambien funciona poniendo transaction en vez de []Transaction
	Store(id int, code, coin string, amount float64, company, receiver, date string)
	LastID()(int, error)
	Update() (id int, code, coin string, amount float64, company, receiver, date string)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, code, coin string, amount float64, company, receiver, date string)(Transaction, error){
	t:= Transaction{id, code, coin, amount, company, receiver, date}
	transactions = append(transactions, t)
	lastID = t.ID
	return t, nil
}

func (r *repository) GetAll()([]Transaction, error){
	return transactions, nil
}

func (r *repository) LastID()(int, error){
	return lastID, nil
}

func (r *repository) Update(id int, code, coin string, amount float64, company, receiver, date string) (Transaction, error){
	t:= Transaction{TransactionCode: code,Coin: coin, Amount: amount,IssuingCompany: company,Receiver: receiver, TransactionDate: date}
	updated:= false
	for i := range transactions{
		if transactions[i].ID == id{
			t.ID = id
			transactions[i] = t
			updated = true
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("Transaccion %d no encontrada", id)
	}
	return t, nil
}


