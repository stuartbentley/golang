package main

import "github.com/naegelejd/brewerydb"
import "fmt"

func main() {
    client := brewerydb.NewClient("505860a4bd7ebd07d586849f189f16ed")
   // Get any random beer
   beer, _ := client.Beer.GetRandom(&brewerydb.RandomBeerRequest{ABV: "8"})
   fmt.Println(beer.Name, beer.Style.Name)
}
