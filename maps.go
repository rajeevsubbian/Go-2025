package main

import "fmt"

func main() {
	m := map[string]string{}
	m["name"] = "John Doe"
	m["age"] = "30"
	m["city"] = "New York"

	fmt.Println("Map contents:")
	for key, value := range m {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Checking if a key exists
	if _, exists := m["country"]; !exists {
		fmt.Println("Country key does not exist")
	}

	// Adding a new key-value pair
	m["country"] = "USA"

	// Updating an existing value
	m["age"] = "31"

	// Deleting a key-value pair
	delete(m, "city")

	fmt.Println("\nUpdated map contents:")
	for key, value := range m {
		fmt.Printf("%s: %s\n", key, value)
	}

}
