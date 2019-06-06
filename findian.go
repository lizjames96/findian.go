package main

import "fmt"
import "strings"

func main() {

  fmt.Printf("Welcome to Golang excercises\n")
  findian()
}

func findian() {
  fmt.Printf("\nThis is findian function to find letter i at the begining, n at the end and a in middle somewhere\n")
  fmt.Printf("\nEnter the string here: ")

  var fstr string
  fmt.Scan(&fstr)
  lString := strings.ToLower(fstr)    // tidy up string
  fstr = strings.TrimSpace(lString)

//  fmt.Printf("\nYou entered = %s\n", fstr)



  if strings.IndexAny(fstr, "i") != 0 {
    fmt.Printf("\nNot Found\n")
    return
  }
  if strings.LastIndex(fstr, "n") == -1 {
    fmt.Printf("\nNot Found\n")
    return
  }
  if strings.Contains(fstr, "a") == false {
    fmt.Printf("\nNot Found\n")
    return
  }

  fmt.Printf("\n Found \n")

}