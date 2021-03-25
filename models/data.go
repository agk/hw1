package models

var DB []Item

type Item struct {
	ID		int	`json:"id"`
	Title	string	`json:"title"`
	Amount	int	`json:"amount"`
	Price	float64	`json:"price"`
}


func init() {
	item1 := Item {
		ID:	1,
		Title:	"New Item",
		Amount: 100,
		Price: 12.55,
	}
	DB = append(DB, item1)
}

func FindItemById(id int) (Item, bool) {
	var item Item
	var found bool = false
	for _, i := range DB {
		if i.ID == id {
			item = i
			found = true
			break
		}
	}
	return item, found
}

func FindAndReplaceItemById(id int, newItem Item) (DB []Item) {
	for index, data := range DB {
		if data.ID == id {
			DB[index] = newItem
			break
		}
	}
	
	return DB
}
