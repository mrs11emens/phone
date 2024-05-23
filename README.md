# Программа по работе с форматом телефона

Пакет Phone для обработки и форматирования телефона.

---
 # Примеры

```go
var v string
	v = "79121115200"
	expectedPhone := "7 (912) 111-52-00"
	v = VvodPhone(v)

	// Создаем экземпляр структуры Phone с помощью функции NewPhone
	phone, err := NewPhone(v)
	if err != nil {
		t.Errorf("Error creating phone object: %v", err)
	}
	formatirovanie := "{country} ({prefix}) {part1}-{part2}-{part3}"
	// Получаем форматированный номер телефона с помощью метода Format
	formatted := phone.Format(formatirovanie)
	t.Log(formatted)

	// Сравниваем полученный формат с ожидаемым форматом
	if formatted != expectedPhone {
		t.Errorf("Expected '%v', got '%v'", expectedPhone, formatted)
	}
```
Данный тест выведет ok(сравнение с exceptedPhone будет выполнено, соответственно форматированный номер будет идентичен exceptedPhone)

По желанию можно сделать собственный формат, указав его. В моем примере, формат указан в переменной formatirovanie.
