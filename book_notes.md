# The Go Programming Language

## Table of Contents
- [The Go Programming Language](#the-go-programming-language)
	- [Table of Contents](#table-of-contents)
	- [Preface](#preface)
	- [Chapter 1 - Tutorial](#chapter-1---tutorial)
	- [Chapter 2 - Program Structure](#chapter-2---program-structure)
	- [Chapter 3 - Basic Data Types](#chapter-3---basic-data-types)
	- [Chapter 4 - Composite Types](#chapter-4---composite-types)
	- [Chapter 5 - Functions](#chapter-5---functions)
	- [Chapter 6 - Methods](#chapter-6---methods)
	- [Chapter 7 - Interfaces](#chapter-7---interfaces)

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

	```go
	var name type = expression
	var b, f, s = true, 2.3, "four" // bool, float64, string
	var f, err = os.Open(name) // os.Open returns a file and an error... or
	f, err := os.Open(name)
	t := 0.0
	i, j = j, i // swap values of i and j
	```

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
11. Raw string literals with `` `...` `` backquotes instead of double quotes. No espcape sequences. May be spread over multiple lines.
	* useful for regular expressions, HTML templates, JSON literals, command usage messages and others which span multiple lines.
12. **TODO?**: exercises 3.10 - 3.12 on p. 74. String manipulation
13. Conversion between strings and Numbers - p. 75
14. Constants - p 75
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
	*  **TODO?**: exercises 4.8, 4.9 on p. 99. Using maps to keep counts by key.
 
	```go
	ages := map[string]int{"alice":31, "charlie": 34,} // initial values are optional
	delete(ages,"alice") // remove an entry
	ages["doug"] = 66  // add doug
	for name, age := range ages {...}
	ages["bob"] == 0 // true
	age, ok := ages["bob"]  // ok == false
	```

1. Structs p. 99
	* Comparing structs - p. 104
	* Struct embedding - p. 104
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
		/*  Output: (note: new lines added below)
		[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},
		{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},
		{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
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
	* See also golang doc: [text/template](https://golang.org/pkg/text/template/) and [html/template](https://golang.org/pkg/html/template/)
	* Exercise 4.14 p. 117. Web server that queries GitHub.

		```go
		// convert boolean to Yes/No
		func yesNo(b bool) string {
			if b {
				return "Yes"
			}
			return "No"
		}

		// format a string array
		func formatStringArray(s []string) string {
			return strings.Join(s,", ")
		}

		const templ = `Movies List
		{{range .}}--------------------------------------------------------------------
		Title: {{.Title | printf "%-18q"}}  Released:{{.Year}}  Color: {{.Color | yesNo}} 
		Actors: {{.Actors | formatStringArray }}
		{{end}}`

		var report = textTemplate.Must(textTemplate.New("movieList").
			Funcs(textTemplate.FuncMap{"yesNo": yesNo, "formatStringArray": formatStringArray}).
			Parse(templ))

		report.Execute(os.Stdout, movies)

		/* Output:
		Movies List
		--------------------------------------------------------------------
		Title: "Casablanca"        Released:1942  Color: No 
		Actors: Humphrey Bogart, Ingrid Bergman
		--------------------------------------------------------------------
		Title: "Cool Hand Luke"    Released:1967  Color: Yes 
		Actors: Paul Newman
		--------------------------------------------------------------------
		Title: "Bullitt"           Released:1968  Color: Yes 
		Actors: Steve McQueen, Jacqueline Bisset
		*/
		```


## Chapter 5 - Functions

[Jupyter Chapter 5 code examples](http://localhost:8888/notebooks/home/GoFolder/gopherNotes/gopl_examples/gopl_examples.ipynb#chapter-5)

1. Function Declarations - p. 119.
	* No default parameter values, no ability to set by name
	* Passed by value
	* pointers, map, function, channel passed by reference
   
		```go
		func name(parameter-list) (result-list) {
			body
		}

		// If no body, function implemented in another language
		package math
		func Sin(x float64) float64 // implemented in assembly language
		```

2. Recursion p. 121

	* Exercises 5.1 - 5.4 p. 124. Traversing HTML documents.

3. Multiple Return Values p. 124
	* NOT limited to two return values: `func CountWordsAndImages(url string) (words, images int, err error) {`
	* "bare return" returns the values by name, so in above, `return` would be equivalent to `return words, images, err`. NOTE: this requires named results.
	* Exercises 5.5, 5.6 p. 127. Implement countWordsAndImages, modify `gopl.io/ch3/surface` to use named results and bare return

4. Errors p. 127
   1. Error-Handling Strategies p. 128
   2. End of File (EOF) p. 131

		* Section 7.11 presents a systematic way to distinguish certain errors. p. 206

5. Function Values p. 132
   	* functions are first class values in Go
	* may be assigned to variable or passed as a parameter
	* NOT comparable, can not be used as a key in a map
	* Exercises 5.7 - 5.9 p. 134. [See gophernotes verison of 5.7 - Produce Outline of HTML from Website](http://localhost:8888/notebooks/home/GoFolder/gopherNotes/gopl_examples/gopl_examples.ipynb#exercise-5.7) (only available on localhost server).

		```go
		// strings.Map applies a function to each character of a string , 
		// joining the results to make another string
		func add1(r rune) rune { return r + 1 }
		fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
		```

6. Anonymous Functions p. 135
	* named functions only at package level but can use anonymous within other functions
	* `strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")` same as previous code
	* Anonymous function can be assigned to a name, thus have a named function in effect
	* Exercises 5.10 - 5.14 p. 140. How to describe?
	* Be careful, in the example below the visitAll var must be defined before the function is defined. 


		```go

		// Error
		visitAll := func(items []string) {
			// ...
			visitAll(m[item]) // compile error: undefined: visitAll
			// ...
		}


		// this works
		var visitAll func(items []string) 
		visitAll = func(items []string) {
			// ...
			visitAll(m[item]) // okay - visitAll is defined
			// ...
		}
		```
	1. 5.6.1. Caveat: Capturation Iteration Variables p. 140
		* Pitfalls of lexical scope rules

7. Variadic Functions - varying number of arguments p. 142
	* `func sum(vals ...int) int {`  elipsis indicates varying number of arguments
	* to accept any type of variable use `interface{}` - see "Chapter 7 Interfaces"
  		*  `func errorf(linenum int, format string, args ...interface{}) {`
    * Exercises 5.15 - 5.17 p. 143. Using variadic functions for previously defined function. 
	* CAN call it with a slice using ellipses after the slice variable name:

		```go
		values := []int{1, 2, 3, 4}
		fmt.Println(sum(values...)) // "10"
		```
8. Deferred Function Calls p. 143
    * executed in reverse order they are declared
	* `defer resp.Body.Close()`
	* `defer f.Close()`
	* `defer mutex.Unlock()`
	* Exercise 5.18 p. 148. Use defer to close file in fetch function

9.  Panic p. 148
	* defers are executed, stack trace printed

10. Recover p. 151
    
## Chapter 6 - Methods

1. Exercises 6.1 through 6.5 p. 167. Extend a bit based set program.
2. Encapsulation: exported outside package through capitalization of names.p. 168
3. Getters **omit** `get...` prefix but use the `set...` prefix. Also omit `Fetch`, `Find` and `Lookup`.


## Chapter 7 - Interfaces

[Jupyter Chapter 7 Interfaces code examples](http://localhost:8888/notebooks/home/GoFolder/gopherNotes/gopl_examples/gopl_examples.ipynb#chapter-6-interfaces)

1. Interfaces as contracts - p. 171
   * reveals only *some* of the possible methods
   * for example `Fprintf(w io.Writer, ...` requires w implment the `Writer interface` which requires a structure which implements a `Write(p []byte) (n int, err error)` method. 
   * This can include `bytes.Buffer` or even your own custom writer, such as `ch7/bytecounter`
   * `fmt.Stringer` - another key interface. Defines `String()` method that returns a string
   * section 7.10 defines **how** `fmt` discovers which values satisfy the interface
   * Exercise 7.1 - 7.3 p. 173. Using interface to "extend" existing stgructures or wrap existing

		```go
		// ByteCounter only counts the number of bytes written.
		// Does not save the bytes anywhere?
		type ByteCounter int
		func (c *ByteCounter) Write(p []byte) (int, error) {
			*c += ByteCounter(len(p)) // convert int to ByteCounter
			return len(p), nil
		}
		```

2. Interface Types p. 174
   * Define minimum requirements to be considered an instance of the interface
   * Usually very fine grained. `io.Reader` and `io.Closer` each define a single method: `Read` can `Close`
   * Interfaces can also be embedded (combined). Example: `io.ReadWriter` combines `io.Reader` and `io.Writer` interfaces
   * Exercises 7.4, 7.5 p. 175. Parse HTML from a string. Implement a `LimitReader`

3. Interface Satisfaction p. 175
   * posses all methods interface requires - "is a"
   * note: that means a struct may satisfy an interface it didn't know about when it was created.
   * **any value** satisfies the *empty interface* `interface{}` - see *type asserstion* in 7.10

4. Parsing Flags with `flag.Value` p. 179
   * Exercise 7.6, 7.7 p. 181. Extend previous Celsius struct

5. Interface Values p. 101
6. Sorting with sort.Interface p. 186
   * See example at [Jupyter Chapter 7 Interfaces code examples](http://localhost:8888/notebooks/home/GoFolder/gopherNotes/gopl_examples/gopl_examples.ipynb#chapter-7-interfaces)
   * Note that `sort` package provides versions to sort `[]int, []string,` and `[]float64`.
   * Exercises 7.8 - 7.10 p. 191. Various uses of sort interface. 

		```go
		package sort
		type Interface interface {
			Len() int
			Less(i, j int) bool // i, j are indices of sequence elements
			Swap(i, j int)
		}

		// Sortable string slice
		type StringSlice []string
		func (p StringSlice) Len() int           { return len(p) }
		func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
		func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

		// sort it
		sort.Sort(StringSlice(names)) // where names is a regular slice
		```

7. The http.Handler interface p. 191
	* 


8. xxx

