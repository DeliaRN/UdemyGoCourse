# TIPS AND DOCUMENTATION ON GOLANG SYNTAX

### **Strings can be declared out of functinos**, but they will cause an error if you try to use the short declaration operator outside of a function.

```var stringOutOfFunction string -> This is valid```

```stringOutOfFunction2 := "This will cause an error"```


### **You can reasign values, but you cannot redeclare variables**

```var card string = "Ace of Spades"```

```card := "Ace of Spades"```

```card = "Different card"```
```card := "Ace of Spades"```

This last line would be a redeclaration of the variable ```card``` and will cause an error


## **Type is inferred**

If you write a function like this one

```
func newCard() string {
	return "Five of Diamonds"
}
```

when you use it: ```card := newCard()``` the type of ```card``` will be inferred to **string** since that's the returning type of ```NewCard()``` function.

## Type conversion
Type we want (value we have)
[]byte("Hi there!")
Float(3)

```
greeting := "Hi there!"
fmt.Printkln([]byte(greeting))
```

This will print:```[72 105 32 116 104 101 114 101 33]```

Deck -> []string -> string -> []byte


# **Slices and arrays**
Arrays are fixed length data structures.
Slices can shrink and grow.

You can fill them directly or by calling functions that return the type expeected for that Slice/Array. Here, cards1 expects only string type.

```cards1 := []string{newCard(), newCard()}```
```cards2 := []string{"Ace of Spades", newCard()}```
	
Since Slices can grow, you can use ```append``` to add new elements to an Slice.

```cards1 = append(cards1, "Six of Spades")```

This doesn't modify the original slice, it creates a new one and reassigns ```cards1``` to it


## Splitting arrays:

Taking the following slice:
```fruits["apple", "banana", "grape", "orange"]```

we use the following pattern:
```slice[startIndexIncluding : upToNotIncluding]```

```fruits[0:2]``` -> ["apple", "banana"]
if you do not put the left side of the collon ```:```, Go will understand it's from the beginning. If it's the right side, Go will understand it's until the end.
- fruits[:3] -> ["apple", "banana", "grape"]
- fruits[3:] -> ["orange"]

## Creating slices
When you create an slice, Go will automatically create an array and a structure that records the length of the slice, the capacity of the slice, and a reference to0 the underlying array.
So, slices are reference type, since they contain a reference to the underlying list of records.

## Byteslice

"Hi there!" string

[72 105 32 116 104 101 114 101 33] byte slice


In **asciitable.com**, in the column decimal, you can see what decimal number corresponds to each character.

So, a byteslice is a computer-friendly way to represent a string.


# Go is not OO
Therefore, we need to "extend" a base type and add some extra functionality to it so it behaves as classes behave in OO languages.

- Java: ```Class Deck```
- Go: ```type deck []string```

In Go, we say "We want an array of strings and add a bunch of functions specifically made to work with that array".

For this, we make **functions with a receiver**. They are like methods, they are functions that belong to an instance.


## Receivers
Any variable of type deck now has access to all the methods we define on it

```
func (d deck) print() {
	for i, card := range d {
		frm.Println(i, card)
	}
}
```

- The ```deck``` is a reference to the type that we want to attach this method to. So any variable of type deck has access to ```print()```

