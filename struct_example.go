// structs are to create a TYPE of collection of fields (typed collection)
// So yeah, we create a new type that holds a collection  of fileds, a structure of data!

type person struct {
  name string
  age int
}


// we can create a struct and return it

func newPerson(name string) *person {
  p := person(name: name, age: 18)
  return &p
}

// use

func main() {
  fmt.Println(person{}) // name: "", age: 0 
  fmt.Println(person{age: 10}) // name: "", age: 10 
  fmt.Println(person{"a", 10}) // name: "a", age: 10
  fmt.Println(newPerson(name: "April") // &{April 18}


  salvia := person{name: "Salvia", age: 16}

  supernaturalSalvia := &salvia

  supernaturalSalvia.age = 17 // automatically denotes, but you can use (*p).x = 17 syntax
  fmt.Println(supernaturalSalvia) // &{Salvia, 17}

  adult := newPerson("Jimmy") // &{Jimmy, 18}

  dog := struct {
    breed string
    age int
  } {
    "Labrador",
    3
  }
  fmt.Println(dog) // {Labrador, 3}


  // at the top-level type is just a type, so you can omit from the elements of literals

  people := map[string]person{
    "wife" : { name: "Jinny", age: 26 },
    "husband" : { name: "Harry", age: 27 },
  }

}
