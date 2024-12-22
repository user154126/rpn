package calculation

import (
	"strconv"
	"strings"
	"unicode"
)

// Calc принимает строку выражения и возвращает вычисленное значение или ошибку
func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

		// Проверка на наличие некорректных символов
		for _, char := range expression {
			if !unicode.IsDigit(char) && char != '.' && char != '+' && char != '-' && char != '*' && char != '/' && char != '(' && char != ')' {
				return 0, ErrInvalidCharacter
			}
		}
		
	var stack []float64
	var opStack []rune
	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	var numberBuffer []rune

	for i := 0; i < len(expression); i++ {
		char := rune(expression[i])

		// tсли цифра или точка, собираем число
		if unicode.IsDigit(char) || char == '.' {
			numberBuffer = append(numberBuffer, char)
			continue
		}

		// если число накопилось в буфере, добавляем его в стек
		if len(numberBuffer) > 0 {
			num, err := strconv.ParseFloat(string(numberBuffer), 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
			numberBuffer = nil // Очищаем буфер
		}

		// если откр скобка
		if char == '(' {
			opStack = append(opStack, char)
			continue
		}

		// если закр скобка выполн операции до открывающей скобки
		if char == ')' {
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				if len(stack) < 2 {
					return 0, ErrNotEnoughOperands
				}
				res, err := applyOperator(&stack, &opStack)
				if err != nil {
					return 0, err
				}
				stack = append(stack, res)
			}
			if len(opStack) == 0 {
				return 0, ErrMismatchedParentheses
			}
			opStack = opStack[:len(opStack)-1] // Убираем '('
			continue
		}

		// если оператор, обрабатываем приоритеты
		if precedence[char] > 0 {
			for len(opStack) > 0 && precedence[char] <= precedence[opStack[len(opStack)-1]] {
				if len(stack) < 2 {
					return 0, ErrNotEnoughOperands
				}
				res, err := applyOperator(&stack, &opStack)
				if err != nil {
					return 0, err
				}
				stack = append(stack, res)
			}
			opStack = append(opStack, char)
		}
	}

	// если числа остались в буфере, добавляем их в стек
	if len(numberBuffer) > 0 {
		num, err := strconv.ParseFloat(string(numberBuffer), 64)
		if err != nil {
			return 0, err
		}
		stack = append(stack, num)
	}

	// применяем оставшиеся операторы
	for len(opStack) > 0 {
		if len(stack) < 2 {
			return 0, ErrNotEnoughOperands
		}
		res, err := applyOperator(&stack, &opStack)
		if err != nil {
			return 0, err
		}
		stack = append(stack, res)
	}

	if len(stack) != 1 {
		return 0, ErrExtraOperands
	}

	return stack[0], nil
}

// applyOperator применяет оператор к двум верхним элементам стека
func applyOperator(stack *[]float64, opStack *[]rune) (float64, error) {
	if len(*stack) < 2 {
		return 0, ErrNotEnoughOperands
	}

	b := (*stack)[len(*stack)-1]
	a := (*stack)[len(*stack)-2]
	*stack = (*stack)[:len(*stack)-2] // Убираем два последних операнда

	op := (*opStack)[len(*opStack)-1]
	*opStack = (*opStack)[:len(*opStack)-1] // Убираем оператор

	var result float64
	switch op {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		result = a / b
	default:
		return 0, ErrInvalidOperator
	}

	return result, nil
}
