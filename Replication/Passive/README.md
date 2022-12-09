# Passive-Replication

## How to run
### Running managers
1. Setup 3 terminal instances in the [manager directory](https://github.com/2rius/DiSys-Algorithms/tree/main/Replication/Passive/manager).
2. Execute `go run .` on each instance. All managers, except 1 (the primary manager), should start receiving heartbeats ("[HB] .").

### Running a frontend client
1. Setup a terminal instance in the [frontend directory](https://github.com/2rius/go-ass-5/tree/main/frontend)
2. Execute `go run .`. A list of commands should be printed, aswell as the message "listening on [frontends ip:port]". If multiple frontends are running, please use the `--port` flag to specify a unique port for each instance (avoid ports 5000-5002, as they are used by the managers).
3. Enter a desired command, as explained by the terminal.