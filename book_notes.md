# The Go Programming Language

## Preface

PDF version of book

1. Set environment variable **GOPATH** 

    ```
    export GOPATH=$HOME/gobook
    ```
                
2. Fetch, build, install hello world app

    ```
    go get gopl.io/ch1/helloworld
    ```
                
3. Run the app

    ```
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

    ```
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
    * IF outside ofa funciton AND var name begins with uppercase, it is exported to all packages
    * Camel case 
    * Acronyms and initialism maintain case, such as `htmlEscape` or `escapeHTML` **not** `escapeHtml`
4. 
