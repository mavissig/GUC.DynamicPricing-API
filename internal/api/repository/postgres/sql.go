package postgres

var (
	sqlInitTables = `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		DO
		$$
		BEGIN
        	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender') THEN
            	CREATE TYPE gender AS enum ('male', 'female');
        	END IF;
    	END
		$$;

        CREATE TABLE IF NOT EXISTS Images
        (
            id    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
            image bytea
        );

        CREATE TABLE IF NOT EXISTS Address
        (
            id      serial PRIMARY KEY,
            country varchar(255),
            city    varchar(255),
            street  varchar(255)
        );

		CREATE UNIQUE INDEX if NOT EXISTS unique_address on address(country, city, street);
        CREATE TABLE IF NOT EXISTS Client
        (
            id                SERIAL PRIMARY KEY,
            client_name       VARCHAR(255),
            client_surname    VARCHAR(255),
            birthday          date,
            gender            VARCHAR(255),
            registration_date TIMESTAMP,
            address_id        int REFERENCES address (id) ON DELETE SET NULL
        );

		CREATE TABLE IF NOT EXISTS Supplier
        (
            id           serial PRIMARY KEY,
            name         varchar(255),
            address_id   int REFERENCES address (id) ON DELETE SET NULL,
            phone_number varchar(255)
        );

        CREATE TABLE IF NOT EXISTS Product
        (
            id               serial PRIMARY KEY,
            name             varchar(255),
            category         varchar(255),
            price            money,
            available_stock  int,
            last_update_date date,
            supplier_id      int REFERENCES supplier (id) ON DELETE CASCADE,
            image_id         UUID REFERENCES images (id) ON DELETE SET NULL
        );
	`

	sqlResetTables = `
	DROP TYPE IF EXISTS gender CASCADE;

	DROP TABLE IF EXISTS images CASCADE;
	DROP TABLE IF EXISTS address CASCADE;
	DROP TABLE IF EXISTS Client CASCADE;
	DROP TABLE IF EXISTS Product CASCADE;
	DROP TABLE IF EXISTS supplier CASCADE;

	DROP INDEX IF EXISTS unique_address CASCADE;

	DROP EXTENSION if EXISTS "uuid-ossp" CASCADE;
	`

	// -----------------------------------------------------------------
	// sql queries for client
	sqlAddClient = `
	INSERT INTO Client(client_name, client_surname, birthday, gender, registration_date, address_id)
	VALUES ('%s', '%s', '%s', '%s', '%s', %d);
	`

	sqlGetClientByName = `
	SELECT * FROM client
	WHERE client_name = '%s' AND client_surname = '%s';
	`

	// -----------------------------------------------------------------
	// sql queries for product
	sqlAddProduct = `
	insert into product(name, category, price, available_stock, last_update_date, supplier_id, image_id)
	VALUES ('%s','%s', %f, %d, '%s', %d, %s)
	`

	sqlDecProduct = `
	-- Params:
	-- 		[1]- id
	-- 		[2]- N
	DO
	$$
       DECLARE
			v_id bigint := %v;
			v_n bigint := %v;
	   BEGIN
		   if (SELECT available_stock from product WHERE id = v_id) < v_n THEN
			   RAISE EXCEPTION 'Not enough product quantity';
			   else
			   update product SET available_stock = available_stock - v_n
			   WHERE id = v_id;
		   END IF;
	   END;
	$$;
	`
	sqlGetProduct = `
	SELECT id, name, category, price::numeric::float8, available_stock, last_update_date, supplier_id, image_id FROM product %s
`
	// -----------------------------------------------------------------
	// sql queries for supplier

	sqlAddSupplier = `
	INSERT INTO supplier (name, address_id, phone_number)
	VALUES ('%s', %d, '%s');
`
	// -----------------------------------------------------------------
	// sql queries for image

	sqlAddImage = `
	WITH id_img as (
		INSERT INTO images (image) VALUES ($2)
		RETURNING id
	)
	UPDATE product SET image_id = (SELECT id_img.id FROM ID_IMG)
	WHERE product.id = $1;
`
	sqlChangeImageById = `
	UPDATE images SET image = $1
	WHERE id = '%s';
`
	sqlGetImageByProductId = `
	SELECT img.image FROM product
	JOIN images img ON product.image_id = img.id
	WHERE product.id = %d;
`
	sqlGetImageById = `
	SELECT image FROM images
	WHERE images.id = '%s';
`
	sqlDeleteImage = `
	DELETE from images 
	WHERE id = '%s';
`
	// -----------------------------------------------------------------
	// Common sql queries
	sqlCommonGetById = `
	SELECT * FROM %s
	WHERE id = %d;
	`

	sqlCommonGet = `
	SELECT * FROM %s
	`

	sqlDeleteById = `
	DELETE FROM %s WHERE id = %d;
	`

	sqlCommonChangeAddress = `
	-- Params:
	-- 		[1]- country
	-- 		[2]- city
	-- 		[3]- street
	-- 		[4]- id

	with addr as (
    INSERT into address(country, city, street)
    VALUES ($1, $2, $3)
           on CONFLICT (country, city, street) DO NOTHING
           RETURNING id as tmp_id
)

	UPDATE %s set address_id = (coalesce((select tmp_id from addr), (SELECT id from address WHERE country = $1 AND city = $2 AND street = $3)))
	WHERE id = $4;
`

	//todo: debug insert
	sqlDebugClient = `
	INSERT into address(country, city, street)
    VALUES ('Russia', 'Novosibirsk', 'Svobodi 513');

	insert into client(client_name, client_surname, birthday, gender, registration_date, address_id)
	VALUES ('test', 'test','1997-05-10', 'male', now()::timestamp, 1);
`
)
