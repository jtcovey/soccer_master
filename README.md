## Soccer Master by Justin Covey

### Building/Running

In order to build the application you will need Go installed on your machine. 
Instructions for doing so can be found [here](https://go.dev/doc/install).

Once you have Go on your machine, simply run `go build` in the base directory of the application.

The build results will be different if you are on Linux/Mac or Windows.

- For Linux/Mac, it will generate a `soccer_master` file 
which can be executed from the command line with the command `./soccer_master`

- For Windows, it will generate a `soccer_master.exe` file 
which can be executed with the command `soccer_master.exe`.

In either environment the two main modes of operation are to either pipe results into application like so:
- `cat .\prompt\sample-input.txt | ./soccer_master`

or pass the filepath of a file to be read for input as an argument like so 
- `./soccer_master .\prompt\sample-input.txt`

If no argument or pipe is provided the app will await standard input.

### Testing
To test you will first need Go installed on your machine, see the Building instructions above.

With Go installed, to run the automated tests simply run `go test -v` from the projects 
base directory.


### Design

The design of the app is to turn either the file supplied or standard input into a bufio Scanner,
then read line by line. Each line is first split by comma, then each half is parsed for team 
name and score. Then points are awarded to the teams based on the game results.
The points are stored in two maps, one for the game day currently being played and one for the 
season as a whole.

The end of a game day is detected when a line is read and one or both of the teams playing has 
already played this game day. At that point the season map is updated with the days results, 
the day is cleared, and the three teams with the highest points so far in the season are announced
to standard out.
