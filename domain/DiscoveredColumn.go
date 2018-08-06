package domain

type DiscoveredColumn struct{
	Field string
	Type string
	Null bool
	Key string
	Default string
	Extra string
	MaxCharacterLimit int
	MaxDigits int
	MaxDecimals int
	Enums []string
}