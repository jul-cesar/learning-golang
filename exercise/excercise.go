package exercise

import "fmt"

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

func (player *Player) PickUpItem(item Item) {
	player.Inventory = append(player.Inventory, item)
}
func (player *Player) DropItem(item *Item) {
	delete(player.Inventory(item))
}
func (player *Player) UseItem(item *Item) {
	if item.Name == "Sword" {
		fmt.Println("Using sword!")
	}
}

func excercise() {

}
