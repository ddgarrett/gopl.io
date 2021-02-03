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


## Chapter 1

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
7. See solution to PDF page 13 exercise 1.4, finding duplicate lines, in folder [ex1/1.4](./ex1/1.4).
8. See solutions to PDF page 15 animated gif exercises 1.5 and 1.6 in [ex1/1.5](./ex1/1.5) and [ex1/1.6](./ex1/1.6)
