# Alien invasion

Mad​ ​aliens​ ​are​ ​about​ ​to​ ​invade​ ​your world ​and​ ​this project simulates the invasion. The simulation requires a map with
the world cities. Under the resources folder in this project there are two files containing cities of Spain and Europe 
that can be used to test the simulation. The format of these files is as follows: The city name first, followed by 1-4
directions (north, south, east or west). Each one represents a road to another city that lies in that direction. For
example:

```
Madrid north=Barcelona west=Lisbon
Lisbon east=Madrid
Barcelona north=Paris south=Madrid
Paris north=London south=Barcelona
```

## Asumptions

* The city and each of the pairs are separated by a single space
* The directions are separated from their respective cities with an equals (=) sign
* The city names can only contain letters
* The directions must be lowercase and exactly one of the 4 mentioned before (north, south, east or west)

## The simulation

The simulation works as follows:

* A number of aliens are generated. This number is specified as a command-line argument
* The aliens start out at random cities on the map , and wander around randomly, following links
* Each iteration, the alien moves in any of the directions leading out of a city
* When two aliens end up in the same city, they fight and in the process kill each other and destroy the city. 
When a city is destroyed, it is removed from the map, and so are any roads that lead into or out of it
* Once a city is destroyed, aliens can no longer travel to or through it. This may lead to aliens getting "trapped"
* When the simulation finishes, the map with what is left is printed out in the console in the same format as in the input file

## Build & Run

To build the project just run the following command from the root path:

``go build``

This command builds main and leaves the result in the current working directory. The program requires two command-line 
arguments to work properly:

* The input file containing the map with all the cities in the format described before
* A second argument indicating the number of aliens we want to generate

To run the simulation just execute:

`./main resources/europe.txt 20`

The command before runs the simulation on the europe map generating 20 aliens. The project also contains test that can
be executed from the root path with the following commands:

```
go test ./...
go test -cover ./... (including coverage)
```

## Result

When the simulation ends, the map with the cities left is printed out in the same format as in the input file. Find below
an example:

```
The world after the alien invasion

Lisboa
Sevilla
Huesca Zaragoza=south
Zaragoza Huesca=north Bilbao=west Barcelona=east
Valencia Barcelona=north
Barcelona Valencia=south Zaragoza=west
Oviedo Santander=east
Santander Oviedo=west Bilbao=east
Bilbao Zaragoza=east Santander=west
```