## Tf
 
Tf is a command line tool to easily generate your Terraform configuration with an interactive prompt.

[![Go Report Card](https://goreportcard.com/badge/github.com/g14a/tf)](https://goreportcard.com/report/github.com/g14a/tf)
[![Go Workflow Status](https://github.com/g14a/tf/workflows/Go/badge.svg)](https://github.com/g14a/tf/workflows/Go/badge.svg)
![CodeQL](https://github.com/g14a/tf/workflows/CodeQL/badge.svg)

### Inspiration
Boredom in Covid-19

### Installation
* ```go get github.com/g14a/tf```  
* Or clone the master branch and run ```go install``` in the root directory.

### Features
* Provider and resource support.
* Boilerplate code without having to go to the official docs.  
* In place documentation of fields.
* Custom Terraform validators for ```int```,```bool```,```string```,```tags```  
* Currently supports AWS EC2, S3, RDS, ELB, Lambda, VPC
* Less development overhead(I guess :blush:)

[![asciicast](https://asciinema.org/a/p6e5I9fNEslVdcaKFAJHgRfdt.svg)](https://asciinema.org/a/p6e5I9fNEslVdcaKFAJHgRfdt)

#### Fetching Boilerplate code for a resource

1. Search through the resources for a given provider
```shell
tf resource -p <provider> -b
```
2. Directly provide the resource as well
```shell
tf resource -p <provider> -r <resource-in-the-provider> -b
```

[![asciicast](https://asciinema.org/a/IMsCtr687FYZKkjJEuJHjMvhH.svg)](https://asciinema.org/a/IMsCtr687FYZKkjJEuJHjMvhH)

### Stability
This is a highly work in progress project, but I do my best to keep it stable so that things don't break.
Please report issues if you find this tool useful and I will try to make time to resolve them.

### Contribution 
Please check [CONTRIBUTION.md](https://github.com/g14a/tf/blob/main/CONTRIBUTING.md)

### Roadmap
* Google and Azure
* Repeatable configuration