package model

type TypeWarCafe string

const (
	INVEST TypeWarCafe = "Invest"
	WAR    TypeWarCafe = "War"
)

type Model struct {
	ID   int64
	Type string

	TargetID int64
	Number   int
}
