package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"T-SHIRT": 20.00,
	"MUG":     7.50,
	"STICKER": 1.50,
	"HAT":     12.00,
	"BOOK":    30.00,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf(" - Item %s (Sale! Original: $%.2f, | Sale Price: $%.2f )\n", originalItemCode, basePrice, salePrice)
				return salePrice, true
			}
		}
		fmt.Printf(" - Item: %s (Product not found)\n", itemCode)
		return 0, false
	}

	return basePrice, true
}

func main() {
	fmt.Println("-------------Simple Sales Order Processor ----------------")

	orderItems := []string{
		"T-SHIRT",
		"MUG",
		"STICKER",
		"HAT",
		"BOOK",
		"SELLER",
		"RAGS",
		"GOLANG",
	}

	var subtotal float64
	fmt.Println("--------------Processing order items: --------------")

	for _, item := range orderItems {
		itemPrice, found := calculateItemPrice(item)
		if found {
			subtotal += itemPrice
		}
	}

	fmt.Printf("Subtotal price: $%.2f\n", subtotal)
}
