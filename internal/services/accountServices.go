package services

import (
	"fmt"
	"module/internal/models"
	"sync"

	log "github.com/sirupsen/logrus"
)

// интерфейс для аккаунтов
type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

// стуктура аккаунта
type Account struct {
	ID      int64   `json:"id"`
	Balance float64 `json:"balance"`
	mu      sync.Mutex
}

// хранение аккаунтов
var accounts = make([]*Account, 0, 10)

// мьютекс для обращения к аккаунтам, чтобы например не создалось 10 аккаунтов одновременно с однаковыми айди.
var accountsMu sync.Mutex

// добавить деньги на счёт
func (a *Account) Deposit(amount float64) error {

	// мьютекс для потокобезопасности. defer для анлока если вдруг случится краш.
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += amount

	// вывод логов
	id := fmt.Sprint(a.ID)
	newAmount := fmt.Sprint(amount)
	log.Info("deposit " + newAmount + " at account id: " + id)

	return nil
}

// снять деньги со счёта
func (a *Account) Withdraw(amount float64) error {

	// мьютекс для потокобезопасности. defer для анлока если вдруг случится краш.
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance -= amount

	// вывод логов
	id := fmt.Sprint(a.ID)
	newAmount := fmt.Sprint(amount)
	log.Info("withdraw " + newAmount + " from account id: " + id)

	return nil
}

// посмотреть сколько денег на счёте
func (a *Account) GetBalance() float64 {

	// мьютекс для потокобезопасности. defer для анлока если вдруг случится краш.
	a.mu.Lock()
	defer a.mu.Unlock()

	// вывод логов
	id := fmt.Sprint(a.ID)
	balance := fmt.Sprint(a.Balance)
	log.Info("show balance " + balance + " from account id: " + id)

	return a.Balance
}

// создание нового аккаунта
func CreateAccount() {

	// мьютекс для потокобезопасности. defer для анлока если вдруг случится краш.
	accountsMu.Lock()
	defer accountsMu.Unlock()

	// создание аккаунта с 0 балансом, а айди автонумеруется по количеству аккаунтов.
	var newAccount Account
	newAccount.ID = int64(len(accounts))
	newAccount.Balance = 0
	accounts = append(accounts, &newAccount)

	// вывод логов
	id := fmt.Sprint(newAccount.ID)
	log.Info("create new account with id: " + id)

}

// получить аккаунт по айди, если аккаунта нет, то вернуть пустой и ошибку
func GetBasicAccount(id int64) (*Account, error) {

	// если аккаунта под таким номером нет, то вернуть ошибку
	var curLen = len(accounts)
	if id >= int64(curLen) {
		return nil, models.ResponseBadRequest()
	}

	var curAccount = accounts[id]
	return curAccount, nil
}

// положить деньги на аккаунт. (выполняется в горутине)
func DepositeToBasicAccount(ch chan error, id int64, amount float64) {

	// тут идёт работа с аккаунтом через интерфейс, и если поменять GetBasicAccount на GetSuperAccount то будет работа уже с другой структурой
	// если у структуры нет всех методов интерфейса, то будет ошибка.
	var curAcc BankAccount
	curAcc, err := GetBasicAccount(id)
	if err != nil {
		ch <- err
		return
	}

	// внести деньги на аккаунт.
	err = curAcc.Deposit(amount)
	if err != nil {
		ch <- err
		return
	}

	ch <- models.ResponseGood()
	return

}

// снять деньги с аккаунта. (выполняется в горутине)
func WithdrawByBasicAccount(ch chan error, id int64, amount float64) {

	// тут идёт работа с аккаунтом через интерфейс, и если поменять GetBasicAccount на GetSuperAccount то будет работа уже с другой структурой
	// если у структуры нет всех методов интерфейса, то будет ошибка.
	var curAcc BankAccount
	curAcc, err := GetBasicAccount(id)
	if err != nil {
		ch <- err
		return
	}

	// снять деньги с аккаунта.
	err = curAcc.Withdraw(amount)
	if err != nil {
		ch <- err
		return
	}

	ch <- models.ResponseGood()
	return
}

// посмотреть деньги на аккаунте. (выполняется в горутине)
func BalanceFromBasicAccount(ch chan error, id int64) {

	// тут идёт работа с аккаунтом через интерфейс, и если поменять GetBasicAccount на GetSuperAccount то будет работа уже с другой структурой
	// если у структуры нет всех методов интерфейса, то будет ошибка.
	var curAcc BankAccount
	curAcc, err := GetBasicAccount(id)
	if err != nil {
		ch <- models.ResponseErrorAtServer()
		return
	}

	// получить баланс с аккаунта.
	value := curAcc.GetBalance()
	ch <- models.ResponseBalanceGood(value)
	return
}
