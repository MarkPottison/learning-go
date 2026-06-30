package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	var name string
	fmt.Println("Как вас зовут?")
	if _, err := fmt.Scan(&name); err != nil {

		log.Fatal("Ошибка ввода имени")
	}
	for _, char := range name {
		if !unicode.IsLetter(char) {
			log.Fatal("Ошибка: имя не должно содержать цифры и символы!")
		}
	}

	var money float64
	fmt.Println("Сколько у вас денег в кошельке?")
	if _, err := fmt.Scan(&money); err != nil {
		log.Fatal("Ошибка: введите число")
	}

	var price = 130.90

	var quantity int
	fmt.Println("Сколько вы хотите чашек кофе?")
	if _, err := fmt.Scan(&quantity); err != nil {
		log.Fatal("Ошибка: введите целое число")
	}
	total := price * float64(quantity)
	fmt.Println("Итоговая сумма " + fmt.Sprintf("%.2f", total))

	if money >= total {
		change := money - total
		fmt.Println("Отлично, ваша сдача " + fmt.Sprintf("%.2f", change))
		if quantity >= 5 || total > 800 {
			fmt.Println("Вы выиграли золотую наклейку!")
		}
		switch quantity {
		case 1:
			fmt.Println("Хорошее начало")
		case 2:
			fmt.Println("Двойная порция!")
		case 3:
			fmt.Println("Ого сколько кофе!")
		default:
			fmt.Println("Вы сегодня в ударе!")
		}
	} else {
		fmt.Println("Извини, тебе не хватает денег")
	}
	fmt.Println("\nНажмите Enter, чтобы выйти...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	_, _ = reader.ReadString('\n')
}
