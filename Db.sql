CREATE TABLE mst_customer (
    id serial PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    email VARCHAR(20),
    no_hp VARCHAR(20)
);

CREATE TABLE mst_employee (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    no_hp VARCHAR(30)
);

CREATE TABLE mst_product (
    id serial PRIMARY KEY,
    name varchar(100),
    price int,
    satuan varchar(30)
);

CREATE TABLE transaction (
    id INT PRIMARY KEY,    
    tanggal_masuk DATE,
    tanggal_selesai DATE,
    employee_id INT NOT NULL,
    customer_id INT NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES mst_customer(id),
    FOREIGN KEY (employee_id) REFERENCES mst_employee(id)
);

CREATE TABLE transaction_detail (
    id INT PRIMARY KEY,    
    product_id int, 
    quantity INT NOT NULL,
    total INT,
    FOREIGN KEY (product_id) REFERENCES mst_product(id)
);

INSERT INTO mst_customer (name, email, no_hp)
VALUES
    ('John ', 'john@gmail.com', '08937343324'),
    ('Kevin', 'Kevin@gmail.com', '0384832473');

INSERT INTO mst_employee (name, no_hp)
VALUES
    ('Alice ', '0893736123'),
    ('Jojo', '08926251463');

INSERT INTO mst_product (name, price, satuan)
VALUES
    ('Laundry Baju ', 7000, 'Kilogram'),
    ('Laundry Boneka', 10000, 'pcs');

INSERT INTO transaction (id, tanggal_masuk, tanggal_selesai, employee_id, customer_id)
VALUES
    (1, '2023-09-21', '2023-09-22', 1, 1),
    (2, '2023-09-22', '2023-09-23', 2, 2);

INSERT INTO transaction_detail (id, product_id, quantity, total)
VALUES
    (1, 1, 3, 21000), 
    (2, 2, 2, 20000);  