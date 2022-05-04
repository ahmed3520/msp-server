# msp-server

back-end API built with Golang (GO) and mongodb..

Endpoints are 

```
/ping
/users/find
/users/create
/users/delete
/users/update
/events/find
/events/create
/events/delete
```
# /ping

Send a GET request,
Return a string..
```
pong
```

# /users/find

Send a GET request(only email as parameter).. Ex)
```
msp-server.onrender.com/users/find?email=msp@msp.com
```
Return a JSON
```
{
    "id": "6271194502**************"
    "name": "Ahmed Taha",
    "image": "",
    "email": "msp@msp.com",
}
```
# /users/create

Send a POST request with a json body. Ex)
```
{
  "name": "Ahmed Taha",
  "image": "<URL For IMAGE>",
  "email": "ataha******@gmail.com",
}
```
Return a JSON
```
{
  "_id": "<user id>",
  "name": "<user name>",
  "image": "<user image URL>",
  "eamil": "<user email>",
}
```
This endpoint creates a document in the users collection of the users database.

# /users/update
Send a GET request. Ex)

```
msp-server.onrender.com/users/update?email=msp@msp.com&field=name&value=ataha
```
```
{
  "_id": "<user id>",
  "name": "<user name>",
  "image": "<image URL>",
  "eamil": "<user email>",
}
```
This endpoint updates the field name of the user with specified email to the value of ataha and returns the updated user profile..

# /users/delete
Send a GET request. Ex)

```
msp-server.onrender.com/users/delete?email=msp@msp.com
```
Returns a json:

```
{
  "isRemoved": true
}
```
This endpoint deletes the user based on the user eamil.

# /events/create

Send a POST request with a json body. Ex)
```
{
	"name": "<Event Name>",
	"image": "<Event main image URL>",
	"thumbnail": "<EVent image thumbnail URL>",
	"speakersevent": [{
		"name": "<speaker name>",
    "image": "<Image URL of the speaker>",
		"email": "<speaker email>",
	}]
}
```
Return a JSON
```
{
  "_id": "<user id>",
  "name": "<user name>",
  "image": "<user image URL>",
  "eamil": "<user email>",
  "speakersevent":"<array of JSON of the speakers data>",
}
```
This endpoint creates a document in the events collection of the msp database.


# Errors

All the endpoints return an error in json format if something goes wrong. Ex)

```
{
   "Message": <message>,
   "Status": <status>,
   "Error": <error>
}
```
