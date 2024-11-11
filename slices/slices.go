package slices

import "fmt"

type Product struct{
	Id int
	Title string
	Price float64
}

func NewProduct (id int, title string, price float64) Product{
	return Product{Id: id, Title: title, Price: price}
}

func main() {
	hobbiesArray := createAnArray()
	fmt.Println("hobbies array: ", hobbiesArray)

	
	printArray(hobbiesArray)


	firstSlice := hobbiesArray[:2]
	fmt.Println("first slice: ", firstSlice)
	secondSlice := hobbiesArray[0:2]
	fmt.Println("second slice: ", secondSlice)


	firstSlice = firstSlice[1:]
	fmt.Println("first slice: ", firstSlice)
	firstSlice = append(firstSlice, hobbiesArray[len(hobbiesArray)-1])
	fmt.Println("first slice: ", firstSlice)

	dynamicArray := []string {"success in career", "health"}
	fmt.Println("dynamicArray: ", dynamicArray)
	dynamicArray[1] = "travelling the world"
	fmt.Println("dynamicArray: ", dynamicArray)
	dynamicArray = append(dynamicArray, "happiness")
	fmt.Println("dynamicArray: ", dynamicArray)

	pastry := NewProduct(1, "Pastry", 1.99)
	tea := NewProduct(2, "Tea", 0.50)
	milk:= NewProduct(3, "Milk", 0.35)

	newDynamicArray := []Product{pastry, tea}
	fmt.Println("newDynamicArray: ", newDynamicArray)
	newDynamicArray = append(newDynamicArray, milk)
	fmt.Println("newDynamicArray: ", newDynamicArray)

}
func createAnArray() []string {
	return []string{"reading", "watching TV", "listening music"}
}


func printArray(arr []string) {
	if len(arr) < 1 {
		return
	}
	fmt.Printf("First element of the array is: %s \n", arr[0])
	if len(arr) > 1 {
		fmt.Printf("Second and the third elements of the array are: %s \n", arr[1:3])
	}
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
