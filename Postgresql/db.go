package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432 // default port
    user     = "postgres"       // your created username
    password = "mysecretpassword"   // your created password
    dbname   = "postgres"         // your created database name
)

var db *sql.DB


func initDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        panic(fmt.Sprintf("Unable to connect to the database: %v\n", err))
    }

    // Check the connection
    err = db.Ping()
    if err != nil {
        panic(fmt.Sprintf("Unable to reach the database: %v\n", err))
    }
}

// GetUserByID retrieves a user by their ID
func GetUserByIDDB(id int) (*User, error) {
    user := User{}
    query := `SELECT id, name, email, organizationid FROM Users WHERE id=$1;`
    err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.OrganizationID)
    if err != nil {
        fmt.Printf("no user found with id %d ", id)
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("no user found with id %d", id)
        }
        return nil, fmt.Errorf("query error: %v", err)
    }
    return &user, nil
}

// CreateRouter adds a new router to the database
func CreateRouterDB(router Router) error {
    query := `INSERT INTO Routers (userid, organizationid, config, status) VALUES ($1, $2, $3, $4);`
    _, err := db.Exec(query, router.UserID, router.OrganizationID, router.Config, router.Status)
    return err
}

// UpdateRouter updates an existing router
func UpdateRouterDB(id int, router Router) error {
    query := `UPDATE Routers SET user_id=$1, organization_id=$2, config=$3, status=$4 WHERE id=$5;`
    _, err := db.Exec(query, router.UserID, router.OrganizationID, router.Config, router.Status, id)
    return err
}

// DeleteRouter removes a router from the database
func DeleteRouterDB(id int) error {
    query := `DELETE FROM Routers WHERE id=$1;`
    _, err := db.Exec(query, id)
    return err
}

// GetBillingInfo retrieves billing information
func GetBillingInfoDB(userID int) (*Billing, error) {
    billing := Billing{}
    query := `SELECT id, user_id, organization_id, billing_info, license_id FROM Billings WHERE user_id=$1;`
    err := db.QueryRow(query, userID).Scan(&billing.ID, &billing.UserID, &billing.OrganizationID, &billing.BillingInfo, &billing.LicenseID)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("no billing info found for user id %d", userID)
        }
        return nil, fmt.Errorf("query error: %v", err)
    }
    return &billing, nil
}

// UpdateBillingRecord updates a billing record
func UpdateBillingRecordDB(billing Billing) error {
    query := `UPDATE Billings SET billing_info=$1, license_id=$2 WHERE id=$3;`
    _, err := db.Exec(query, billing.BillingInfo, billing.LicenseID, billing.ID)
    return err
}