📦 Product API (Fiber + GORM + PostgreSQL)

This is a simple CRUD API for managing products, built with Go (Fiber framework) and GORM.

Base URL
http://localhost:3000/api/v1

📌 API Endpoints
1️⃣ Create a Product

Method: POST

URL:

http://localhost:3000/api/v1/products


Headers:

Content-Type: application/json


Body → raw → JSON:

{
"name": "iPhone 16",
"description": "Latest Apple smartphone",
"price": 129999.99,
"category": "Electronics",
"in_stock": true
}


✅ Response (example):

{
"id": 1,
"name": "iPhone 16",
"description": "Latest Apple smartphone",
"price": 129999.99,
"category": "Electronics",
"in_stock": true,
"created_at": "2025-09-17T12:30:45Z",
"updated_at": "2025-09-17T12:30:45Z"
}

2️⃣ Get All Products

Method: GET

URL:

http://localhost:3000/api/v1/products


✅ Response (example):

[
{
"id": 1,
"name": "iPhone 16",
"description": "Latest Apple smartphone",
"price": 129999.99,
"category": "Electronics",
"in_stock": true,
"created_at": "2025-09-17T12:30:45Z",
"updated_at": "2025-09-17T12:30:45Z"
}
]

3️⃣ Get Product by ID

Method: GET

URL:

http://localhost:3000/api/v1/products/1


✅ Response (example):

{
"id": 1,
"name": "iPhone 16",
"description": "Latest Apple smartphone",
"price": 129999.99,
"category": "Electronics",
"in_stock": true,
"created_at": "2025-09-17T12:30:45Z",
"updated_at": "2025-09-17T12:30:45Z"
}

4️⃣ Update a Product

Method: PUT

URL:

http://localhost:3000/api/v1/products/1


Headers:

Content-Type: application/json


Body → raw → JSON:

{
"name": "iPhone 16 Pro",
"description": "Updated Apple smartphone with better cameras",
"price": 149999.99,
"category": "Electronics",
"in_stock": false
}


✅ Response (example):

{
"id": 1,
"name": "iPhone 16 Pro",
"description": "Updated Apple smartphone with better cameras",
"price": 149999.99,
"category": "Electronics",
"in_stock": false,
"created_at": "2025-09-17T12:30:45Z",
"updated_at": "2025-09-17T12:40:10Z"
}

5️⃣ Delete a Product

Method: DELETE

URL:

http://localhost:3000/api/v1/products/1


✅ Response:

{
"message": "Product deleted successfully"
}