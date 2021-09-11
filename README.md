# Garage-2021-Hackathon
Demonstration of openconfig ygot yang compiler

# Step 1: generate yang structs
clone the ygot repo: https://github.com/openconfig/ygot
set enviroment variable YGOT to the repo path
run 'make generate', this will create go structs corresponding to yang/animal.yang in structs/structs.go

# Step 2 : build and run rest server
go build, this will create exe 'hack.com'
execute ./hack.com

# Step 3 : demonstrate rest server
It would be useful to read the code before this step, but basically we have a rest interface to the yang model.

Try get for animal/cat\n
curl localhost:3000/animal/cat
{"does" : "meow" }
The cat meows.

Try setting it,
cat test1.json
{
   "does" : "meowwwwww"
}
curl -X POST -d @test1.json localhost:3000/animal/cat

curl localhost:3000/animal/cat
{"does" : "meowwwwwwwww" }
The cat meowwwwwws.

Now, note the regex for the cat, [A-Za-z]+. The cat can not understand numbers.
cat test1.json
{
   "does" : "meow1"
}
curl -X POST -d @test2.json localhost:3000/animal/cat
/device/cat/does: schema "does": "meow1" does not match regular expression pattern "^([A-Za-z]+)$"

# What did we demonstrate?
In essence we demonstate translation of a simple yang model to compiled structs and it's usage by a simple rest server. It can be observed that the code is largely dynamic, and that metadata such as the pattern can be used easily.
