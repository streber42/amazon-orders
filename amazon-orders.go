package main

import "fmt"
import "encoding/csv"
import "os"

type Shipment struct {
	shipment_date         string
	order_id              string
	total_charged         string
	shipping_charge       string
	subtotal              string
	tax_before_promotions string
	total_promotions      string
	tax_charged           string
}

type Item struct {
	shipment_date string
	order_id      string
	title         string
	item_total    string
}

func main() {
	shipmentsfile, err := os.Open("shipments.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer shipmentsfile.Close()

	itemsfile, err := os.Open("items.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer itemsfile.Close()

	shipments_index := make(map[string]int, 30)
	items_index := make(map[string]int, 30)

	shipments_reader := csv.NewReader(shipmentsfile)
	shipments_headers, err := shipments_reader.Read()
	for index, each := range shipments_headers {
		shipments_index[each] = index
	}
	rawCSVdata, err := shipments_reader.ReadAll()

	items_reader := csv.NewReader(itemsfile)
	items_headers, err := items_reader.Read()
	for index, each := range items_headers {
		items_index[each] = index
	}
	rawItemsData, err := items_reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Shipment
	var allRecords []Shipment

	for _, each := range rawCSVdata {
		oneRecord.shipment_date = each[shipments_index["Shipment Date"]]
		oneRecord.order_id = each[shipments_index["Order ID"]]
		oneRecord.total_charged = each[shipments_index["Total Charged"]]
		oneRecord.subtotal = each[shipments_index["Subtotal"]]
		oneRecord.shipping_charge = each[shipments_index["Shipping Charge"]]
		oneRecord.tax_before_promotions = each[shipments_index["Tax Before Promotions"]]
		oneRecord.total_promotions = each[shipments_index["Total Promotions"]]
		oneRecord.tax_charged = each[shipments_index["Tax Charged"]]
		allRecords = append(allRecords, oneRecord)
	}

	//fmt.Println(allRecords)

	var oneItem Item
	var allItems []Item

	for _, each := range rawItemsData {
		oneItem.shipment_date = each[items_index["Shipment Date"]]
		oneItem.order_id = each[items_index["Order ID"]]
		oneItem.title = each[items_index["Title"]]
		oneItem.item_total = each[items_index["Item Total"]]
		allItems = append(allItems, oneItem)
	}
	//fmt.Println(allItems)

	for i := len(allRecords) - 1; i > 0; i-- {
		fmt.Printf("Date: %s Order Id: %s Subtotal: %s Shipping: %s TBP: %s TP: %s Tax: %s Total: %s\n", allRecords[i].shipment_date, allRecords[i].order_id, allRecords[i].subtotal, allRecords[i].shipping_charge, allRecords[i].tax_before_promotions, allRecords[i].total_promotions, allRecords[i].tax_charged, allRecords[i].total_charged)
		for j := 0; j < len(allItems); j++ {
			if allItems[j].shipment_date == allRecords[i].shipment_date && allItems[j].order_id == allRecords[i].order_id {
				fmt.Printf("%s %s\n", allItems[j].title, allItems[j].item_total)
			}
		}
		fmt.Println("")
	}
}
