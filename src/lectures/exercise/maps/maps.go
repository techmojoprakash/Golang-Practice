// //--Summary:
// //  Create a program to display server status. The server statuses are
// //  defined as constants, and the servers are represented as strings
// //  in the `servers` slice.
// //
// //--Requirements:
// //* Create a function to print server status displaying:
// //  - number of servers
// //  - number of servers for each status (Online, Offline, Maintenance, Retired)
// //* Create a map using the server names as the key and the server status
// //  as the value
// //* Set all of the server statuses to `Online` when creating the map
// //* After creating the map, perform the following actions:
// //  - call display server info function
// //  - change server status of `darkstar` to `Retired`
// //  - change server status of `aiur` to `Offline`
// //  - call display server info function
// //  - change server status of all servers to `Maintenance`
// //  - call display server info function

// package main

// import "fmt"

// const (
// 	Online      = 0
// 	Offline     = 1
// 	Maintenance = 2
// 	Retired     = 3
// )

// func getAllServersDetails(serverMap map[string]int) {
// 	fmt.Printf("total number of servers in serverMap is %v \n", len(serverMap))
// 	for key, value := range serverMap {
// 		fmt.Printf("server name is %v and status is %v \n", key, value)
// 	}
// }

// func changeServerStatus(serverMap map[string]int, singleserver string, status int) {
// 	serverMap[singleserver] = status

// }

// func main() {
// 	servers := []string{"tcs", "wipro", "adani", "paytm", "amazon"}

// 	//* Create a map using the server names as the key and the server status
// 	//  as the value
// 	//* Set all of the server statuses to `Online` when creating the map
// 	serverMap := make(map[string]int)
// 	serverMap[servers[0]] = Online
// 	serverMap[servers[1]] = Online
// 	serverMap[servers[2]] = Online
// 	serverMap[servers[3]] = Online

// 	fmt.Printf("server type is %T \n", serverMap)

// 	// call display server info function
// 	getAllServersDetails(serverMap)
// 	// change server status
// 	changeServerStatus(serverMap, servers[2], Offline)
// 	// change server status
// 	changeServerStatus(serverMap, servers[3], Retired)
// 	// call display server info function
// 	getAllServersDetails(serverMap)
// 	//  - change server status of all servers to `Maintenance`
// 	changeServerStatus(serverMap, servers[0], Maintenance)
// 	changeServerStatus(serverMap, servers[1], Maintenance)
// 	changeServerStatus(serverMap, servers[2], Maintenance)
// 	changeServerStatus(serverMap, servers[3], Maintenance)

// 	// call display server info function
// 	getAllServersDetails(serverMap)
// }
