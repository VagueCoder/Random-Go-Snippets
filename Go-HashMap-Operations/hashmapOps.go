package hashmapOps

import (
	"fmt"
	"net/http"
	"os"

	"github.com/VagueCoder/Random-Go-Snippets/Go-HashMap-Operations/data"
	"github.com/VagueCoder/Random-Go-Snippets/Go-HashMap-Operations/hashmap"
)

// Person is a random struct type for demonstration purposes
type Person struct {
	Name  string
	Age   int
	Phone string
}

func RunHashMapOperationsShort() {
	// initialize HashMap
	fmt.Println("Initialized HashMap:")
	h := hashmap.NewHashMap()
	fmt.Println("Put: (MyKey, MyVal)")
	h.Put("MyKey", "MyVal")
	fmt.Println("PutIfAbsent: (MyKey2, MyValInserted)")
	h.PutIfAbsent("MyKey2", "MyValInserted")
	fmt.Printf("Print HashMap: ")
	h.Print()
}
func RunHashMapOperations() {
	var key, val, Square, Default interface{}

	// initialize HashMap
	fmt.Println("Initialized HashMap:")
	h := hashmap.NewHashMap()
	h.Print()

	// initialize Item
	fmt.Println("--------------------------------------------\nInitialized Item:")
	item := hashmap.NewItem("MyKey", "MyVal")
	item.Print()

	// Inserting different and random types of variables and values into hashMap
	fmt.Println("--------------------------------------------\nInserting Key-Value Pairs:")
	h.Put("MyKey", "MyVal")
	h.Put("HTTP-Request", &http.Request{})
	h.Put(os.File{}, 8.8)
	h.Put(struct{}{}, 6)
	h.Put(true, make(map[string]string))
	h.Put(&Person{}, &data.DataType{ID: "ID-1", Value: 1})
	h.Put(hashmap.NewHashMap(), h)
	h.Print()

	// Putting value if key is absent
	fmt.Println("--------------------------------------------\nInserting Key-Value Pairs if absent:")
	h.PutIfAbsent("MyKey", "MyValUpdated")   // Shouldn't update
	h.PutIfAbsent("MyKey2", "MyValInserted") // Should insert
	h.Print()

	// Merging HashMaps
	fmt.Println("--------------------------------------------\nMerging hashMaps:")
	h2 := hashmap.NewHashMap()
	h2.Put("MyKey", "MyVal")                  // Updates value
	h2.PutIfAbsent("MyKey3", "MyValInserted") // Inserts value
	h.Merge(h2)
	h.Print()

	// Merging HashMaps those keys which are missing in original HashMap
	fmt.Println("--------------------------------------------\nerging HashMaps those keys which are missing in original HashMap:")
	h2.Put("MyKey", "---")          // Ignored in merge
	h2.PutIfAbsent("MyKey4", "---") // Inserted while merging
	h.MergeMissing(h2)
	h.Print()

	// Clearing Second HashMap
	fmt.Println("--------------------------------------------\nClearing Second HashMap:")
	h2.Print()
	fmt.Println("Clearing ...")
	h2.Clear()
	h2.Print()

	// Checking if HashMao is empty
	fmt.Println("--------------------------------------------\nChecking if HashMao is empty:")
	fmt.Printf("Is Second HashMap empty? - %v\n", h2.IsEmpty())

	// Iterating over items
	fmt.Println("--------------------------------------------\nIterating over items:")
	ch := h.Items()
	for item := range ch {
		item.Print()
	}

	// List of Keys
	fmt.Println("--------------------------------------------\nList of Keys:")
	fmt.Printf("%+v\n", h.Keys())

	// List of Values
	fmt.Println("--------------------------------------------\nList of Values:")
	fmt.Printf("%+v\n", h.Values())

	// Cloning Original HashMap and assigning the clone to second HashMap
	fmt.Println("--------------------------------------------\nCloning HashMap:")
	h2.Clear()
	h2.Print()
	fmt.Println("\nCloning ...")
	h2 = h.Clone()
	h2.Print()

	// Size of HashMap
	fmt.Println("--------------------------------------------\nSize of HashMap:")
	fmt.Printf("Size of HashMap = %d\n", h.Size())

	// Getting Value for Key
	fmt.Println("--------------------------------------------\nGetting Value for Key:")
	key = "MyKey"
	fmt.Printf("Value for Key %v (%T) = %v (%T)\n", key, key, h.Get(key), h.Get(key))
	key = os.File{}
	fmt.Printf("Value for Key %v (%T) = %v (%T)\n", key, key, h.Get(key), h.Get(key))

	// Getting Value for Key, default value if key not found
	fmt.Println("--------------------------------------------\nGetting Value for Key, default value if key not found:")
	key = "MyKey"
	err := fmt.Errorf("Key-Not-Found")
	fmt.Printf("Value for Key %v (%T) = %v (%T)\n", key, key, h.GetOrDefault(key, err), h.GetOrDefault(key, err))
	key = os.LinkError{}
	fmt.Printf("Value for Key %v (%T) = %v (%T)\n", key, key, h.GetOrDefault(key, err), h.GetOrDefault(key, err))

	// Replacing value
	fmt.Println("--------------------------------------------\nReplacing value:")
	h.Replace("MyKey5", "MyValUpdated")     // Shouldn't update
	h.Replace("MyKey2", "MyKey2ValUpdated") // Should update
	h.Print()

	// Replacing all values
	fmt.Println("--------------------------------------------\nReplacing all values:")
	h2.Clear()
	h2.Put("MyKey2", "MyVal")
	h2.Print()
	fmt.Println("Replacing All Items ...\n")
	h2.ReplaceAll(h)
	h2.Print()

	// Checking Equality
	fmt.Println("--------------------------------------------\nChecking Equality:")
	fmt.Printf("Checking before assigning. Equal: %v\n", h.Equals(h2)) // Shouldn't be equal
	h2.Merge(h)
	fmt.Printf("Checking after assigning. Equal: %v\n", h.Equals(h2)) // Should be equal

	// Check if contains key
	fmt.Println("--------------------------------------------\nCheck if contains key:")
	key = "MyKey"
	fmt.Printf("HashMap contains key '%v (%T)' : %v\n", key, key, h.ContainsKey(key))
	key = os.LinkError{}
	fmt.Printf("HashMap contains key '%v (%T)': %v\n", key, key, h.ContainsKey(key))

	// Check if contains Value
	fmt.Println("--------------------------------------------\nCheck if contains Value:")
	val = http.Client{}
	fmt.Printf("HashMap contains value '%v (%T)' : %v\n", val, val, h.ContainsValue(val))
	val = 8.8
	fmt.Printf("HashMap contains value '%v (%T)': %v\n", val, val, h.ContainsValue(val))

	// Removing
	fmt.Println("--------------------------------------------\nRemoving:")
	h2.Clear()
	h2.Put("MyVal", 100)
	h2.Print()
	fmt.Println("Removing key ...")
	h2.Remove(hashmap.NewItem("MyVal", 100))
	h2.Print()

	// For Each Value
	fmt.Println("--------------------------------------------\nFor Each Value:")
	h2.Clear()
	h2.Put("k1", 1)
	h2.Put("k2", 2)
	h2.Put("k3", 3)
	h2.Print()
	Square = func(v interface{}) { fmt.Printf("Square of %d = %d\n", v.(int), v.(int)*v.(int)) }
	h2.ForEach(Square.(func(interface{})))

	// Compute
	fmt.Println("--------------------------------------------\nCompute:")
	h2.Print()
	fmt.Println("Computing with func Square()...")
	Square = func(v interface{}) (interface{}, error) { return v.(int) * v.(int), nil }
	h2.Compute(Square.(func(v interface{}) (interface{}, error)))
	h2.Print()

	// Compute If Absent
	fmt.Println("--------------------------------------------\nCompute If Absent:")
	h2.Print()
	fmt.Println("Computing with func Default() when key is absent...")
	Default = func() interface{} { return 100 }
	h2.ComputeIfAbsent("k1", Default.(func() interface{})) // Shouldn't Update
	h2.ComputeIfAbsent("k4", Default.(func() interface{})) // Should Update
	h2.Print()

	// Compute If Present
	fmt.Println("--------------------------------------------\nCompute If Present:")
	h2.Print()
	fmt.Println("Computing with func Default() when key is present...")
	Default = func() interface{} { return 200 }
	h2.ComputeIfPresent("k1", Default.(func() interface{})) // Should Update
	h2.ComputeIfPresent("k5", Default.(func() interface{})) // Shouldn't Update
	h2.Print()
}
