// Package vars holds variables which are used throughout the program.
// The variables are set and retrieved to communicate between packages.
package vars

// TargetIP is the IP address of the target machine.
var TargetIP string
// Port is the port on which the program listens and sends data.
var Port = 9911

// SelectedInterface is the interface selected by the user when sending data.
var SelectedInterface string

// AvailableInterfaces is a list of all the available hosts on the subnet when
// sending data.
var ConnectedHosts []string