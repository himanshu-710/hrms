
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,

    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,

    dob DATE,
    gender TEXT,
    blood_group TEXT,

    personal_email TEXT UNIQUE,
    mobile_no TEXT,
    work_no TEXT,

    department TEXT,
    date_of_joining DATE,

    employment_context_role TEXT DEFAULT 'EMPLOYEE',

    relations JSONB,

    is_active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE IF NOT EXISTS employee_education (
    id SERIAL PRIMARY KEY,

    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    degree TEXT,
    branch TEXT,
    university TEXT,

    cgpa_or_pct FLOAT,

    year_of_joining INT,
    year_of_completion INT
);

CREATE INDEX IF NOT EXISTS idx_employee_education_employee_id
ON employee_education(employee_id);



CREATE TABLE IF NOT EXISTS employee_experience (
    id SERIAL PRIMARY KEY,

    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    company TEXT,
    role TEXT,

    start_date DATE,
    end_date DATE,

    currently_working BOOLEAN DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_employee_experience_employee_id
ON employee_experience(employee_id);


CREATE TABLE IF NOT EXISTS employee_addresses (
    id SERIAL PRIMARY KEY,

    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    address_type TEXT, -- CURRENT / PERMANENT
    street TEXT,
    city TEXT,
    state TEXT,
    country TEXT,
    postal_code TEXT,

    UNIQUE(employee_id, address_type)
);

CREATE INDEX IF NOT EXISTS idx_employee_addresses_employee_id
ON employee_addresses(employee_id);


CREATE TABLE IF NOT EXISTS employee_identity_documents (
    id SERIAL PRIMARY KEY,

    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    doc_type TEXT, 
    encrypted_doc_number TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(employee_id, doc_type)
);

CREATE INDEX IF NOT EXISTS idx_employee_identity_documents_employee_id
ON employee_identity_documents(employee_id);



CREATE TABLE IF NOT EXISTS employee_documents (
    id SERIAL PRIMARY KEY,

    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    doc_category TEXT,
    file_name TEXT,
    s3_url TEXT,

    verification_status TEXT DEFAULT 'PENDING',

    verified_by INT,
    verified_at TIMESTAMP,
    rejection_note TEXT,

    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_employee_documents_employee_id
ON employee_documents(employee_id);

CREATE INDEX IF NOT EXISTS idx_employee_documents_status
ON employee_documents(verification_status);


CREATE TABLE IF NOT EXISTS employee_assets (
    id SERIAL PRIMARY KEY,

    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    asset_code TEXT,
    asset_type TEXT,
    asset_name TEXT,

    assigned_on DATE,

    acknowledgement_status TEXT DEFAULT 'PENDING',
    acknowledged_at TIMESTAMP,

    UNIQUE(employee_id, asset_code)
);

CREATE INDEX IF NOT EXISTS idx_employee_assets_employee_id
ON employee_assets(employee_id);