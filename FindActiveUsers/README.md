# Find Active Users gRPC

This example demonstrates searching for users based on an attribute value.

In this example we search for users who have a 'last_active' attribute timestamp of less than a certain number of days ago. This value could be set during the user's login or logout so we'll use this attribute to determine when the user was last active on the platform.

Here is the userdb code that performs the search:

```go
timeVal, err := structpb.NewValue(time.Now().AddDate(0, 0, -inactive))
if err != nil {
    log.Println(err)
}

attrFilter := userapi.AttributeFilter{
    Name: "last_active",
    Value: timeVal,
    Operator: userapi.Operator_LESS_THAN,
}

if err != nil {
    log.Println(err)
}

q := &userapi.UserQuery{
    OrderBy:    "username",
    AttributeFilters: []*userapi.AttributeFilter{&attrFilter},
}

userResponse, err := client.Find(context.Background(), q)
if err != nil {
    log.Println(err)
}
```

First the attribute filter is created to specify we are looking for users with a `last_active` attribute that is less than `inactive` days ago. `AttributeFilter` is a protobuf message that specifies the name of the attribute, the value to compare against, and the comparison operator.

Then that filter is supplied to the `UserQuery` protobuf message and passed to the `Find` method. `Find` will return all users that meet the criteria specified in the query including all AttributeFilters.
