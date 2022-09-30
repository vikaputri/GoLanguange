package main

func main() {
	rs := []interface{}{1, "Airell", true, "Ananda", true}

	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
