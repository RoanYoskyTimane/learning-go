package main

import "fmt"

type Combatant interface {
	Attack() int
	Defend() int
	TakeDamage(amount int)
	IsDead() bool
	GetName() string
	HealDamage(amount int)
}

// Dragon Implementation
type Dragon struct {
	Name      string
	Health    int
	FirePower int
	Scales    int
}

func (d *Dragon) Attack() int {
	return d.FirePower
}

func (d *Dragon) Defend() int {
	return d.Scales
}

func (d *Dragon) IsDead() bool {
	return d.Health <= 0
}

func (d *Dragon) GetName() string {
	return d.Name
}

func (d *Dragon) TakeDamage(amount int) {
	netDamage := amount - d.Defend()
	if netDamage < 1 {
		netDamage = 1
	}
	d.Health -= netDamage
	fmt.Printf("%s takes %d damage! (Health: %d)\n", d.Name, netDamage, d.Health)
}

func (d *Dragon) HealDamage(amount int) {
	d.Health += amount
}

// Cleric Implementation
type Healer interface {
	Heal(target Combatant)
}

type Cleric struct {
	Name     string
	Health   int
	Strength int
	Wisdom   int
}

func (cl *Cleric) Attack() int {
	return cl.Strength
}

func (cl *Cleric) Defend() int {
	return cl.Wisdom / 2
}

func (cl *Cleric) IsDead() bool {
	return cl.Health <= 0
}

func (cl *Cleric) GetName() string {
	return cl.Name
}

func (cl *Cleric) TakeDamage(amount int) {
	cl.Health -= amount
	fmt.Printf("%s takes %d damage! (Health: %d)\n", cl.Name, amount, cl.Health)
}

func (cl *Cleric) Heal(target Combatant) {
	healAmount := cl.Wisdom * 2
	target.HealDamage(healAmount)
	fmt.Printf("%s casts Heal on %s for +%d HP!\n", cl.Name, target.GetName(), healAmount)
}

func (cl *Cleric) HealDamage(amount int) {
	cl.Health += amount
}

func BattleReport(c Combatant) {
	fmt.Printf("This combatant attacks for %d and defends with %d\n", c.Attack(), c.Defend())
}

func BattleArena(c1 Combatant, c2 Combatant) {
	fmt.Printf("\n--- Arena: %s VS %s ---\n", c1.GetName(), c2.GetName())
	round := 1

	// Dynamic loop that runs until someone dies
	for !c1.IsDead() && !c2.IsDead() {
		fmt.Printf("\n[Round %d]\n", round)

		// c1 attacks c2
		c2.TakeDamage(c1.Attack())
		if c2.IsDead() {
			break
		}

		// c2 attacks c1
		c1.TakeDamage(c2.Attack())

		round++
	}

	if c1.IsDead() {
		fmt.Printf("\n %s is victorious!\n", c2.GetName())
	} else {
		fmt.Printf("\n %s is victorious!\n", c1.GetName())
	}
}

func SupportRound(c Combatant, h Healer) {
	fmt.Println("--- The Support Round ---")
	h.Heal(c)
}

func main() {
	dragon := &Dragon{Name: "Smaug", Health: 80, FirePower: 20, Scales: 5}
	cleric := &Cleric{Name: "Faith", Health: 50, Strength: 8, Wisdom: 12}

	SupportRound(dragon, cleric)
	BattleArena(cleric, dragon)
}
