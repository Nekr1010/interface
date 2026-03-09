package structs

import "fmt"

type Apartment struct {
	address string
	Area    float64
	Floor   int
	Price   int
}

func (a *Apartment) ChangePrice(newPrice int) {
	a.Price = newPrice
}

func NewApartment(
	address string,
	area float64,
	floor int,
	prince int,
) Apartment {
	if address == "" {
		fmt.Println("Адрес не может быть пустым")
		return Apartment{}
	}
	if area < 10 {
		fmt.Println("Плошадь должна быть больше 10 метров")
		return Apartment{}
	}
	if floor < 0 || floor > 100 {
		fmt.Println("Этаж должен быть больше 0 и меньши либо равен 100")
		return Apartment{}
	}
	return Apartment{
		address: address,
		Area:    area,
		Floor:   floor,
		Price:   prince,
	}
}
