package main

import "fmt"

func dayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return "Invalid number."
	}
}

func dayOfTheWeek2(number int) string {
	var day string

	switch {
	case number == 1:
		day = "Monday"
		fallthrough //ignora a condição que entrou e executa a próxima(uso muito raro)
	case number == 2:
		day = "Tuesday"
	case number == 3:
		day = "Wednesday"
	case number == 4:
		day = "Thursday"
	case number == 5:
		day = "Friday"
	case number == 6:
		day = "Saturday"
	default:
		day = "Invalid number."
	}

	return day
}

func main() {
	fmt.Println("Swtichs")

	//forma tradicional do switch
	dayOfTheWeek := dayOfTheWeek(1)
	fmt.Println(dayOfTheWeek)

	//outra forma de usar o switch (condições dentro dos cases)
	dayOfTheWeek2 := dayOfTheWeek2(1)
	fmt.Println(dayOfTheWeek2)
}
