package models

type Staff struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Role     string `json:"role"`     // e.g., Doctor, Nurse, Surgeon
    Shift    string `json:"shift"`    // e.g., Day, Night
}