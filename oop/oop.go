package main


import(
  "fmt"
)

func main() {

  // compositionOverInheritence()

  methodDemo()
}

type Person struct {
  Id, Age int
  Name string
}

type Player struct {
  Person
  Role string
  Rank int
}

func compositionOverInheritence() {

  p := Player{}
  p.Id = 1
  p.Age = 22
  p.Name = "Josh"
  p.Role = "Tank"
  p.Rank = 10

  fmt.Println("New person: ", p)

  p0 := Player{Person:Person{Id:1, Age:22, Name:"Josh"}, Role:"Tank", Rank:10}
  
  fmt.Println("Literal person: ", p0)

}

func (person *Person) personGreetings() string {
  return fmt.Sprintf("Hello from %s with id %d", person.Name, person.Id)
}

func methodDemo() {
  
  p0 := Player{Person:Person{Id:1, Age:22, Name:"Josh"}, Role:"Tank", Rank:10}

  fmt.Println(p0.personGreetings())
}

