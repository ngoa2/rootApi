CREATE TABLE if not exists brands (
  id int not NULL auto_increment primary key,
  BrandName VARCHAR(100) not NULL,
  Environment int not NULL,
  EthicalPractices int not NULL,
  Transparency int not NULL,
  Average DECIMAL(3,2) not NULL,
  Tags VARCHAR(256) not NULL,
  AltBrands VARCHAR(3000) not NULL
);