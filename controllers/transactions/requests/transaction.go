package requests

import (
	"project/business/transactions"
)

type Transaction struct {
	Payment_method_id int
	User_Id           int
	Plan_id           int
}

func (pay *Transaction) ToDomain() transactions.Transaction {
	return transactions.Transaction{
		Payment_method_id: pay.Payment_method_id,
		User_Id:           pay.User_Id,
		Plan_Id:           pay.Plan_id,
	}
}
