# gotask

Users need to register
• Users need to login
• Users cannot login until confirmed
• Users need to confirm Registration via email (use mailinator.com)
• To-dos can have subtasks
• All data needs to be validated (UI and backend)
• Users to be notified on creation or updates to tasks
• Use JWT as authentication layer for the API
• Describe a deployment process for this application

Naive Thought Process:

Create basic todo app first, before adding users, email confirmation, or authentication. Reasoning: learning Go, Echo, Quasar and Vue in one weekend is already a big task. Create a MVP first, extend later.

Step 2:

Register provides the route for registration and email confirmation
Login provides the route for authentication
Authentication is persisted by JWT

GET     /user/todos         show all user todos, with todo ids stored within the HTML elements as names or data attributes
POST    /user/todos/new     add user todo by multipart form, with :id dynamically determined by page javascript (Leave subtodos until v0.1 POC achieved - unnecessary complexity)
GET     /user/todos/:id     get user todo by id (could add an associated notes textfield here?)
PUT     /user/todos/:id     update todo by id
DELETE  /user/todos/:id     delete todo by id