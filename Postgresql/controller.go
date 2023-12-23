package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "strconv"
)

// GetUser ...
func GetUser(c *gin.Context) {
    // Dummy data for demonstration
    // user := User{
    //     ID:   1,
    //     Name: "John Doe",
    // }

    user, err := GetUserByIDDB(1)
    fmt.Println("user: ", user)
    fmt.Println(err)
    c.JSON(http.StatusOK, user)
}

// AddRouter ...
func AddRouter(c *gin.Context) {
    var newRouter Router
    if err := c.ShouldBindJSON(&newRouter); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Here, you'd call a function to add the router to the DB
	err := CreateRouterDB(newRouter)
	if err != nil {
		fmt.Printf("%s",err)
	}
    c.JSON(http.StatusOK, gin.H{"message": "Router added", "router": newRouter})
}
// INSERT INTO Users (id, name, email, organization_id) VALUES (1, "John Doe", "jd@gmail.com", "org1")
// UpdateRouter ...
func UpdateRouter(c *gin.Context) {
    var updatedRouter Router
    if err := c.ShouldBindJSON(&updatedRouter); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    routerID := c.Param("id")
    routerIDcvt, _ := strconv.Atoi(routerID)
    // Here, you'd call a function to update the router in the DB
    err := UpdateRouterDB(routerIDcvt, updatedRouter)
	if err != nil {
		fmt.Printf("%s",err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Router " + routerID + " updated", "router": updatedRouter})
}

// DeleteRouter ...
func DeleteRouter(c *gin.Context) {
    routerID := c.Param("id")
    routerIDcvt, _ := strconv.Atoi(routerID)
    // Here, you'd call a function to delete the router from the DB
    err := DeleteRouterDB(routerIDcvt)
	if err != nil {
		fmt.Printf("%s",err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Router " + routerID + " deleted"})
}

// GetBillingInfo ...
func GetBillingInfo(c *gin.Context) {
    // Dummy data for demonstration
    billingInfo := Billing{
        ID:         1,
        UserID:     123,
        LicenseID:  "abc-123",
    }
    c.JSON(http.StatusOK, billingInfo)
}

// UpdateBillingRecord ...
func UpdateBillingRecord(c *gin.Context) {
    var updatedBilling Billing
    if err := c.ShouldBindJSON(&updatedBilling); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Here, you'd call a function to update the billing record in the DB
	err := UpdateBillingRecordDB(updatedBilling)
	if err != nil {
		fmt.Printf("%s",err)
	}
    c.JSON(http.StatusOK, gin.H{"message": "Billing updated", "billing": updatedBilling})
}
