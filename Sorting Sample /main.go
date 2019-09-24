package main

//Importing packages to use inside of functions
import (
	"fmt"
	"sort"
)

//creating struct for coffee
type coffee struct {
	Name   string
	Rating float64
	Price  float64
}

//Sorting functions for rating
type ByRating []coffee

func (rate ByRating) Len() int           { return len(rate) }
func (rate ByRating) Swap(i, j int)      { rate[i], rate[j] = rate[j], rate[i] }
func (rate ByRating) Less(i, j int) bool { return rate[i].Rating > rate[j].Rating }

//Sorting functions for first letter in the name
type ByName []coffee

func (sem ByName) Len() int           { return len(sem) }
func (sem ByName) Swap(i, j int)      { sem[i], sem[j] = sem[j], sem[i] }
func (sem ByName) Less(i, j int) bool { return sem[i].Name < sem[j].Name }

//Sorting functions for price
type ByPrice []coffee

func (pic ByPrice) Len() int           { return len(pic) }
func (pic ByPrice) Swap(i, j int)      { pic[i], pic[j] = pic[j], pic[i] }
func (pic ByPrice) Less(i, j int) bool { return pic[i].Price > pic[j].Price }

//Main function
func main() {

	//List of coffee brands with my raing and price
	coffees := []coffee{
		{"Twin Engine", 5.0, 19.95},
		{"Folgers", 2.1, 5.99},
		{"Dunken Dounuts", 5.5, 18.99},
		{"Starbucks", 2.5, 22.99},
	}
	//Using sort to sort by rating
	fmt.Println("~~Sorted by Ryan's rating~~")
	sort.Sort(ByRating(coffees))
	for _, c := range coffees {
		fmt.Println(c.Rating, c.Name)
	}
	//Using sort to sort by first letter in the name
	fmt.Println("---------------")
	fmt.Println("~~Sorted by Name~~")
	sort.Sort(ByName(coffees))
	for _, u := range coffees {
		fmt.Println(u.Name, u.Rating)
	}
	//Using sort to sort by price
	fmt.Println("---------------")
	fmt.Println("~~Sorted by Price~~")
	sort.Sort(ByPrice(coffees))
	for _, q := range coffees {
		fmt.Println(q.Price, q.Name, q.Rating)
	}
}
