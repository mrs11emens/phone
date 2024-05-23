# Программа по работе с форматом телефона

Пакет Phone для обработки и форматирования телефона.

---
 # Пример

```go
    // Изначальные параметры телефона, в formatted изменятся
	v := "7 (912) 123-45-67"

	// Создаем экземпляр структуры Phone с помощью функции NewPhone
	phone, err := NewPhone(v)
	if err != nil {
		t.Errorf("Error creating phone object: %v", err)
	}
	// Здесь задаётся шаблон телефона
	formatirovanie := "+{country} {prefix} {part1}-{part2}-{part3}"
	// Получаем форматированный номер телефона с помощью метода Format
	formatted := phone.Format(formatirovanie)
	// t.Log(formatted) выведет +7 912 123-45-67

```
Данный тест выведет ok(сравнение с exceptedPhone будет выполнено, соответственно форматированный номер будет идентичен exceptedPhone)

По желанию можно задать собственный шаблон, указав его. В моем примере, шаблон указан в переменной formatirovanie
