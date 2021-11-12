## Shift 32 character porject installation
-------------------------------
1. The installation script is basically made for Linux operating system.
2. All source code contains under ```github/ashish041/``` directory
3. Project is implemented by Golang programming language.

## Project Installation
--------------------
1. `make all` command to install the whole project.
	make all command will do following operations.
	1. Install go
	2. Build application
	3. Run application on port 3333
	4. Run unit test
	5. Call api endpoint

	command
	-------
	make all

2. `make up` command  to run the project without installing go,
	in case go already installed in the operating system.
	`make up` command will do following operations.
	1. Build application
	2. Run app on port 3333
	3. Run unit test
	4. Call api endpoint
	
	command
	-------
	make up

3. `make build` command to build the application.

	command
	-------
	make build

4. `make run` command to run the application on port 3333.

	command
	-------
	make run

5. `make test` command to run unit test.

	command
	-------
	make test

6. `make callApi` command to call the endpoint.

	command
	-------
	make callApi

7. `make docker-all` command to build a image and run the application on docker.

	command
	-------
	make docker-build

8. `make docker-build` command to build a image on docker for the application.

	command
	-------
	make docker-build

9. `make docker-run` command to run application on docker

	command
	-------
	make docker-run


## Questions And Answers
---------------------
1. What does each line do?

export CGO_ENABLED=0
CGO_ENABLED=0 means the flag is disabled to compile cgo (C and GO mixing code) code.
CGO_ENABLED is by default enabled means CGO_ENABLED=1. So if we turn off the flag
which rebuilds all pre-build packages.

export GO111MODULE=on
GO111MODULE=on enable to use GO module, which means GOPATH is no longer supported. The GO module installs packages by specific version and keeps track of each package's version in go.mod.

build:  go build -o www
This command builds GO code to generate an executable binary by name www.

2. What do you expect to happen if you type make?

make command will execute the first target of the Makefile. For my case, make all.

3. How can you improve the encryption/decryption algorithm?

To optimize the encryption/decryption algorithm I declare two global variables containing a-z lowercase alphabet and A-Z uppercase alphabet. I convert the input string to byte slice. Then I check if the input string is uppercase or lowercase by a loop. And based on that I calculate the position of the new 32 shift character in the global declaration variable. By this position I get the byte value of the new 32 shift character. Then replace the input character by a new 32 shift character. End the loop when all input characters are replaced. And finally we convert the byte slice to string to get the final output.

```
Formula I use to calculate the position of the new 32 shift character: (26 + (inputvalue - 'a') + (shift % 26)) % 26
For example if the input character is d.
alphabetLowerCase "abcdefghijklmnopqrstuvwxyz"
byte value of a is 97
byte value of d is 100
shift 32

new position = (26 + (100 - 97) + (shift % 26)) % 26
			=  (29 + 6) % 26
			=  35 % 26
			=  9
			=  j
If we count from 0, the value of position 9 is j in the global variable alphabetLowerCase.
```

4. How can you reduce the Docker image size?

I use a multi-stage build which divides Dockerfile into multiple stages. In The first stage I install Go tools and build Go code which generates an executable binary file and copies the binary file in the second stage. In this way final docoker image won’t have any unnecessary content except required executable binary.
