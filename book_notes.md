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
2. See solutions to PDF page 8 exercises 1.1 through 1.3 in folder `ex1`.
3. 