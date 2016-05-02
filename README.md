# Journey PlannerGives the time taken to travel between two co-ordinates by car, bicycle and walking.## Endpoint#/timeParams:  `start` and `end`Each param must be a pair of comma-separated lattitude and longitude co-ordinates.e.g. `/time?start=51.5034070,-0.1275920&end=51.4838940,-0.6044030This will give the following result:```{  "journey_times": {      "car": "48 mins", //2906      "bicycle": "2 hours 21 mins", //8475      "walking": "7 hours 25 mins" //26695  }}```Should an error occur, you will recieve a 400 response, with the response bodycontaining specifics about the error.## Technologies Used- Golang- Google Maps Distance Matrix API## How to run it