# extended-awx-client-go

This is an inofficial fork of Red Hat's "awx-client-go" Git Repo: moolitayer/awx-client-go.

## Usage
For general usage please have a look at original Git Repo mentioned above.

Following example shows how to gather all hosts regarding an inventory with their according groups.
The for loop is an example how to build inventory file as a map.

	hostsResource := connection.Hosts()

	// Get a list of all Hosts.
	getHostsRequest := hostsResource.Get()
	getHostsResponse, err := getHostsRequest.Filter("host_filter", "inventory__name=\"Demo Inventory\"").Send()
	if err != nil {
		panic(err)
	}

	// Build inventory from result
	hosts := getHostsResponse.Results()
	inventory := make(map[string][]Host, len(hosts))

	for _, host := range hosts {

		groupsArrayRef := host.GroupsArray()

		for i := 0; i < len(groupsArrayRef); i++ {
			tmpGroup := (*groupsArrayRef[i])
			inventory[tmpGroup] = append(inventory[tmpGroup], (*host))
		}
	}

## Additions
New objekt types can be gathered from awx:
- Groups
- Hosts


