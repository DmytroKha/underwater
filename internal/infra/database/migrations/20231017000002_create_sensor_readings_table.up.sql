CREATE TABLE IF NOT EXISTS public.sensor_readings (
    id SERIAL PRIMARY KEY,
    sensor_id INT REFERENCES sensors(id),
    timestamp TIMESTAMP NOT NULL,
    temperature FLOAT,
    transparency INT,
    fish_species JSONB
);
