package types

type ITemplater interface {
	Serve(port string) error
	Build() error
}