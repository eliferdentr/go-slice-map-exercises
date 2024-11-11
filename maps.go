package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := map[string]string{
	// 	"Google":              "www.google.com",
	// 	"Amazon Web Services": "www.aws.com",
	// }

	// fmt.Println(websites)
	// fmt.Println(websites["Amazon Web Services"])
	// websites["Linkedin"] = "linkedin.com"
	// fmt.Println(websites["Linkedin"])

	// delete(websites,"Google")
	// fmt.Println(websites)

	///
	// userNames := make([]string, 2)

	// userNames = append(userNames, "Elif")
	// userNames = append(userNames, "Elain")
	// fmt.Println(userNames)

	userNames := make([]string, 2, 5)
	userNames[0] = "Rhysand"
	userNames = append(userNames, "Elif")
	userNames = append(userNames, "Elain")

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 6.0
	courseRatings["java"] = 5.0
	courseRatings.output()

	for index, value := range userNames {
		fmt.Printf ("index: %d . value : %s\n", index, value)
	}

	for key, value := range courseRatings {
		fmt.Printf ("key: %s . value : %2.f\n", key, value)
	} 

}