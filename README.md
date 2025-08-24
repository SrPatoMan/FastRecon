# FastRecon
Fastrecon is a go tool that allows you to view the response from a list of subdomains through your proxy. This lets you to see any suspicious or notable difference between subdomains in an easy way.


<img width="1881" height="889" alt="image" src="https://github.com/user-attachments/assets/77ed8f92-cc4a-4b9a-89f9-5506c5bce237" />


## Installation


### ** NOTE: You need to install your Proxy's certificate in your OS's certificate store **   

<br>
<br>
You can install the tool with `go install` or download the repo and do `go build .` inside the repo

Installing with go install:

```go install github.com/SrPatoMan/fastrecon@latest```

## Usage

Examples of use:

Do requests to the root with default concurrency (20):

```fastrecon -l subdomains.txt```

Do requests to the root with 40 threads and following the redirects:

```fastrecon -l subdomains.txt -t 40 -redirect```

Do requests to a specific path:

```fastrecon -l subdomains.txt -p admin/admin.php```


## Options

```Usage of fastrecon:
    -l string
      	Subdomains file from your target
    -p string
      	URL Path. Example: -p .htaccess | -p directory1/directory2/endpoint.php
    -proxy string
      	Proxy address (default "http://127.0.0.1:8080")
    -redirect
      	Follow the redirects and logs the last request
    -t int
      	Amount of threads (default 20)
```
