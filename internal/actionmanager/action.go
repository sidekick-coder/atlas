package actionmanager 

type ActionContext map[string]any

type Action interface {
	Execute(ctx ActionContext) error
}
