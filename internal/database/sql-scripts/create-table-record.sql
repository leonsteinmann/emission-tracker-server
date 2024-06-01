DROP TABLE IF EXISTS record;

CREATE TABLE record (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    datetime DATETIME NOT NULL,
    user_id INT,
    category_id INT NOT NULL,
    subcategory_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    unit_type VARCHAR(50) DEFAULT 'CO2e' NOT NULL,
    input_datetime DATETIME NOT NULL,
    PRIMARY KEY (id)
);

-- Insert example emission records into the record table
INSERT INTO record (name, datetime, user_id, category_id, subcategory_id, amount, unit_type, input_datetime) VALUES
('Avocado', NOW(), NULL, 1, 0, 12.34, 'CO2e', NOW()),
('iPhone', NOW(), NULL, 2, 0, 56.78, 'CO2e', NOW()),
('Bus', NOW(), NULL, 3, 0, 23.45, 'CO2e', NOW()),
('Electricity', NOW(), NULL, 4, 0, 34.56, 'CO2e', NOW());