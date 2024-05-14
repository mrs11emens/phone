package Phone

import (
	"errors"
	"strings"
	"unicode"
)

func VvodPhone(phone string) string {
	// Remove letters from the phone string
	phone = strings.Map(func(r rune) rune {
		if !unicode.IsDigit(r) {
			return -1
		}
		return r // Keep other characters unchanged
	}, phone)

	return phone
}

type Phone struct {
	country string
	prefix  string
	part1   string
	part2   string
	part3   string
}

func NewPhone(phone string) (*Phone, error) {
	// вырезать все символы кроме цифр
	phon := VvodPhone(phone)

	//проверить длину строки на 11 символов
	if len(phon) != 11 {
		return nil, errors.New("длина номера должна быть 11 символов")
	}
	//разбить строку на части
	country := phon[:1]
	prefix := phon[1:4]
	part1 := phon[4:7]
	part2 := phon[7:9]
	part3 := phon[9:11]

	//вернуть обьект phone
	obj := &Phone{country: country,
		prefix: prefix,
		part1:  part1,
		part2:  part2,
		part3:  part3,
	}
	return obj, nil

}

// теги:
//
//	{country} - код страны,
//	{prefix} - префикс,
//
// {part1} - первая часть номера,
// {part2} - вторая часть номера,
// {part3} - третья часть номера.
func (p *Phone) Format(s string) string {
	s = strings.ReplaceAll(s, "{country}", p.country)
	s = strings.ReplaceAll(s, "{prefix}", p.prefix)
	s = strings.ReplaceAll(s, "{part1}", p.part1)
	s = strings.ReplaceAll(s, "{part2}", p.part2)
	s = strings.ReplaceAll(s, "{part3}", p.part3)

	return s
}
