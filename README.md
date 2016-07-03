# mongo-insert-demo
A rough demo of data insertion for Mongo and Go.

Sanitized data will be fed from another process and it will just parse the JSON file and insert it to the relevant Mongo document.

This is meant to run only on `localhost` so that another process residing in the same machine can access it and no external access is possible.

## Installation Instructions

* Type in the command `git clone https://github.com/code-shoily/mongo-insert-demo.git insert_server`
* Install all dependancies by typing `go get`
