package arraySlicing

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, amount float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: amount}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, t Transaction) Account {
	if t.From == a.Name {
		a.Balance -= t.Sum
	}
	if t.To == a.Name {
		a.Balance += t.Sum
	}
	return a
}
