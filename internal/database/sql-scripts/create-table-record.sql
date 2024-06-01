DROP TABLE IF EXISTS record;

CREATE TABLE record (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    datetime DATETIME NOT NULL,
    user_id INT,
    category VARCHAR(50) NOT NULL,
    subcategory VARCHAR(50) NULL,
    amount DECIMAL(10, 2) NOT NULL,
    unit_type VARCHAR(50) DEFAULT 'CO2e' NOT NULL,
    input_datetime DATETIME NOT NULL,
    PRIMARY KEY (id)
);

-- Insert example emission records into the record table
INSERT INTO record (name, datetime, user_id, category, subcategory, amount, unit_type, input_datetime) VALUES
('Avocado', NOW(), NULL, 'Food', 'Avocado', 12.34, 'CO2e', NOW()),
('iPhone', NOW(), NULL, 'Consumption', 'iPhone', 56.78, 'CO2e', NOW()),
('Bus', NOW(), NULL, 'Transportation', 'Bus', 23.45, 'CO2e', NOW()),
('Electricity', NOW(), NULL, 'Household', 'Electricity', 34.56, 'CO2e', NOW());