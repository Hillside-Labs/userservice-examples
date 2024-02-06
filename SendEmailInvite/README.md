# Email invites via the magic link

This feature allows users to create teams and invite their teammates to the app via a magic link.
Use case: I want my users to be able to create teams and invite their teammates to the app. I would also love to be able to find out which users were invited to use the app and which users are the biggest referrers.

We currently offer such capabilities through our api. 
- `/api/teams/new/form` - POST request creates a team.
- `/api/teams/new/form` - GET request allows to get a sample form for team creation with a `submit` button, which gets posted to the URL above.
- `/api/teams/invites/new/form` - POST request issues the invitation to the team.
- `/api/teams/invites/new/form` - GET request allows to get a sample form with inputs to issue an invitation with a `submit` button, which gets posted to the URL above.
- `/api/teams/invites/:id` - This magic link the invited user gets in the email. Once accepted, the user is added to the team, automatically authenticated and redirected to the redirect that was configured for this invite at an earlier stage via `/api/teams/invites/new/form ` request.