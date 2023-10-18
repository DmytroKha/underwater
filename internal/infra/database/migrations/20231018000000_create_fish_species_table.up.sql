CREATE TABLE IF NOT EXISTS public.fish_species (
    id SERIAL PRIMARY KEY,
    reading_id INT REFERENCES sensor_readings(id),
    name VARCHAR(100),
    count INT
);
