# Products API

Welcome to the Products API! This API allows you to manage products, including creating, reading, updating, and deleting products. Below you'll find details on how to use this API, including endpoint information, request and response structures, and authentication requirements.

## Table of Contents

- [Getting Started](#getting-started)
  - [Base URL](#base-url)
  - [Authentication](#authentication)
- [Endpoints](#endpoints)
  - [Get List of Products](#get-list-of-products)
  - [Create Product](#create-product)
  - [Get Product By Id](#get-product-by-id)
  - [Update Product By Id](#update-product-by-id)
  - [Delete Product By Id](#delete-product-by-id)
- [Schemas](#schemas)
  - [Product](#product)

## Getting Started

### Base URL

The Base URL for the API is :

`http://127.0.0.1:8080/`

### Authentication

The API uses API Key authentication. You need to include your API Key in the `X-API-KEY` header for each request.

Example :

`X-API-KEY: YOUR_API_KEY_HERE`

## Endpoints

### Get List of Products

- **Endpoint** : `/api/v1/products`
- **Method** : GET
- **Description** : Retrieves a list of products.
- **Security** : Requires API Key.
- **Response** :

  - **200 OK** : Successful response containing a list of products.

  ```
  {
    "status": "OK",
    "code": 200,
    "data": [
        {
        "id": "string",
        "name": "string",
        "price": "number"
        }
    ]
  }
  ```

### Create Product

- **Endpoint** : `/api/v1/products`
- **Method** : POST
- **Description** : Creates a new product.
- **Security** : Requires API Key.
- **Request Body** :

  ```
  {
    "name": "string",
    "price": "number"
  }
  ```

- **Response** :

  - **201 Created** : Successful response containing the created product.

  ```
  {
    "status": "Created",
    "code": 201,
    "data": {
        "id": "string",
        "name": "string",
        "price": "number"
    }
  }
  ```

### Get Product By Id

- **Endpoint** : `/api/v1/products/{productId}`
- **Method** : GET
- **Description** : Retrieves a product by its id.
- **Security** : Requires API Key.
- **Parameters** :
  - `productId` (path): The Id of the product to retrieve.
- **Response** :

  - **200 OK** : Successful response containing the product.

  ```
  {
    "status": "OK",
    "code": 200,
    "data": {
        "id": "string",
        "name": "string",
        "price": "number"
    }
  }
  ```

### Update Product By Id

- **Endpoint** : `/api/v1/products/{productId}`
- **Method** : PUT
- **Description** : Updates a product by its id.
- **Security** : Requires API Key.
- **Parameters** :
  - `productId` (path): The Id of the product to update.
- **Request Body** :

  ```
  {
    "name": "string",
    "price": "number"
  }
  ```

- **Response** :

  - **200 OK** : Successful response containing the updated product.

  ```
  {
    "status": "OK",
    "code": 200,
    "data": {
        "id": "string",
        "name": "string",
        "price": "number"
    }
  }
  ```

### Delete Product By Id

- **Endpoint** : `/api/v1/products/{productId}`
- **Method** : DELETE
- **Description** : Deletes a product by its id.
- **Security** : Requires API Key.
- **Parameters** :
  - `productId` (path): The Id of the product to delete.
- **Response** :

  - **200 OK** : Successful response indicating the product has been deleted.

  ```
  {
    "status": "success",
    "code": 200,
    "data": null
  }
  ```

## Schemas

### Product

- **Type** : Object
- **Properties** :

  - `id` (string): The unique identifier for the product.
  - `name` (string): The name of the product.
  - `price` (number): The price of the product.

- **Example** :

  ```
  {
    "id": "string",
    "name": "string",
    "price": "number"
  }
  ```
