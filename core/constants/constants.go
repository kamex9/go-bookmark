package constants

type StoreMode uint8

const (
	_ StoreMode = iota
	MEMORY
	FILE
	DATABASE
)
