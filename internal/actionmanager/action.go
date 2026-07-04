package actionmanager 

type Action interface {
	Execute(params []string) error
}
