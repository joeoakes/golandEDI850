package main

import (
	"fmt"
	"strings"
)

// Define structures to hold the information
type Item struct {
	ItemNumber string
	Quantity   int
	Price      float64
}

type Order struct {
	ISA   string
	GS    string
	ST    string
	BEG   string
	Items []Item
	SE    string
	GE    string
	IEA   string
}

// Function to generate a PO1 segment for each item
func createPO1Segment(item Item) string {
	return fmt.Sprintf("PO1*%s*%d*EA*%.2f**IN*Item001*VN*ABC123", item.ItemNumber, item.Quantity, item.Price)
}

// Function to generate the EDI 850 document based on the order information
func createEDI850(order Order) string {
	segments := []string{
		order.ISA,
		order.GS,
		order.ST,
		order.BEG,
	}

	// Dynamically add PO1 segments based on items in the order
	for _, item := range order.Items {
		po1Segment := createPO1Segment(item)
		segments = append(segments, po1Segment)
	}

	segments = append(segments, order.SE, order.GE, order.IEA)

	return strings.Join(segments, "\n")
}

func main() {
	// Example usage
	order := Order{
		ISA: "ISA*00*          *00*          *01*SenderID     *01*ReceiverID   *231025*0934*U*00401*000000123*0*P*>",
		GS:  "GS*PO*SenderCode*ReceiverCode*20231025*0934*123*X*004010",
		ST:  "ST*850*0001",
		BEG: "BEG*00*SA*1001**20231025",
		Items: []Item{
			{"1", 5, 10.00},
			{"2", 3, 20.00},
		},
		SE:  "SE*6*0001",
		GE:  "GE*1*123",
		IEA: "IEA*1*000000123",
	}

	edi850 := createEDI850(order)
	fmt.Println(edi850)
}
