package redis

type Client struct{}

// will create redis client

func NewClient() *Client {
	// get the config and assign arangodb client
	return &Client{}
}

func (c *Client) Exists(key string) (bool, error) {
	return false, nil
}

func (c *Client) Get(key string, val interface{}) error {
	return nil
}

func (c *Client) Set(key string, val interface{}) error {
	return nil
}
func (c *Client) Del(key string) error {
	return nil
}
