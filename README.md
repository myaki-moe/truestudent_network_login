# TrueStudent Network Auto Login

TrueStudent Network Auto Login is a utility designed to automate the process of logging into the TrueStudent network. This software was born out of frustration with HTML webpage authentication, which is notoriously unfriendly to headless devices. By offering a command line based solution, this program provides a more streamlined and device-friendly alternative.

## Quick Start

To use this program, simply download the executable file and run it from your command line or terminal.

## Usage

This program offers several command-line arguments to customize its behavior:

- `-c int`: Set the retry count. This is the number of times the program will attempt to connect to the network before giving up. The default value is 3.

- `-d duration`: Set the connectivity check duration. This is the amount of time the program will wait before checking the network connection again. The default value is 5 minutes (5m0s).

- `-f`: Force login. Use this option if you want the program to attempt to log in immediately, regardless of the current network status.

- `-r`: Run once. Use this option if you want the program to attempt to log in only once, rather than continuously checking the network status and attempting to log in as needed.

- `-v`: Verbose output. Use this option if you want the program to provide detailed output about its actions and the current network status.

## Example

Here is an example of how you might use this program:

```
./truestudent_network_login -c 5 -d 10m0s -f -v
```

In this example, the program will:

- Attempt to connect to the network up to 5 times (`-c 5`)
- Wait 10 minutes before checking the network connection again (`-d 10m0s`)
- Attempt to log in immediately, regardless of the current network status (`-f`)
- Provide detailed output about its actions and the current network status (`-v`)

## Important Note

This program uses HTTPS connections. Therefore, it's crucial that you correctly configure your HTTPS root certificates. Failure to do so may result in connectivity issues. Please refer to your operating system's documentation or your network administrator for guidance on managing HTTPS root certificates.

## Build

This repository includes a `Makefile` that automates the process of building the application for different architectures with both static and dynamic linking. You need to have `make` and `Go` installed on your machine and make sure you have a `git` versioning system in your project.

To build the application for all architectures and linking types, simply run the following command:
```
make
```
This will create four output files in the `output` directory:

- `truestudent_network_login_amd64_static`: The application built for the linux/amd64 architecture with static linking.
- `truestudent_network_login_amd64`: The application built for the linux/amd64 architecture with dynamic linking.
- `truestudent_network_login_arm64_static`: The application built for the linux/arm architecture with static linking.
- `truestudent_network_login_arm64`: The application built for the linux/arm architecture with dynamic linking.

If you want to build the application for a specific architecture and linking type, you can run one of the following commands:

```
make build-linux-amd64-static
make build-linux-amd64
make build-linux-arm-static
make build-linux-arm
```

## Support

If you encounter any issues or have any questions about this program, please feel free to open an issue on this repository or contact the developer directly.