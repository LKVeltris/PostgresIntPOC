1. `docker pull postgres`
2. `docker run --name sqlserver -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres`

3. `docker exec -it sqlserver psql -U postgres`

4. queries:
```
CREATE ROLE myuser WITH LOGIN PASSWORD 'mypassword';
CREATE DATABASE mydb;
GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;

```
Then : `\q`

5. run `main.go`


Adding table:

```
CREATE TABLE Users (id SERIAL PRIMARY KEY, name TEXT, email TEXT, password_hash TEXT, organization_id INT);
CREATE TABLE Organizations (id SERIAL PRIMARY KEY, name TEXT, settings JSON);
CREATE TABLE Routers (id SERIAL PRIMARY KEY, user_id INT, organization_id INT, config JSON, status TEXT);
CREATE TABLE Billings (id SERIAL PRIMARY KEY, user_id INT, organization_id INT, billing_info JSON, license_id TEXT);
```

```
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Organizations;
DROP TABLE IF EXISTS Routers;
DROP TABLE IF EXISTS Billings;
```


```
-- Creating the Users table
CREATE TABLE Users (
    ID SERIAL PRIMARY KEY,
    Name TEXT NOT NULL,
    Email TEXT UNIQUE NOT NULL,
    OrganizationID INT,
    Settings TEXT
);

-- Creating the Organizations table
CREATE TABLE Organizations (
    ID SERIAL PRIMARY KEY,
    Name TEXT NOT NULL,
    Settings TEXT
);

-- Creating the Routers table
CREATE TABLE Routers (
    ID SERIAL PRIMARY KEY,
    UserID INT REFERENCES Users(ID),
    OrganizationID INT REFERENCES Organizations(ID),
    Config TEXT,
    Status TEXT
);

-- Creating the Billings table
CREATE TABLE Billings (
    ID SERIAL PRIMARY KEY,
    UserID INT REFERENCES Users(ID),
    OrganizationID INT REFERENCES Organizations(ID),
    BillingInfo TEXT,
    LicenseID TEXT
);
```