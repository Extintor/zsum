CREATE TABLE category (
	id UUID PRIMARY KEY,
	name VARCHAR(255) NOT NULL
);

CREATE TABLE transaction (
	id UUID PRIMARY KEY,
	datetime TIMESTAMP NOT NULL,
	category_id UUID REFERENCES category(id) NOT NULL,
	memo VARCHAR(255) NOT NULL,
	value DECIMAL NOT NULL
);

INSERT INTO category (id, name)
VALUES ('cccc0320-a5ac-4991-a744-5059aefcd9da', 'TEST');

INSERT INTO transaction (id, datetime, category_id, memo, value)
VALUES ('66ed0829-c452-4a32-af68-5a86d0ee50db', '03/03/2021 02:03:04', 'cccc0320-a5ac-4991-a744-5059aefcd9da', 'Test Transaction', '-12.34');

