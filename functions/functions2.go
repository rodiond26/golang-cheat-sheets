package main

import "fmt"

func main() {
	i := new(int32)
	fmt.Printf("i = %T : %v\n", i, i) // *int32 0xc00001a098

	client := Client{}
	// применение функции new при создании поля структуры
	client.IMG = new(Avatar)
	// применение ссылки при создании поля структуры
	client.IMG = &Avatar{}
	fmt.Printf("client = %T : %#v\n", client, client)

	updateAvatar(client)
	fmt.Printf("client = %T : %#v\n", client, client)

	updateClient(client)
	fmt.Printf("client = %T : %#v\n", client, client)

	b := client.HasAvatar()
	fmt.Printf("b = %v\n", b)

	client.UpdateAvatar2()
	fmt.Printf("Avatar = %T : %#v\n", client.IMG, client.IMG)
}

type Client struct {
	ID   int64
	Name string
	Age  int
	IMG  *Avatar
}

type Avatar struct {
	URL  string
	Size int64
}

func updateAvatar(client Client) {
	client.IMG.URL = "updated_url"
}

func updateClient(client Client) {
	client.Name = "new_name"
}

func (c *Client) HasAvatar() bool {
	if c.IMG != nil && c.IMG.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar2() {
	c.IMG.URL = "new_url"
}
