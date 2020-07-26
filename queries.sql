# tables created on elephantsql

CREATE TABLE product (
	product_id serial PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL,
	description VARCHAR ( 255 ) NOT NULL,
	category_id int,
	FOREIGN KEY (category_id)
      REFERENCES category (category_id)
);


CREATE TABLE category (
	category_id serial PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL,
	description VARCHAR ( 255 ) NOT NULL	  
);

SELECT * FROM PRODUCT

