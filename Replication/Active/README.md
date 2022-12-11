# Active-Replication
(not sequentially consistent (yet!!) & not linearisable)

## How to run
### Running managers
1. Setup 3 terminal instances in the [manager directory](https://github.com/2rius/DiSys-Algorithms/tree/main/Replication/Active/manager).
2. Execute `go run .` on each instance.
   1. If more than 5 managers, please specify --port flag for additional instances, e.g. `go run . --port 5030`.

### Running a frontend client
1. Setup a terminal instance in the [frontend directory](https://github.com/2rius/DiSys-Algorithms/tree/main/Replication/Passive/frontend)
2. Execute `go run .`. A list of commands should be printed.
   1. If more than 3 managers, please specify each manager's port as args, e.g. `go run . 5000 5001 5002 5003`.
3. Enter a desired command, as explained by the terminal.

## How to make sequentially consistent and/or linearisable
Linearisable implies sequentially consistent, but sequentially consistent doesn't imply linearisable

### Make sequentially consistent
Add logical clock timestamps (Lamport/Vector) to frontends/managers.

### Make linearisable
Add clock synchronization algorithm (Christian's/Berkeley) to frontends/managers.