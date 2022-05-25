package config

type Configer interface {
	Read() error
	validate() error
	Print()
}
