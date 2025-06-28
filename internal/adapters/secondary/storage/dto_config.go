package storage

// Config represents the main configuration structure for the game.
// It contains all the initial parameters, level definitions, item effects, and enemy configurations.
type Config struct {
	CharacterStartParams Character         `yaml:"character_start_params"`
	Levels               []Level           `yaml:"levels"`
	Elixir               ItemEffects       `yaml:"elixir"`
	Scroll               ItemEffects       `yaml:"scroll"`
	Food                 FoodEffects       `yaml:"food"`
	Weapon               WeaponEffects     `yaml:"weapon"`
	EnemyAgility         map[string][2]int `yaml:"enemy_agility"`
	EnemyStrength        map[string][2]int `yaml:"enemy_strength"`
	EnemyAnimosity       map[string][2]int `yaml:"enemy_animosity"`
	EnemyHealth          map[string][2]int `yaml:"enemy_health"`
	Enemies              map[string]Enemy  `yaml:"enemies"`
}

// Character defines the base attributes for a game character.
type Character struct {
	MaxHealth int `yaml:"max_health"` // Maximum health points
	Health    int `yaml:"health"`     // Current health points
	Strength  int `yaml:"strength"`   // Strength attribute
	Agility   int `yaml:"agility"`    // Agility attribute
}

// Level contains configuration for a specific game level or range of levels.
type Level struct {
	Range        [2]int         `yaml:"range"`         // Level range this configuration applies to [min, max]
	EnemyChances map[string]int `yaml:"enemy_chances"` // Probability weights for different enemies
	EnemyCount   [2]int         `yaml:"enemy_count"`   // Number of enemies [min, max]
	ItemsCount   [2]int         `yaml:"items_count"`   // Number of items [min, max]
	Treasure     [2]int         `yaml:"treasure"`      // Treasure amount range [min, max]
}

// ItemEffects defines the possible effects of consumable items.
type ItemEffects struct {
	MaxHealth []int    `yaml:"max_health"` // Possible max health modifications
	Strength  []int    `yaml:"strength"`   // Possible strength modifications
	Agility   []int    `yaml:"agility"`    // Possible agility modifications
	Name      []string `yaml:"name"`       // Possible item names
	Duration  []int    `yaml:"duration"`   // Effect durations in turns
}

// FoodEffects defines the effects of food items.
type FoodEffects struct {
	Health []int    `yaml:"health"` // Possible health restorations
	Name   []string `yaml:"name"`   // Possible food names
}

// Enemy defines a specific enemy instance with references to its attribute segments.
type Enemy struct {
	EnemyAgility   string `yaml:"enemy_agility"`   // Reference to agility segment
	EnemyStrength  string `yaml:"enemy_strength"`  // Reference to strength segment
	EnemyAnimosity string `yaml:"enemy_animosity"` // Reference to animosity segment
	EnemyHealth    string `yaml:"enemy_health"`    // Reference to health segment
}

// WeaponEffects defines the attributes of weapons.
type WeaponEffects struct {
	Strength []int    `yaml:"strength"` // Possible strength bonuses
	Name     []string `yaml:"name"`     // Possible weapon names
}

// EnemyNum is a map counting different enemy types.
// The key is the enemy type name, value is the count.
type EnemyNum map[string]int
