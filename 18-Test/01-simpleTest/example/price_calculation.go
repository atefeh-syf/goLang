package example

func CalculateRoomPrice(roomType string, nights int, personCount int) (finalPrice int) {
	price := 0
	switch roomType {
	case "standard":
		price = nights * 140000 * personCount
	case "double":
		price = nights * 220000 * personCount
	case "suite":
		price = nights * 350000 * personCount
		// default:
		// 	panic("room type not supported")
	}
	tax := float64(price) * 0.09
	finalPrice = price + int(tax)
	return
}
