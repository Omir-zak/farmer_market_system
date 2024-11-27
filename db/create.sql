CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role VARCHAR(20) NOT NULL,
                       city VARCHAR(100),
                       is_approved BOOLEAN DEFAULT FALSE
);


CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          farmer_id INT REFERENCES users(id),
                          name VARCHAR(100),
                          category VARCHAR(50),
                          price NUMERIC(10, 2),
                          quantity INT,
                          description TEXT,
                          images TEXT[], -- Array of image URLs
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
