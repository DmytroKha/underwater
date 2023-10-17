CREATE TABLE IF NOT EXISTS public.sensors (
    id SERIAL PRIMARY KEY,
    codename VARCHAR(50) NOT NULL,
    group_id INT REFERENCES sensor_groups(id),
    x FLOAT NOT NULL,
    y FLOAT NOT NULL,
    z FLOAT NOT NULL,
    data_output_rate INT NOT NULL
);

INSERT INTO sensors (codename, group_id, x, y, z, data_output_rate) VALUES
    ('alpha 1', 1, 10.0, 20.0, 30.0, 60),
    ('alpha 2', 1, 12.0, 22.0, 31.0, 60),
    ('beta 1', 2, 15.0, 25.0, 32.0, 90),
    ('gamma 1', 3, 18.0, 28.0, 33.0, 120);
