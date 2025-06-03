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


**KILO CODE AGENT:**
I've successfully debugged the user authentication and authorization feature. Key issues resolved include:
1.  Correct loading of `JWT_SECRET` from the `.env` file.
2.  Resolution of API call connectivity issues by changing the web server port to `8081` (as `8080` was causing conflicts, possibly with the expected PostgreSQL port).
3.  Correction of the `DATABASE_URL` to use the actual PostgreSQL port (`5432`) and disable SSL (`sslmode=disable`) for local development.
4.  Rectification of database permissions for sequence generation (`users_id_seq`).
5.  Ensured JWT claims (`userID` and `userRole`) are correctly set by the `JWTMiddleware` and retrieved by the `Profile` handler using consistent context keys.

The `/api/register`, `/api/login`, and protected `/api/profile` endpoints are now working as demonstrated by successful API calls. You can register a user, log in to receive a JWT, and use that JWT to access protected user profile information.