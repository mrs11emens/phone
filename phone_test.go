package Phone

import (
	"testing"
)

func TestPhone(t *testing.T) {
	var v string
	v = "79121115200"
	expectedPhone := "7 (912) 111-52-00"

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

}
