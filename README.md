# 1. creating a docker container for the postgres database
### docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5

# 2. check if the container is running

### docker ps
# 3. enter postgres terminal inside docker container
### docker exec -it api-todo psql -U postgres

# 4. inside postgres, create a database, a user and a password, respectively
### create database api_todo;
### create user user_todo;
### alter user user_todo with encrypted password '1122';

# 5. giving permission for the user to access the database
### grant all privileges on database api_todo to user_todo;

# 6. conect with to database
### \c api_todo;

# 7. create a table in api_todo
### create table todos ( id serial primary key, title varchar, description text, done bool default FALSE);

# 8. permission to change tables
### grant all privileges on all tables in schema public to user_todo;
### grant all privileges on all sequences in schema public to user_todo;

# 9. TO RUN THE PROGRAM, JUST RUN MAIN.GO IN THE ROOT FOLDER
### go run main.go
