LOAD DATA INFILE './brands.csv'
INTO TABLE brands
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '/n'
IGNORE 1 ROWS;
