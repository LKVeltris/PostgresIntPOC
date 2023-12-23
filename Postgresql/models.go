package main

type User struct {
    ID             int    `json:"id"`
    Name           string `json:"name"`
    Email          string `json:"email"`
    OrganizationID string `json:"organization_id"`
    Settings       string `json:"settings"`
}

type Organization struct {
    ID       int      `json:"id"`
    Name     string   `json:"name"`
    Settings string `json:"settings"`
}

type Router struct {
    ID             int      `json:"id"`
    UserID         int      `json:"user_id"`
    OrganizationID int      `json:"organization_id"`
    Config         string `json:"config"`
    Status         string   `json:"status"`
}

type Billing struct {
    ID             int      `json:"id"`
    UserID         int      `json:"user_id"`
    OrganizationID int      `json:"organization_id"`
    BillingInfo    string   `json:"billing_info"`
    LicenseID      string   `json:"license_id"`
}