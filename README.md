# Go: Silent Database Update Errors

This repository demonstrates a common error in Go applications: silently ignoring errors from database operations. The `UpdateUser` function attempts to update a user in a database but always returns `nil`, even if the database operation fails.  This prevents the application from properly handling errors and reporting them to the user.

The `bug.go` file contains the problematic code, while `bugSolution.go` provides a corrected version that handles errors appropriately.

## How to Reproduce

1. Clone this repository.
2. Run `bug.go`.  Observe that even if the database update fails (simulated here), the program prints "User updated successfully".
3. Run `bugSolution.go`.  Observe that the program now handles the error properly and prints an appropriate message.