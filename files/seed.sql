-- Create a table for the PetCategory sub-entity
CREATE TABLE IF NOT EXISTs pet_categories (
        id SERIAL PRIMARY KEY, -- Use appropriate data type for your database
        name VARCHAR(255) NOT NULL
);

-- Create a table for the Pet entity
CREATE TABLE IF NOT EXISTs pets (
        id SERIAL PRIMARY KEY, -- Use appropriate data type for your database
        category_id INT, -- Foreign key to link with the PetCategory table
        name VARCHAR(255) NOT NULL,
        photo_urls TEXT[], -- Use an appropriate array type for your database
        status VARCHAR(255) NOT NULL
);
