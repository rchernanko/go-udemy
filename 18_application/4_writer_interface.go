package main

/*

The writer interface is really important to understand, it's found all through the standard library

But before that, let's look at json encode and decode and how they differ from Marshal and Unmarshal

From stack overflow - https://stackoverflow.com/questions/33061117/in-golang-what-is-the-difference-between-json-encoding-and-marshalling

"
	Marshal => String
	Encode => Stream

	Marshal and Unmarshal convert a string into JSON and vice versa.
	Encoding and decoding convert a stream into JSON and vice versa.

"

And another (from the same post):

"
	Generally, encoding/decoding JSON refers to the process of actually reading/writing the character data to a string
	or binary form. Marshaling/Unmarshaling refers to the process of mapping JSON types from and to Go data types and primitives.
"

Encode and Decode is basically for sending / receiving JSON to and from external applications

TODO up to 2:10

*/
