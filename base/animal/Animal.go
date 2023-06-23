package animal

import "fmt"

type Health = float64;
type MaxHealth = Health;
type Hunger = float64;
type MaxHunger = Hunger;

const ( 
	DefaultMaxHealth MaxHealth = 100;
	DefaultMaxHunger MaxHunger = 100;
)

type Animal struct {
	health Health
	MaxHealth MaxHealth
	hunger Hunger
	MaxHunger MaxHunger
}

func (a *Animal) Reset() {
	a.health = a.MaxHealth;
	a.MaxHealth = DefaultMaxHealth;
	a.hunger = 0;
	a.MaxHunger = DefaultMaxHunger;
}

// ----------------
// Health Methods
// ----------------

func (a *Animal) SetHealth(health Health) Health {
	// If given health is less than 0, just set the animal health to 0
	if health < 0 {
		a.health = 0;
		return a.health;
	}
	// If given health is greater than the Maximum health allowed, set the health to Maximum health
	if health > a.MaxHealth {
		a.health = a.MaxHealth;
		return a.health;
	}

	// Set the health
	a.health = health;
	return a.health;
}

func (a *Animal) GetHealth() Health {
	return a.health;
}

func (a *Animal) IncreaseHealth(health Health) Health {
	a.SetHealth(a.GetHealth() + health)
	return a.GetHealth();
}

func (a *Animal) DecreaseHealth(health Health) Health {
	return a.IncreaseHealth(-health);
}

// ----------------
// Hunger Methods
// ----------------

func (a *Animal) SetHunger(hunger Hunger) Hunger {
	// If given hunger is more than Maximum hunger allowed, set the hunger to Maximum hunger
	if hunger > a.MaxHunger {
		a.hunger = a.MaxHunger;
		return a.hunger;
	}

	// Set the hunger
	a.hunger = hunger;
	return a.hunger;
}

func (a *Animal) GetHunger() Hunger {
	return a.hunger;
}

func (a *Animal) IncreaseHunger(hunger Hunger) Hunger {
	a.SetHunger(a.GetHunger() + hunger)
	return a.GetHunger();
}

func (a *Animal) DecreaseHunger(hunger Hunger) Hunger {
	return a.IncreaseHunger(-hunger);
}

// --------------

func (a *Animal) ApplyDamage(damage Health) Health {
	return a.DecreaseHealth(damage);
}

// --------------

func (a *Animal) Heal(max Health) Health {
	beforeHealth := a.GetHealth();
	a.IncreaseHealth(max);
	a.IncreaseHunger(a.GetHealth() - beforeHealth)
	return a.GetHealth();
}

func (a *Animal) IsDead() bool {
	return a.GetHealth() == 0
}

// ----------------

func (a *Animal) Print() {
	fmt.Printf("%#v\n", *a);
}

func NewAnimal() *Animal {
	animal := Animal{
		health: DefaultMaxHealth,
		MaxHealth: DefaultMaxHealth,
		hunger: 0,
		MaxHunger: DefaultMaxHunger,
	}

	return &animal;
}