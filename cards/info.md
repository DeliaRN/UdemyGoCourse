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
This will print: [72 105 32 116 104 101 114 101 33]

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





## Go is not OO
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





# PACKAGES

## ioutil
It implements some common operations for working with the underlying hard drive on our working machine.

