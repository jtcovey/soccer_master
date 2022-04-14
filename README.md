## Soccer Master by Justin Covey

### Building
In order to build the application you will need Go installed on your machine. 
Instruction for doing so can be found [here](https://go.dev/doc/install).

Once you have Go on your machine, simply run `go build` in the base directory of the application.

The build results will be different if you are on Linux/Mac or Windows.

- For Linux/Mac, it will generate a `soccer_master` file 
which can be executed from the command line with the command `./soccer_master`

- For Windows, it will generate a `soccer_master.go.exe` file 
which can then be executed with the command `soccer_master.go.exe`.

In either environment you can either provide a file to be read as argument to the above command or use a pipe to feed input to the executable.
if no argument or pipe is provided, the app will await standard input.

### Testing
To test you will first need Go installed on your machine, see the Building instructions above.

With Go installed, to run the automated tests simply run `go test -v` from the projects 
main directory.


### Design

The design of the app is to turn either the file supplied or standard in into a bufio Scanner,
then read line by line. Each line is first split by comma, then each half is parsed for team 
name and score. Then points are awarded to the teams based on the game results.
The points are stored in two maps, one for the game day currently being played and one for the 
season as a whole.

The end of a game day is detected when a line is read and one or both of the teams playing has 
already played this game day. At that point the season map is updated with the days results, 
the day is cleared, and the three teams with the highest points so far in the season are announced
to standard out.
