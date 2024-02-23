# Querying Specific events in Userup

This example demonstrates a simple app to query for event details for a specific event for a past X minutes. A hypothetical scenario is a customer has encountered an issue that the engineers have resolved. But to verify the issue is resolved and collect information if it does reoccur, we can check if this event ever reoccurs.

In the product code, whenever the issue would be encountered it is logged to Userup events. (See the [SendEmailInvite](https://github.com/Hillside-Labs/userservice-examples/tree/main/SendEmailInvite) example for a demonstration of logging events.)

This QueryEvent example tool is set up to run on a schedule to check if the event has occured in the past few minutes then produces an action. In this example it is printing the event details, but it could just as easily send an email or Slack notification.

```go
resp, err := client.SearchEvents(context.Background(),
    &userapi.SearchEventsRequest{
        UserId: &userapi.UserID{
            Id: userid,
        },
        Names: []string{"userup.io/example/issue-encountered"},
        Begin: timestamppb.New(since),
    })
```
This call to SearchEvents has multiple filters to extract only the specific events in question. The UserID, the event Name, and the time range of the events are all specified. Any event not matching all three conditions will not be returned in the response.