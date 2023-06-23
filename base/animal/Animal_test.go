package animal

import "testing"


func TestAnimalCreation(t *testing.T){
	animal := NewAnimal();

	if animal.GetHealth() != animal.MaxHealth {
		animal.Print();
		t.Fatalf("Initial Health is not equal to Max Health: [health]:%v [MaxHealth]:%v", animal.GetHealth(), animal.MaxHealth);
	}
	
	if animal.GetHunger() != 0 {
		animal.Print();
		t.Fatalf("Initial Hunger is not equal to ZERO: [hunger]:%v", animal.GetHunger());
	}
}


func TestSetHealth(t *testing.T){
	animal := NewAnimal();

	animal.SetHealth(90);

	if animal.GetHealth() != 90 {
		animal.Print();
		t.Fatalf("Failed to Set Health: [health]:%v", animal.GetHealth());
	}
}

func TestNegativeSetHealth(t *testing.T){
	animal := NewAnimal();
	
	animal.SetHealth(-10);
	
	if animal.GetHealth() < 0 {
		animal.Print();
		t.Fatalf("Found Negative Health: [health]:%v", animal.GetHealth());	
	}
}

func TestSetMoreHealthThanMaxHealth(t *testing.T){
	animal := NewAnimal();
	
	animal.SetHealth(animal.MaxHealth + 10);
	if animal.GetHealth() > animal.MaxHealth {
		animal.Print();
		t.Fatalf("Found Health more than MaxHealth: [health]:%v [MaxHealth]:%v", animal.GetHealth(), animal.MaxHealth);	
	}
}

func TestApplyDamage(t *testing.T){
	animal := NewAnimal();
	
	animal.ApplyDamage(10);

	if animal.GetHealth() != animal.MaxHealth - 10 {
		animal.Print();
		t.Fatalf("[health]:%v [expected]:%v", animal.GetHealth(), animal.MaxHealth - 10)
	}

}

func TestApplyMoreDamageThanMaxHealth(t *testing.T){
	animal := NewAnimal();
	
	animal.ApplyDamage(animal.MaxHealth + 10);
	
	if animal.GetHealth() != 0 {
		animal.Print()
		t.Fatalf("[health]:%v [expected]:%v", animal.GetHealth(), 0);
	} 

	if animal.IsDead() != true {
		animal.Print();
		t.Fatalf("[IsDead]:%v [expected]:%v", animal.IsDead(), true);
	}
}

func TestHeal(t *testing.T){
	animal := NewAnimal();

	animal.ApplyDamage(10);
	animal.Heal(5)

	if animal.GetHealth() != 95 {
		animal.Print();
		t.Fatalf("[health]:%v [expected]:%v", animal.GetHealth(), 95)
	}

	if animal.GetHunger() != 5 {
		animal.Print();
		t.Fatalf("[hunger]:%v [expected]:%v", animal.GetHunger(), 5)
	}

}

func TestHealMoreThanMaximumHealth(t *testing.T){
	animal := NewAnimal();

	animal.ApplyDamage(10);
	animal.Heal(20);

	if animal.GetHealth() != animal.MaxHealth {
		animal.Print();
		t.Fatalf("[health]:%v [expected]:%v", animal.GetHealth(), animal.MaxHealth)
	}

	if animal.GetHunger() != 10 {
		animal.Print();
		t.Fatalf("[hunger]:%v [expected]:%v", animal.GetHunger(), 10)
	}

}