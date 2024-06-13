CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL,
    permissions TEXT[],
    customer_id VARCHAR(255) NOT NULL
);

CREATE TABLE permissions (
    name VARCHAR(255) PRIMARY KEY
);

CREATE TABLE products (
    id VARCHAR(255) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    names JSONB,
    descriptions JSONB,
    categories TEXT[],
    prices JSONB,
    price_history JSONB,
    customer_id VARCHAR(255) NOT NULL
);

CREATE TABLE product_categories (
    id VARCHAR(255) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    customer_id VARCHAR(255) NOT NULL
);

CREATE TABLE coupons (
    id VARCHAR(255) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    promo_code VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL,
    unit VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP,
    categories TEXT[],
    is_active BOOLEAN NOT NULL,
    customer_id VARCHAR(255) NOT NULL
);

CREATE TABLE orders (
    id VARCHAR(255) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    total_cost FLOAT NOT NULL,
    currency CHAR(3) NOT NULL,
    country CHAR(2) NOT NULL,
    city VARCHAR(255) NOT NULL,
    postal_code VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    basket JSONB,
    tax_id VARCHAR(255),
    customer_id VARCHAR(255) NOT NULL
);

INSERT INTO permissions(name) VALUES 
('get-user'),
('get-user-self'),
('create-user'),
('update-user'),
('update-user-self'),
('delete-user'),
('delete-user-self'),
('create-product'),
('update-product'),
('delete-product'),
('create-category'),
('update-category'),
('delete-category'),
('create-coupon'),
('update-coupon'),
('delete-coupon'),
('get-orders'),
('update-order'),
('delete-order');

INSERT INTO users (id, username, password, email, is_active, permissions, customer_id, created_at, updated_at) VALUES (
    '243277f5-146d-4dc1-9c66-11b41ced6ead',
    'admin',                  
    'password',                 
    'admin@example.com',      
    true,                    
    ARRAY[
        'get-user',
        'get-user-self',
        'create-user',
        'update-user',
        'update-user-self',
        'delete-user',
        'delete-user-self',
        'create-product',
        'update-product',
        'delete-product',
        'create-category',
        'update-category',
        'delete-category',
        'create-coupon',
        'update-coupon',
        'delete-coupon',
        'get-orders',
        'update-order',
        'delete-order'
    ]::text[],                    
    'c6a2b5a1-6851-438b-a055-2ae0d1116b50',
    NOW(),
    NOW()   
);

INSERT INTO product_categories (id, name, created_at, updated_at, customer_id) VALUES 
('c1a2b5a1-6851-438b-a055-2ae0d1116b50', 'electronics', NOW(), NOW(), 'c6a2b5a1-6851-438b-a055-2ae0d1116b50'),
('d2a2b5a1-6851-438b-a055-2ae0d1116b51', 'books', NOW(), NOW(), 'c6a2b5a1-6851-438b-a055-2ae0d1116b50'),
('e3a2b5a1-6851-438b-a055-2ae0d1116b52', 'clothing', NOW(), NOW(), 'c6a2b5a1-6851-438b-a055-2ae0d1116b50'),
('f4a2b5a1-6851-438b-a055-2ae0d1116b53', 'home&kitchen', NOW(), NOW(), 'c6a2b5a1-6851-438b-a055-2ae0d1116b50'),
('g5a2b5a1-6851-438b-a055-2ae0d1116b54', 'sports', NOW(), NOW(), 'c6a2b5a1-6851-438b-a055-2ae0d1116b50');
