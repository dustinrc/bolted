package bolted

type Configurer interface {
	Get(directive string) (value string, ok bool)
}
