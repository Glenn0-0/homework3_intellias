package main

import "fmt"

var (
	applePrice float32 = 5.99 //declaring price per apple (could have declared those as "const" but why would we bother :/ )
	pearPrice float32 = 7 //declaring price per pear
	availableMoney float32 = 23 //the amount of money we possess
	appleQuantity float32 = 9 //the amount of apples we need to buy in task1
	pearQuantity float32 = 8 //the amount of pears we need to buy in task1
	quantity float32 = 2 //the amount of each fruit we have to buy in task4
	answer string //a variable holding if...else statement result for task4
)

func main() {
	
	//task 1
	fmt.Printf("Task 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\nTo buy %v apples and %v pears we need %v uah.\n\n", appleQuantity, pearQuantity, (applePrice)*(appleQuantity)+(pearPrice)*(pearQuantity))

	//task2
	fmt.Printf("Task 2. Скільки груш ми можемо купити?\nRegarding the amount of money we have (%v uah), we can only buy %v pears.\n\n", availableMoney, int((availableMoney)/(pearPrice)))

	//task3
	fmt.Printf("Task 3. Скільки яблук ми можемо купити?\nRegarding the amount of money we have (%v uah), we can only buy %v apples.\n\n", availableMoney, int((availableMoney)/(applePrice)))

	//task4
	if (applePrice*quantity + pearPrice*quantity) <= availableMoney {
		answer = "Yes, we can."  //i.e. comparing the amount of money that is required to buy certain quantity of fruits and the amount we possess
	} else {
		answer = "No, we don't have enough money available."
	}
	fmt.Printf("Task 4. Чи ми можемо купити 2 груші та 2 яблука?\n%v\n\n", answer)
}

// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
//
// Задача:
// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
// Xід опишіть, я маю на увазі наступне:
// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані 
// відповіді на поставлені вище запитання. Перед відповідю треба роздрукувати саме питання.
// Правильно обирайте типи даних та назви змінних чи констант.
// Публікація:
// Створити папку в своєму репозиторії в github та завантажити туди main.go файл, в котрому буде зроблено дане завдання.

