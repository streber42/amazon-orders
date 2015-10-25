package main

import "fmt"
import "encoding/csv"
import "os"

type Shipment struct {
	shipment_date string
	order_id      string
	total_charged string
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

	shipments_reader := csv.NewReader(shipmentsfile)
	rawCSVdata, err := shipments_reader.ReadAll()
	items_reader := csv.NewReader(itemsfile)
	rawItemsData, err := items_reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Shipment
	var allRecords []Shipment

	for _, each := range rawCSVdata {
		oneRecord.shipment_date = each[6]
		oneRecord.order_id = each[1]
		oneRecord.total_charged = each[20]
		allRecords = append(allRecords, oneRecord)
	}

	//fmt.Println(allRecords)

	var oneItem Item
	var allItems []Item

	for _, each := range rawItemsData {
		oneItem.shipment_date = each[16]
		oneItem.order_id = each[1]
		oneItem.title = each[2]
		oneItem.item_total = each[27]
		allItems = append(allItems, oneItem)
	}
	//fmt.Println(allItems)

	for i := len(allRecords) - 1; i > 0; i-- {
		fmt.Printf("Date: %s Order Id: %s Total: %s\n", allRecords[i].shipment_date, allRecords[i].order_id, allRecords[i].total_charged)
		for j := 0; j < len(allItems); j++ {
			if allItems[j].shipment_date == allRecords[i].shipment_date && allItems[j].order_id == allRecords[i].order_id {
				fmt.Printf("%s %s\n", allItems[j].title, allItems[j].item_total)
			}
		}
		fmt.Println("")
	}
}
