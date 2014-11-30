####Dependencies

    go get github.com/stretchr/testify

####Run exercises

	go run exercise-....go

####Run tests

	cd exercise-7/mathematics/
	go test

####Go get

    go get code.google.com/p/go-tour/wc

####Exercise with HTTP and expose port

    $ sudo docker ps
    [sudo] password for zampieri:
    CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
    ef4c64eb792b        go:latest           "/bin/bash"         19 hours ago        Up 19 hours                             naughty_kirch

    $ sudo docker inspect go | grep IPAddress

    $ sudo docker inspect naughty_kirch | grep IPAddress
            "IPAddress": "172.17.0.4",

Access browser `http://172.17.0.4:8000/`, example in exercise-57.go.
