ToDo API
    This is a simple ToDo API built with Go, using Chi router and JWT for authentication. The API supports creating, updating, and deleting todo items for users. An admin user can manage all todos.

Getting Started

Prerequisites
- Go 1.16 or higher

Running the Project
1. Clone the repository:
   git clone <repository-url>
2. Navigate to the project directory:
   cd <project-directory>
3. Install dependencies:
   go mod download
4. Run the server:
   go run main.go

*Users:
	
User 1 (admin):
- Username: admin	
- Password: adminpw11

user 2:
- Username:user
- Password:userpw22


Endpoints
1. Login:
   - POST /login
   - Request Body (x-www.form-urlencoded): 
     {
       "username": "<username>",
       "password": "<password>"
     }
   - Response: 
     - 200: "Login successful. Token: <token>"
     - 401: "Login failed."

2. User Endpoints:
   - List ToDos:
     - GET /{username}/todos
   - Create ToDo:
     - POST /{username}/todos
     - Request Body: 
       {
         "description": "<description>",
         "percentComplete": <percent>
       }
   - Update ToDo:
     - PUT /{username}/todos/{id}
     - Request Body: 
       {
         "description": "<description>",
         "percentComplete": <percent>
       }
   - Delete ToDo:
     - DELETE /{username}/todos/{id}

3. Admin Endpoints:
   - List All ToDos:
     - GET /admin/todos
   - Create Admin ToDo:
     - POST /admin/todos
     - Request Body: 
       {
         "description": "<description>",
         "percentComplete": <percent>
       }
   - Update Any User's ToDo:
     - PUT /admin/users/{username}/todos/{id}
     - Request Body: 
       {
         "description": "<description>",
         "percentComplete": <percent>
       }
   - Delete Any User's ToDo:
     - DELETE /admin/users/{username}/todos/{id}



Note: Regular users can only manage their own ToDos. Admin can manage all users' ToDos.

