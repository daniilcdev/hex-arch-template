package api

import "github.com/daniilcdev/hex-arch-template/hex/internal/ports"

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (app Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := app.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = app.db.AddToHistory(answer, "addition")

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (app Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := app.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = app.db.AddToHistory(answer, "subtraction")

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (app Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := app.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = app.db.AddToHistory(answer, "multiplication")

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (app Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := app.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = app.db.AddToHistory(answer, "division")

	if err != nil {
		return 0, err
	}

	return answer, nil
}
