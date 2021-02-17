package client

//Client interact with 3-rd party joke API
type Client interface {
	//GetJoke returns one joke
	GetJoke() (string, error)
}
