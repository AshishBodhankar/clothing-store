# eComm


**Summary of Bugs Fixed:**

1. JWT_SECRET Not Loading: Resolved by ensuring the application runs from the project root so godotenv finds backend/.env, and by adding detailed logging to confirm loading.
2. Incorrect user.NewService Call: Corrected by creating user.Repository first and passing it to NewService.
3. Incorrect product.RegisterRoutes Call: Corrected by passing the jwtSecret as required by its signature.
4. Incorrect userHandler.RegisterRoutes Call: Corrected by passing the jwtSecret after its signature was updated.
5. API Call Hang/No Response (Port Conflict): Resolved by changing the Gin web server port from 8080 to 8081, as 8080 was likely conflicting with the expected PostgreSQL port or another service.
6. Database Connection Refused (Incorrect Port): Resolved by changing DATABASE_URL to use the correct PostgreSQL port 5432 instead of 8080.
7. Database SSL Error: Resolved by adding ?sslmode=disable to the DATABASE_URL.
8. Database Sequence Permission Error: Resolved by granting USAGE, SELECT permissions on users_id_seq to the ashishb database user.
9. "User ID not found in context" Error: Resolved by ensuring JWTMiddleware sets the "userID" claim into the context.
10. "Role not found in context" Error: Resolved by correcting the context key in the Profile handler from "role" to "userRole" to match what JWTMiddleware sets.