- The ```d``` is the actual "instance" of deck type we are working on. ```cards.print()``` (It's not really an instance because it's not OO, it's actually a "copy")


We do not have ```this``` or ```self``` like in Java, Python or Ruby. 

We do have the convention of using the first letter of the type for naming the receiver. ```deck``` -> ```d```


For example, in the code:

```
func (ls laptopSize) getSizeOfLaptop() {
    return ls
}
```
The variable ```ls``` represents a value of type ```laptopSize```


## Returning double values

To assign different values returned at the same time in different variables, we have the following syntax:

```hand, remainingDeck = deal(cards, 5)```
```hand, remainingDeck := <deck>, <deck>```
First result to first var, second to second.


byteSlice, err := os.ReadFile(filename)
If nothing goes wrong, 'error' will have a value of 'nil'

# FUNCTIONS
A function allows you to group code into a reusable unit.

func keyword + name of the function + comma-separated list of zero or more parameters and types in round brackets.

## Function Parameters
All parameters must be explicitly typed; there is no type inference for parameters.

There are no default values for parameters so all function parameters are required.

- No parameters:
```
func PrintHello() {
    fmt.Println("Hello")
}
```

- One parameter:
```
func PrintHelloName(name string) {
  fmt.Println("Hello " + name)
}
```
Parameters of the same type can be declared together, followed by a single type declaration.

```
func PrintGreetingName(greeting, name string) {
  fmt.Println(greeting + " " + name)
}
```

## Parameters vs. Arguments

- **Function parameters** are the names defined in the function's signature,
such as `greeting` and `name` in the function `PrintGreetingName` above.

- **Function arguments** are the concrete values passed to the function parameters when we invoke the function.
`"Hello"` and `"Katrina"` are the arguments passed to the `greeting` and `name` parameters:

```
PrintGreetingName("Hello", "Katrina")
```


## Return Values

Function parameters can be followed by return values that must also be explicitly typed.

- Single return values:
```
func Hello(name string) string {
  return "Hello " + name
}
```

- Multiple return values: 

```
func HelloAndGoodbye(name string) (string, string) {
  return "Hello " + name, "Goodbye " + name
}
```

Values are returned using `return` keyword at the end of the function.
There can be multiple `return` statements in a function.
The function ends as it hits one of those `return` statements


## Invoking Functions
Invoking a function is done by specifying the function name and passing arguments for each of the function's parameters in parenthesis.


- **No parameters**, no return value:
```
func PrintHello() {
    fmt.Println("Hello")
}
```
Called like this:	`PrintHello()`


- **One parameter**, one return value:
```
func Hello(name string) string {
  return "Hello " + name
}
```
Called like this: `greeting := Hello("Dave")`

- **Multiple parameters**, multiple return values:
```
func SumAndMultiply(a, b int) (int, int) {
    return a+b, a*b
}
```
Called like this: `aplusb, atimesb := SumAndMultiply(a, b)`



### Named Return Values and Naked Return
As well as parameters, return values can optionally be named.
If named return values are used, a return statement without arguments
will return those values. This is known as a **'naked' return**.

```
func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
    sum, mult = a+b, a*b
    sum -= c
    mult -= c
    return
}
```



# TESTING

Go testing is not RSpec, selenium, mocha, jasmine...

- To make a test we create a new file ending in ```_test.go``` 

- To run tests in a package, we use ```go test```

Writing tests in go is actually done by just writting simple go code that tests the desired file.


## How to know what to test

Deciding what to test is part of the logic behind testing.
It's about deciding what do we care about as developers:

With the example of the deck, it makes sense to test that...
- ... for ```func TestNewDeck``` ...
	- ... the deck has 52 cards
	- ... the first one is an Ace and last is a King.
	- ... it's not empty
- ... for ```func TestSaveToFile``` ...
	- ... the file is created
	-
- ... for ```func TestNewDeckFromFile``` ...
	- ... the deck is there
	- ... the deck is lodaed

etc.

Then for ```func TestNewDeck```, we should create a new test and write an if statement that checks the number of cards.


## Naming conventions

- Why is it ```func newDeck``` with regular 'n', and the test is ```func TestNewDeck``` with capital 'n'?

- What is ```t *testing.T```in funcion declaration?


## Running them

```go test``` : Almost no Feedback in case they pass.
```
go test
PASS
ok  	cards	0.776s
```



```go test -v``` : Adding '-v' flag to the command provides a more detailed (verbose) feedback.
```
go test -v
=== RUN   TestNewDeck
--- PASS: TestNewDeck (0.00s)
PASS
ok  	cards	0.384s
```


# STRUCTS

Structs are collections of fields. They are useful to group data together to form records. They are similar to classes in other languages, but they don't have methods.

```
type person struct {
	firstName string
	lastName  string
}
```

There are several approaches on how to use them and declare them:

Standard:
```
alex := person{
	firstName: "Alex",
	lastName:  "Anderson",
}
```

Short one:

```jack := person{"Jack", "Johnson"}```
We **shouldn't use Jack's approach**, since it depends on the order of fields in the struct.
With Alex, the order doesn't matter because we are using field names. It's ok as long as we use all fields.

```fmt.Println(alex)```// {Alex Anderson}

With Emily, we can create empty srtucts that will be filled with zero values for all fields
```var emily person```

```fmt.Println(emily)``` // {  } - both fields are empty strings

We can also use ```Printf``` and ```%+v``` verb to show field names, so we can have more info:

```fmt.Printf("%+v", emily)``` // {firstName: lastName:} 

We can also asign values to the files of an struct just by calling them like in some many other languages:
```
emily.firstName = "Emily"
emily.lastName = "Evans"
```

```fmt.Printf("%+v", emily)``` // {firstName:Emily lastName:Evans}


## Embedded structs

Structs can be embedded inside one another.
```
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
```

```
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jp@email.com",
			zipCode: 12345,
		},
	}
```

If we print this with the %+v verb, as in Emily's sample, we'll get:
```{firstName:Jim lastName:Party contact:{email:jp@email.com zipCode:12345}}```

## Anonymous fields

An anonymous field is a field without a name, only a type. We can use this to embed a struct within another struct. This way we can access the fields of the embedded struct directly from the outer struct:

```
type person struct {
	firstName string
	lastName  string
	contactInfo		
}
```
Here, ```contactInfo``` is an anonymous field. We can access the fields of contactInfo directly from person

This, however, is not recommended because it can lead to confusion.

It would be used as: 
```
jim := person{
	firstName: "Jim",
	lastName:  "Party",
	contactInfo: contactInfo{
		email:   "email.com",
		zipCode: 12345,
	},
}
```

## Receivers and structs

This function takes a person as a receiver:
```
func (p person) print() {
	fmt.Printf("%+v", p)
}
```

This way, we can call ```jim.print()``` instead of doing ```fmt.Printf("%+v, jim")``` and the result will be the same.

If we do not really use the value (`p`or `eb`in these examples)...
```
func (eb englishBot) getGreeting() string {
	// VERY custom logic for English blablabla
	return "Hi!"
}
```
we can take it out!
```
func (englishBot) getGreeting() string {
	// VERY custom logic for English blablabla
	return "Hi!"
}
```


# Pointers

```
jim.print()
jim.updateName("Jimmy")
jim.print()

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
```

This is not going to work because we are working with a copy of the value.

When we do ```jim:=person{}``` , it gets located on RAM memory address, eg 0001, with the whole value of ```jim```-.

Go is a **pass by value** language. Whenever you pass a value into a function, Go will take it and copy all of that data inside it, and place it in a new container/address inside the RAM.

When we run ```jim.update("Jimmy")```, ```jim``` will still exist in slot 0001, and its copy called ```p``` will be located on another slot, eg 0003. So now we have ```jim``` on 0001, and ```p``` on 0003. And it's ```p``` the one being updated.


If we want to modify the actual value we need to use **a pointer**.

A pointer holds the memory address of a value.

```
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
```

- We can get the memory address of a value using the ```&``` operator

```&variable```: Give me the memory address og the value this variable is pointing at

- We can dereference a pointer using the ```*``` operator.

```*pointer``` : Give me the value this memory address is pointing at


A pointer is a type that "points to" the memory address of a value of a specific type.

Like this, we get the variable ```jim```'s address using the address ```&``` operator.
Then, we use the derefence ```*```operator, to access the value stored at the memory address that a pointer points to.

We can say it "follows" the pointer to the value it references.

```
jimPointer := &jim
jimPointer.updateName("Jimmy")
jim.print()
```
'```jimPointer := &jim```': Give me access to the memory address of jim variable and asign it to jimPointer.

If we printed jimPointer, we would see the memory address of jim, not the values it holds.


```
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
```

When we do this above, in the receiver we use '```*person```', which is a type description. It means we are working with a pointer to a person.

When, inside the method, we use '```*pointerToPerson```', this is an operator that means we want to manipulate the value the pointer is referencing. It's like saying _**I don't want to know the address anymore, instead give me direct address to what there is inside the memory address**_.

We put it inside parenthesis to make it an actual person value that is sitting in memory: ```(*pointerToPerson).firstName```

	address	|  value
	0001    | person{firstName:"Ji...}

Turn address into value with *address

Turn value into address with &value



## Pointer shortcut
If we have a function like ```func (pointerToPerson *person) updateName(newFirstName string)```, Go will let us use it both with a ```person```and with a ```*person```

This way, we can do: 
```
jimPointer.updateName("Jimmy")
jim.print()
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
```
**_Notice we've erased ```pointerToJim = &jim```, and it still works_**. That's because Go will understand that the function will work perfectly with both, and it will work with it as a pointer automatically.


## Gotchas with Pointers

With slices, it does not apply the address and pointer situation.

If we have:
```
func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
```

## Reference Trype VS Value Trype

### Value types:
Use pointers to change these things in a function
- int
- float
- string
- bool
- structs

### Reference types:
Do not worry aboutv pointers with these
- slices
- maps
- channels
- pointers
- function


# Maps

Collections of key value pairs.

Both keys and values are statically typed, so all keys must be all the same type.
In the same way, values of a map must be all the same type too.

```
color := map[string]string {
		"Red":   "#ff0000",
		"Green": "#4bf745",
	}
```

- We can declare it and not initialize it:

`var colorMap map[string]string`
 but, if we print it, it will be initialized by go with a zero value. 

- We can also create it with ```make```:

`colorMap := make(map[string]string)`

- We can **add** key-value pairs just by declaring them:

`color["White"] = #ffffff`

Now, if we print `color`, it will be: `map[green:#4bf745 red:#ff0000 white:#ffffff]`


- We can **delete** key-value pairs with `delete`function,
passing the name of the map and the key we want to delete.

`delete(colors, "White")`

## Iterating over maps
```
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
```

Each run of this function, will print the map 
in a different order. That's because, in Go,
 maps are unordered.


White #ffffff
Red #ff0000
Green #4bf745


Red #ff0000
Green #4bf745
White #ffffff

In the same way, maps do not allow repeated 
elements. If we tried to add a key that already
 exists, it will be marked as an error by the IDE and won't compile 


## Maps VS Structs

In maps keys must share type, and values must share type.
In structs, types can be of different type.

In maps, keys are indexed, and we can iterate them. 
In structs, keys do not support indexing.

In maps, you don't need to know all the keys at compile time.
In structs, you do.

Map is a reference type
Structs are value type.

Maps are used to represent a collection of related properties.
Structs are sued to represent a "thing" with a lot of different properties.


# INTERFACES

We knoe values have types and functions must specifies the type of the args

Does that mean EVERY function will be rewritten to accomodate different types even if they share the logic?
Nope :) 

Imagine the `func (d deck) shuffle()`. It will be also `func (s []int) shuffle()` , and so on.

That's one of the problems interfaces try to solve.

If we do this: 
```
type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}
```

We just need to have a getGreeting with both bots as receivers:
```
func (englishBot) getGreeting() string {
	//ommit the value 'eb' since it's not being used
	// VERY custom logic for English blablabla
	return "Hi!"
}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic for Spanish blablabla
	return "Hola!"
}
```


From the moment both spanishBot and englishBot have a function
with them as receivers, and the bot interface declares that
same function -> Both conform to the interface and then turn themselves
into honorable members of type `bot`too.



# PACKAGES

## ioutil or os
It implements some common operations for working with the underlying hard drive on our working machine.

ioutil is deprecated since 1.16. Now we use 'os' instead, which provided platform-independent interface to operating system functionality.

os.Exit(0) -> 0 code indicates success.
os.Exit(1) -> Non-zero indicates an error.

```func Remove(name string) error```
Removes the named file or directory. If there's an error, it will be of type ```*PathError```


## strings
It implements operations to work with UTF-8 encoded strings


## math
### math/rand
It implements pseudo-random number generators
```rand.Intn(n int) int```
It use a Seed as a 'source' in which is based the randomization.
```func NewSource (seed int64) Source```
```func (r *Rand) Intn(n int) int```

## time
func (t Time) UnixNano() int64 : It returns t as a Unix time, so everythime the application

