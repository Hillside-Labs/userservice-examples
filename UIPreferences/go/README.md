# User Preferences

This code snippet showcases a basic application that interacts with the userservice API to set user preferences using the userup Go client The application reads the user ID and service URI from environment variables, connects to the userservice using gRPC, and saves user preferences.
The preferences are stored in a struct and saved as a single user Attribute. This data is serialized and stored as JSON.

