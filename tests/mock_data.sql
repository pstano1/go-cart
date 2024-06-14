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

INSERT INTO coupons (id, created_at, updated_at, promo_code, amount, unit, categories, is_active, customer_id) VALUES
('505892df-ac66-42f6-9fab-74fd03dbc5f3', NOW(), NOW(), 'MAY20', 20, 'percentage', ARRAY[
    'books', 
    'sports'
    ]::text[], 
    true, 
    'c6a2b5a1-6851-438b-a055-2ae0d1116b50'
);

INSERT INTO products(id, created_at, updated_at, names, descriptions, categories, prices, price_history, customer_id) VALUES 
(
    '0883981f-dd7e-436f-a753-be8172324e28', 
    NOW(), 
    NOW(),
    '{"en": "Example Product", "pl": "Przykładowy Produkt"}',
    '{"en": "This is an example product description.", "pl": "To jest przykładowy opis produktu."}',
    ARRAY['electronics']::text[],
    '{"USD": 50.25, "EUR": 50}',
    NULL,
    'c6a2b5a1-6851-438b-a055-2ae0d1116b50'
);


INSERT INTO orders(id, created_at, updated_at, total_cost, currency, country, city, postal_code, address, status, basket, tax_id, customer_id) VALUES
(
    '4709fbb1-6462-4518-859c-f804b43b0d2e', 
    NOW(), 
    NOW(), 
    100.50, 
    'USD',
    'US', 
    'New York', 
    '10001', 
    '123 Main St', 
    'placed', 
    '{"items": [{"0883981f-dd7e-436f-a753-be8172324e28": "Example Product", "quantity": 2, "price": 50.25}]}', 
    '1234567890', 
    'c6a2b5a1-6851-438b-a055-2ae0d1116b50'
);
