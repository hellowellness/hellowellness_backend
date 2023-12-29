-- db creation
CREATE DATABASE hellowellness;
USE hellowellness;

-- configure logging
SET global general_log = on;

-- {En Workbench en Ubuntu, me ha generado error de permisos.}
SET global general_log_file='/var/log/mysql/mysql.log';
SET global log_output = 'file'; 

-- Data creation
CREATE TABLE tests (texto VARCHAR(255));
INSERT INTO tests (texto) values ("hola hellowellness");
SELECT * FROM tests;


CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    variants_ids JSON, -- Puedes usar JSON para almacenar un slice de strings
    seo_list JSON, -- Almacena los valores separados por coma o algún otro delimitador
    brand_id VARCHAR(255) NOT NULL,
    image_id VARCHAR(255) NOT NULL,
    images_ids JSON, -- Puedes usar JSON para almacenar un slice de strings
    description TEXT NOT NULL
);