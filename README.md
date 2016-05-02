# Journey Planner

Finds the travel time between two co-ordinates by car, bicycle and walking.

## Endpoint

###/times

Params:  `start` and `end`

Each param must be a pair of comma-separated lattitude and longitude co-ordinates.

e.g. `/times?start=51.5034070,-0.1275920&end=51.4838940,-0.6044030`

Will give the following result:

```json
{
  "journey_times": {
      "car": "48 mins",
      "walking": "7 hours 25 mins",
      "bicycle": "2 hours 21 mins"
  }
}
```

Should an error occur, you will recieve a 400 response, with the response body
containing specifics about the error.

If there is no information about travelling between two coordinates (e.g. if one is in the middle of an ocean), you will receive a 200 response with the message:
`"not possible to calculate journey times between these coordinates"`

## How to run it

Clone the repo into `$GOPATH/src/github.com/aitkenster/`

```
cd journey-planner
make build
$GOPATH/bin/journey-planner
```

Visit localhost:8080 in your browser to run a query. An example one, with the coordinates to get from Lands End to John O'Groats is:
`http://localhost:8080/times?start=50.0657979,-5.714962&end=58.6367099,-3.1002123`

## To run the tests
```
make test
```

## Technologies Used
- Golang
- [Google Maps Distance Matrix API](https://developers.google.com/maps/documentation/distance-matrix/)
