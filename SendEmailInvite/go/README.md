# Recording Email Invites in UserService

This example demonstrates creating a user and logging events to the UserService using the userup Go client.

In this example we are sending an email invite to a user. On invite, a user is created with an invite status added to its attributes. An email invite event is also create to record the details of the invite.