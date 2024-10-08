package game

import "time"

type Settings struct {
	// time until inputs are evaluated and game state is updated
	TurnDuration time.Duration `json:"turnDuration"`

	// size of the map
	GridSize int `json:"gridSize"`

	// how many bombs can be stored at once
	InventorySize int `json:"inventorySize"`

	// how many turns it takes before a player can get another bomb
	BombRespawnTime int `json:"bombRespawnTime"`

	// how many lives the player has 
	StartHealth int `json:"startHealth"`

	// how many turns until the airstrike detonates
	AirstrikeFuseLength int `json:"airstrikeFuseLength"`

	// how many turns until a player's dropped bomb detonates
	BombFuseLength int `json:"bombFuseLength"`

	// how many turns the corpose of a ship remains on the map
	DeathTime int `json:"deathTime"`

	// don't allow joining of new players
	Locked bool `json:"locked"`

	// don't execute the game but allow joining of new players
	Paused bool `json:"paused"`
}

