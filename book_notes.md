# The Go Programming Language

## Table of Contents
- [The Go Programming Language](#the-go-programming-language)
	- [Table of Contents](#table-of-contents)
	- [Preface](#preface)
	- [Chapter 1 - Tutorial](#chapter-1---tutorial)
	- [Chapter 2 - Program Structure](#chapter-2---program-structure)
	- [Chapter 3 - Basic Data Types](#chapter-3---basic-data-types)
	- [Chapter 4 - Composite Types](#chapter-4---composite-types)

## Preface

PDF version of book

1. Set environment variable **GOPATH** 

    ```go
    export GOPATH=$HOME/gobook
    ```
                
2. Fetch, build, install hello world app

    ```go
    go get gopl.io/ch1/helloworld
    ```
                
3. Run the app

    ```go
    $GOPATH/bin/helloworld
    ```

**NOTE**: the above actually fetched all of the source files for the book into `$HOME/gobook/src` creating both `/src` and `/bin` directories.


Some useful links:
- GoLang documentation: https://golang.org
- Go Blog: https://blog.golang.org
- Package documentations: https://golang.org/pkg
- Playground where you can try out code: https://play.golang.org

Note that the playground provides a persistent link, so it's possible to use it to store code snippets or share code within documents. However only the standard libraries can be used and some functions, such as networking, are restriced.


## Chapter 1 - Tutorial

1. Command line arguments available via `os` package
   - `os.Args` is a slice of strings
   - `os.Args[0]` is the command name
   - `strings.Join(os.Args[1:], " ")` would produce a string of all arguments
2. See solutions to PDF page 8 exercises 1.1 through 1.3 on command line arguments in folder [ex1](./ex1).
3. To create  a map with a string key and int values: `counts := make(map[string]int)`.
4. To then read input file, counting occurrences of each unique line: 

    ```go
        input := bufio.NewScanner(os.Stdin)
        for input.Scan() {
            counts[input.Text()]++
        }
    ```

5. For complete example see [.../ch1/dup1/main.go](https://github.com/ddgarrett/gopl.io/blob/master/ch1/dup1/main.go)
6. Printf format strings see https://golang.org/pkg/fmt/
7. See solution to PDF page 13 exercise 1.4, finding duplicate lines, in folder [ex1/1.4_dup2](./ex1/1.4_dup2).
8. See solutions to PDF page 15 animated gif exercises 1.5 and 1.6 in [ex1/1.5_lissajous](./ex1/1.5_lissajous) and [ex1/1.6_lissajous](./ex1/1.6_lissajous)
9. See solutions to PDF page 17 fetch exercises 1.7 - 1.9 in [ex1](./ex1)
10. See solution to PDF page 22 Lissajous server [1.12_lissajous_server](./ex1/1.12_lissajous_server)

## Chapter 2 - Program Structure
1. 25 reserved keywords:

```
break	    default	    func	interface   select
case	    defer	    go	        map	    struct
chan	    else	    goto	package	    switch
const	    fallthrough	    if	        range	    type
continue    for             import	return	    var
```

2. Three dozen predeclared names, not reserved but should not be redefined:

```
Constants:  true false iota nil
	
Types: 	    int int8 int16 int32 int64
	    uint uint8 uint16 uint32 uint64 uintptr
	    float32 float64 complex128 complex64
	    bool byte rune string error
        
Functions: make len cap new append copy close delete
	   complex real imag
	   panic recover
```

3. Variables:
    * Declared within function, local to function
    * Declared outside function visible to all files with same package
    * IF outside of a function AND var name begins with uppercase, it is exported to all packages
    * Camel case 
    * Acronyms and initialism maintain case, such as `htmlEscape` or `escapeHTML` **not** `escapeHtml`
4. Variables: (p. 30)
	* `var name type = expression`
	* `var b, f, s = true, 2.3, "four" // bool, float64, string`
	* `var f, err = os.Open(name) // os.Open returns a file and an error`  or
	* `f, err := os.Open(name)`
	* `t := 0.0`
	* `i, j = j, i // swap values of i and j`

5. Pointers (p. 32)

	* Pointer "zero" value is nil
	* See [./ch2/echo4](./ch2/echo4) for example of flag package which uses pointers to parse command-line arguments

	```go
	x := 1	
	p := &x         // p, of type *int, points to x
	fmt.Println(*p)	// "1"
	*p = 2	        // equivalent to x = 2
	fmt.Println(x)	// "2"
	```

6. Type declarations p. 39
	* defines new named type
	* conversion must be explicit even if same underlying data type, via T(x) where T = type
	* Types usually used for structures?
	* Types also allow you to define functions associated with the type - more on this in Chp 6
	* example of two non-comparable types and associating functions with a type. In particular, defining String() will be used by fmt package:

	```go
	// cannot directly compare two variables of these different types
	type Celsius float64
	type Fahrenheit float64
	```

	```go
	func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
	c := FToC(212.0)
	fmt.Println(c.String())  // "100°C"
	fmt.Printf("%v\n", c)    // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)    // "100°C"
	fmt.Println(c)           // "100°C"
	```

7. Packages and Files p. 41
	* One or more.go files within the same directory, usually all within directory same package name
	* Each package is a separate name space
	* Exported identifiers begin with uppercase letter
	* Comment preceding package documents the package
	* Only one file in the package should have such a comment and it should begin with a single line package summary
	* **TODO?** exercise 2.1 - add Kelvin to tempconv pkg in Ch2


8. Imports p. 42  **NOTE**: `GOPATH` has been deprecated by use of Modules
	* path, such as `import ("gopl.io/ch2/tempconv")` would bind the short name `tempconv` to the package, allowing subsequent references such as `tempconv.CToF` (see gopl.io/ch2/cf)
	* for above to work, would have to first execute `export GOPATH=$HOME/gobook` assuming source was intalled following directions at beginning of this document which would place it in `$HOME/gobook/src/gopl.io`
	* or `export GOPATH=$HOME/Documents/GitHub`, since I keep my gopl.io source in `/home/doug/Documents/GitHub/src/gopl.io`
	* `golang.org/x/tools/cmd/goimports` will add and remove import statements as needed

9. Package Initialization p. 44
	* Package level variables. Can reference other package level variables or functions.
	* If multiple .go in single package, order in which passed to compiler. Go tools sort names.
	* Within file, initialized in order declared, except dependencies resolved first (as encountered in declaration?)
	* `func init(){ /* ... */ }` **only** called automatically during init. Can not be referenced otherwise. For complex init, such as maps,tables.
	* One package init at a time, last is `main`

10. Scope of Variable p. 45
	* Carefull in use of `:=` where you **want** to keep access to outer defined variable

	```go
		if f, err := os.Open(fname); err != nil { // compile error: unused: f
			return err
		}
		f.Stat()	// compile error: undefined f
		f.Close()	// compile error: undefined f
	```

	```go
		var cwd string  // even if below compile, cwd would remain uninitialized
		func init() {
			cwd, err := os.Getwd() // compile error: unused: cwd
			if err != nil {
				log.Fatalf("os.Getwd failed: %v", err)
			}
		}
	```

## Chapter 3 - Basic Data Types

1. Integers (p. 51), Floating-Point (p. 56), complex numbers (p. 61), booleans (p. 63), Strings (p. 64), UTF-8 (p. 67), String to Number conversions (p. 75), constants (p. 75), constant generator iota (p. 77), untyped constants (p. 78)
2. Variety of bit operations and examples, see p. 53
3. **ToDo?**: exercises 3.1 - 3.4 on page 60. Different float64 svg drawings.
4. **ToDo?**: exercises 3.5 - 3.9 on page 62. Different complex number drawings and other high precision computations (big.Float and big.RAT)
5. String escape charactes, p. 66
6. Unicode - UTF-8 description p. 67
7. bytes.Buffer for efficient string manipulation p. 71
8. Important packages for strings: bytes, strings, strconv, unicocode - p. 71
9. Convert between string and []byte:
   
	```go
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	```

9. Common string functions:
    
	```go
	func Contains(s, substr string) bool
	func Count(s, sep string) int
	func Fields(s string) []string         // split string by white space
	func HasPrefix(s, prefix string) bool
	func Index(s, sep string) int
	func Join(a []string, sep string) string
	```

10. ALL above available in bytes package, using bytes instead of strings
11. Raw string literals with ` ``...`` ` backquotes instead of double quotes. No espcape sequences. May be spread over multiple lines.
	* useful for regular expressions, HTML templates, JSON literals, command usage messages and others which span multiple lines.
7. **TODO?**: exercises 3.10 - 3.12 on p. 74. String manipulation
8. Conversion between strings and Numbers - p. 75
9. Constants - p 75
    * `const pi = 3.14159 // approximately; math.Pi is a better approximation`
    * constant generator iota - p. 77  (exercise 3.13 p. 78)
    * untyped constants, up to 256 bits of precision - p. 78

## Chapter 4 - Composite Types

[Jupyter Chapter 4 code examples](http://localhost:8888/notebooks/home/GoFolder/gopherNotes/gopl_examples/gopl_examples.ipynb#chapter-4)

1. Arrays p. 81
	* `r := [...]int{99: -1}`  array of 100 ints, all 0 except last one
	* `func zero(ptr *[32]byte)` don't pass array directly, pass pointer instead
	* **TODO?**: exercises 4.1, 4.2 on p. 84. SHA hashes on array of bytes.
2. Slices p. 84
	* passing a slice to a function, actually passes and address to the underlying array, so no need to explicitly pass a pointer to the slice
	* slices not comparable
	* `runes = append(runes, r)`
	* **TODO?**: exercises 4.3 - 4.7 on p. 93. Various manipulations of slices.
	* if you need to test whether a slice is empty, use len(s) == 0 , not s == nil
  
	```go
	var s []int    // len(s) == 0  s == nil
	s = nil        // len(s) == 0  s == nil
	s = []int(nil) // len(s) == 0  s == nil
	s =[]int{}     // len(s) == 0  s != nil
	```

	```go
	t := make([]int,3,6)       // make([]T,len, cap), len == 3, cap == 6
	t = append(t, 100)         // len == 4, cap == 6
	t = t[:6]                  // now len == 6, okay as long as <= capacity
	t = append(t, 11)          // append 11 to the slice, len == 7, capacity doubles to 12
	v := make([]int,len(t),10) // len == 7, cap == 10
	copy(v,t)                  // copy to v from t, len == 7, cap == 10
	```

3. Maps aka Hash Tables p. 93
	*  `ages := map[string]int{"alice":31, "charlie": 34,} // initial values are optional`
	*  `delete(ages,"alice") // remove an entry`
	*  `ages["doug"] = 66  // add doug`
	*  `for name, age := range ages {...}`
	*  `ages["bob"] == 0 // true` 
	*  `age, ok := ages["bob"]  // ok == false`
	*  **TODO?**: exercises 4.8, 4.9 on p. 99. Using maps to keep counts by key.

4. Structs p. 99
	* Comparing structs - p. 104
	* Structu embedding - p. 104
	* Examples:

		```go
		type Employee struct {
			ID         int
			Name       string
			Address    string
			DoB        time.Time
			Position   string
			Salary     int
			ManagerID  int
		}

		dilbert.Salary += 5000  // dot notation, can use +=

		employeePtr := &dilbert // pointer can also use dot notation
		employeePtr.Salary

		employeePtr = new(Employee)  // NOTE: new(...) returns a point to the struct
		_,_ := fmt.Printf("%+v",employeePtr)

		// NOTE: lower case names not accessible outside same package
		// Can point to themself, can NOT contain themself
		type tree struct {
			value    int
			left, right *tree
		}

		// struct literals - IF small number of fields
		type Point struct{ X, Y int }
		p := Point{1, 2}

		// otherwise, use field names
		dilbert := Employee{ID:12345, Name: "Dilbert", Position: "software engineer", 
							DoB: parseDate("09/19/1985"), Salary: 125000}
		```

5. JSON - p. 107
   * There is also XML, ASN.1, and Google’s Protocol Buffers - supported by `encoding/json , encoding/xml , encoding/asn1`
   * Exercises 4.10 - 4.13 p. 112. Using JSON interface from GitHub, xkcd, [open movies](https://omdbapi.com/)

```go
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

// convert struct to JSON string
data, err := json.Marshal(movies)
fmt.Printf("%s\n", data)
/*  Output: 
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
*/

// more human readable, use indent
data, err := json.MarshalIndent(movies, "", "    ")

/* Output: 
[
    {
        "Title": "Casablanca",
        "released": 1942,
        "Actors": [
            "Humphrey Bogart",
            "Ingrid Bergman"
        ]
    },
	...
*/

// Unmarshalling similar - but we can selectively unmarshall fields
// NOTE: unmarshalling - names are case insensitive, so "color" unmarshalls to "Color"
// so no need for field tags just for different case names
var titles []struct{ Title string ; Color bool}
if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Printf("%+v \n",titles)
/* Output:
[{Title:Casablanca Color:false} {Title:Cool Hand Luke Color:true} {Title:Bullitt Color:true}] 
*/
```

6. Text and HTML Templates - p. 113



