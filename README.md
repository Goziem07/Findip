# findip

`findip` is a command-line tool written in Go that resolves domain names to their corresponding IP addresses.

## Installation

To install `findip`, you can use the following `go get` command:

```
go get github.com/Goziem07/findip
```

## Usage

```
findip -l <domain_list_file> [-n] [-o <output_file>]
```

The tool accepts the following command-line options:
```
-l <domain_list_file>: Specifies the file containing a list of domain names to resolve.

-n: Optional. Exclude domain names from the output. Only IP addresses will be displayed.

-o <output_file>: Optional. Specifies the output file path to save the results.
```

## WRN:
The <domain_list_file> should contain one domain name per line.

## Example
Suppose you have a file named domains.txt with the following content:
```
google.com
github.com
example.com
```

To resolve the IP addresses for these domains, you can run the following command:

```
findip -l domains.txt
```

This will output the IP addresses of the domains:
```
example.com: 93.184.216.34
google.com: 142.250.185.174
github.com: 140.82.121.3
```

If you want to exclude the domain names from the output, you can use the -n option:

```
findip -l domains.txt -n
```
This will display only the IP addresses:

```
93.184.216.34
142.250.185.174
140.82.121.3
```

You can also save the results to a file using the -o option:
```
findip -l domains.txt -o results.txt
```
This will save the IP addresses to the results.txt file.
