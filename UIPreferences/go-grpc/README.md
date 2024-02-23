# User Preferences Setter

This code snippet showcases a basic application that interacts with the userservice API to set user preferences. The application reads the user ID and service URI from environment variables, connects to the userservice using gRPC, and saves user preferences.
Prerequisites

Make sure to set the following environment variables before running the application:

    userid: The user ID for whom preferences will be set.
    userservice_uri: The URI of the userservice to connect to.

## Code Overview

The main function:

    Parses the user ID and service URI from environment variables.
    Establishes a gRPC connection to the userservice using the GetUserClient function.
    Saves user preferences using the SavePreferences function.

The GetUserClient function:

    Accepts the userservice URI.
    Establishes a gRPC connection with insecure credentials.
    Returns a UsersClient and the connection.

## Usage

Run the application and observe the logged output for any errors during the process. The user preferences, such as theme, font size, language, and notification settings, are set to default values in this example.

```bash
go run main.go
```

Notes

    Ensure that the userservice is running and accessible at the specified URI.
    Modify the preferences in the SavePreferences function as needed.

Feel free to adapt and integrate this code into your own projects as a starting point for interacting with the userservice API.