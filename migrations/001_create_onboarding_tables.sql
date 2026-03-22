CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE,
    phone TEXT,
    department TEXT,
    joining_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS education (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    degree TEXT,
    institution TEXT,
    year_of_completion INT
);

CREATE TABLE IF NOT EXISTS experience (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    company TEXT,
    role TEXT,
    start_date DATE,
    end_date DATE
);

CREATE TABLE IF NOT EXISTS documents (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    document_type TEXT,
    file_url TEXT,
    verified BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS assets (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    asset_name TEXT,
    assigned_date DATE
);