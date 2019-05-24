# gotask

Users need to register
• Users need to login
• Users cannot login until confirmed
• Users need to confirm Registration via email (use mailinator.com)
 - Todo lists are protected via auth
 - Deploy application to one server that configures nginx to serve the public SPA files and run the Go API from the same host
• To-dos can have subtasks
• All data needs to be validated (UI and backend)
• Users to be notified on creation or updates to tasks
• Use JWT as authentication layer for the API
• Describe a deployment process for this application

Naive Thought Process:

Create basic todo app first, before adding users, email confirmation, or authentication. Reasoning: learning Go, Echo, Quasar and Vue in one weekend is already a big task. Create an MVP first, extend later.

Step 2: What I would do if I had more time...

1.0 Backend: JWT
    1.1 Setup JWTWithConfig
    1.2 Create a testing sub-group router /auth/*
    1.3 Create a dummy login route that has a hardcoded username/password that maps claims and returns a token
    1.4 Test dummy login with a GET request to a restricted handler that tests against mapped claims "user"

2.0 UI: Login
    2.1 On login, retrieve JWT token and store in localStorage
    2.2 Set up axios defaults to retrieve the authorisation header

    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('jwtToken');

    2.3 Place all protected routes into sub-group router and test against these settings
    2.4 On validation, store the username in localStorage so that it can be accessed in app

3.0 Backend: Register
    3.1 Add users table to db.migrate() (username must be distinct)
    3.2 Create a testing route that uses smtp/goemail to send an email and lock this down
    3.3 On register, store the user in the database along with a hash generated from their email
    3.4 Send a link to the user email address with the hash as a query parameter under a new route /email/:hash
    3.5 Have this link make a GET request to a verification route, that checks the has against stored db value. If valid, switch the user to activated (bool)
    3.6 Alter login function to require the attempting login to check against the db activation value

4.0 UI: Register
    4.1 After submitting details, redirect to a 'please check your email page'.
    4.2 After email hash validation, prompt a success alert, and redirect back to login page

5.0 Backend: User todos
    5.1 On todo creation, retreive user id and store this as a foreign key on the todo

6.0 UI: User todos
    6.1 Pass username along with created todo to API
    6.2 Add a seperate q-field to display the username of the person who created the todo

7.0 Backend: Sub-todos POC
    7.1 Add subtodos table to db.migrate, with todo_id as foriegn key
    7.2 Create a new route and handler that can insert into the subtodo table
    7.3 Create new route and handler that shows and individual todo, along with its linked subtodos

8.0 UI: Sub-todos POC
    8.1 New page that shows an individual main todo, with another list underneath for subtodos

9.0 Backend/UI: Sub-todo CRUD
    9.1 Layer in existing functions to handle both main and subtodos

10.0 Deployment: Personal use
    10.1 Create a deployment branch and set up with Procfile etc.
    10.2 Upload compiled backend code and sqlite3 database file to Heroku. Use os.Getenv to set secret keys
    10.3 Use Quasar CLI to build the static files, setup a minimal express server and upload these to another Heroku server, with axios/vue-router configured to point at the IP of the Go API.

11.0 Further research/learning
    11.1 Migrate to MySQL/Postgres as opposed to sqlite3
    11.2 Look into docker and create containers for the API, UI and database
