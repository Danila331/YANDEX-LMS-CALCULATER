package pkg

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func CalculateExpression(expression string) (float64, error) {
	// Создаем новый экземпляр вычислителя
	eval, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}

	// Вычисляем значение выражения
	result, err := eval.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	// Преобразуем результат в тип float64
	resultFloat, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("результат не является числом")
	}

	return resultFloat, nil
}
