//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func statusLabel(status int) string {
	var statusLabel = ""
	switch status {
	case Online:
		statusLabel = "Online"
	case Offline:
		statusLabel = "Offline"
	case Maintenance:
		statusLabel = "Maintenance"
	case Retired:
		statusLabel = "Retired"
	}
	return statusLabel
}

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)

func showServerStatus(servers map[string]int) {
	fmt.Println("\nThere are ", len(servers), "servers")
	fmt.Println("---")

	statuses := map[int]int{
		Online:      0,
		Offline:     0,
		Maintenance: 0,
		Retired:     0,
	}

	for _, status := range servers {
		statuses[status]++
	}

	for status, element := range statuses {
		statusLabel := statusLabel(status)

		fmt.Println(statusLabel, ":", element)
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverMap := make(map[string]int, len(servers))

	for i := 0; i < len(servers); i++ {
		serverMap[servers[i]] = Online
	}

	fmt.Println(serverMap)

	//  - call display server info function
	showServerStatus(serverMap)
	//  - change server status of `darkstar` to `Retired`
	serverMap["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serverMap["aiur"] = Offline
	//  - call display server info function
	showServerStatus(serverMap)
	//  - change server status of all servers to `Maintenance`
	for server, _ := range serverMap {
		serverMap[server] = Maintenance
	}
	//  - call display server info function
	showServerStatus(serverMap)

}
