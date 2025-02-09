func UpdateUser(user *User) error { 
  // ... some code to update user in database ...
  // Simulate a potential database error
  if user.ID == 0 {
    return fmt.Errorf("invalid user ID")
  }
  // ... actual database update code ... 
  return nil
}

func main() { 
  user := &User{ID: 1, Name: "John Doe"} 
  err := UpdateUser(user) 
  if err != nil { 
    fmt.Println("Error updating user:", err) 
  } else { 
    fmt.Println("User updated successfully") 
  }
  user2 := &User{ID: 0, Name: "Jane Doe"}
  err = UpdateUser(user2)
  if err != nil {
    fmt.Println("Error updating user:", err)
  } else {
    fmt.Println("User updated successfully")
  }
} 

type User struct {
  ID   int
  Name string
}