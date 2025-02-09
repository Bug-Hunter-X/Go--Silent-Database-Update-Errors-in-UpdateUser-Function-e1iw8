func UpdateUser(user *User) error { 
  // ... some code to update user in database ...
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
}