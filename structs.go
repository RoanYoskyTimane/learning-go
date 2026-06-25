package main

import (
	"fmt"
)

type Character struct {
	Name           string
	Health         int
	Strength       int
	Level          int
	Experience     int
	Inventory      []string
	EquippedWeapon *Item
}

func NewCharacter(name string) *Character {
	return &Character{
		name,
		100,
		10,
		1,
		0,
		[]string{},
		nil}
}

func (c *Character) DisplayStats() {
	fmt.Printf("The new character - health %d\n strength %d\n level %d\n", c.Health, c.Strength, c.Level)
}

func (c *Character) GainXP(amount int) {
	c.Experience += amount
	fmt.Printf("Gained %d XP! ", amount)

	if amount >= 100 {
		c.Level++
		c.Strength += 2
		c.Health += 20
		c.Experience -= 100
		fmt.Printf("Level %d reached! Stregth %d, Health %d", c.Level, c.Strength, c.Health)
	}

}

func (c *Character) AddItem(item string) {

	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i] == item {
			fmt.Printf("Item %s is already added to Inventory\n", item)
		}
	}
	c.Inventory = append(c.Inventory, item)
}

func (c *Character) ShowInventory() {
	fmt.Println("Inventory:", c.Name)
	for _, item := range c.Inventory {
		fmt.Println(item)
	}
}

type Party struct {
	Members []*Character
}

func (p *Party) addMember(character *Character) {
	p.Members = append(p.Members, character)
}

func (p *Party) TotalPartyHealth() int {
	totalHealth := 0
	for _, member := range p.Members {
		totalHealth += member.Health
	}
	return totalHealth
}

func (p *Party) FindStrongest() *Character {
	if len(p.Members) == 0 {
		return nil
	}

	strongest := p.Members[0]
	for _, member := range p.Members {
		if member.Strength > strongest.Strength {
			strongest = member
		}
	}
	return strongest
}

type Item struct {
	Name          string
	Durability    int
	MaxDurability int
	BonusStrength int
}

func (c *Character) EquipWeapon(newItem *Item) {
	if c.EquippedWeapon != nil {
		fmt.Printf("Unequipped %s\n", c.EquippedWeapon.Name)
		c.Strength -= c.EquippedWeapon.BonusStrength
	}
	c.EquippedWeapon = newItem
	c.Strength += newItem.BonusStrength
	fmt.Printf("Equipped %s! Strength is now %d.\n", newItem.Name, c.Strength)
}

func (c *Character) UseWeapon() {
	if c.EquippedWeapon == nil {
		fmt.Println("Fighting with bare fists!")
		return
	}

	c.EquippedWeapon.Durability -= 10
	fmt.Printf("Used %s! Durability dropped to %d.\n", c.EquippedWeapon.Name, c.EquippedWeapon.Durability)

	if c.EquippedWeapon.Durability < 0 {
		fmt.Printf("%s has broken!", c.EquippedWeapon.Name)
		c.EquippedWeapon.BonusStrength = 0
		c.EquippedWeapon = nil
	}
}

func main() {
	hero := NewCharacter("Roan")
	hero.DisplayStats()
	//hero.GainXP(150)
	hero.AddItem("Sword")
	hero.ShowInventory()
}
