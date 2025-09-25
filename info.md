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

