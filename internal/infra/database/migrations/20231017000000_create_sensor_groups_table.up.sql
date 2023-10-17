CREATE TABLE IF NOT EXISTS public.sensor_groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

INSERT INTO sensor_groups (name) VALUES
    ('alpha'),
    ('beta'),
    ('gamma');
