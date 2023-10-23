Objective
The objective of this test is to create an API in GO language to manage parking lot of a rental agency.

A car is characterized by the following elements:

Car model
Its registration number
Its mileage
Its condition (available or being rented).
Your goal is to create an API providing the functionality listed below using gorilla mux.

 

1. Listing available cars
GET /cars
The API returns a list of all the cars registered. Te response should include every car's model, its registration number, its mileage and its rental status are displayed.

2. Adding a Car
POST /cars

{
    "model": "Tesla M3",
    "registration": "BTS812",
    "mileage":6003
}
The API takes the new car to register, if the car already exists (the registration number already exists) the API reports an err, if not the API stores the new car and returns its inserted id
3. Rent a Car
POST /cars/:registration/rentals
The API uses the registration number of the car to be rented. If the car does not exist, the API reports an error; if it is already rented, the API indicates that it is already rented otherwise the car is marked as being rented.

4. Returning a Car
POST /cars/:registration/returns
The API uses the registration number of the car to be returned. If the car does not exist, the program reports an error; if the car was not marked as being rented, the program stipulates it otherwise the program takes the number of kilometers driven and adds them to the mileage of the car. The car is then marked as available

5. Database
You can either create a database using mysql, sqlite (or any engine)- or an in-memory database for the cars.

--

What will I be assessed on
If your code works
If you wrote tests
Your application structure
How you modeled the entities and relations
How complicated is your code, can a junior understand what you did without help
How long it took you to complete the task
Useful resources
https://github.com/gorilla/mux
https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html
https://github.com/benbjohnson/wtf
   

Good Luck!