package main

import "fmt"

//here first we should create struct with map as a key
//2. have to create
type customSet struct {
	container map[string]struct{}
}

func makeSet() *customSet {
	return &customSet{container: make(map[string]struct{})}
}

func (c *customSet) add(key string) {
	c.container[key] = struct{}{}
}

func (c *customSet) remove(key string) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("key doesnt exists")
	}
	delete(c.container, key)
	return nil
}

func (c *customSet) ifExists(key string) bool {
	_, exists := c.container[key]
	return exists
}

func main() {
	customSet := makeSet()
	customSet.add("C")
	customSet.add("A")
	customSet.add("T")
	fmt.Println(customSet.ifExists("C"))
	customSet.remove("C")
	fmt.Println(customSet.ifExists("C"))
}
