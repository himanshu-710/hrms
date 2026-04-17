
DO $$ BEGIN
    CREATE TYPE gender_enum AS ENUM ('MALE','FEMALE','NON_BINARY','PREFER_NOT_TO_SAY');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE marital_status_enum AS ENUM ('SINGLE','MARRIED','DIVORCED','WIDOWED');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE employment_type_enum AS ENUM ('PERMANENT','CONTRACT','INTERN','PROBATION');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE address_type_enum AS ENUM ('CURRENT','PERMANENT');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE ownership_type_enum AS ENUM ('OWNED','RENTED');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE doc_type_enum AS ENUM ('AADHAAR','PAN','VOTER_ID','PASSPORT','DRIVING_LICENSE','ADDRESS_PROOF');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE verification_status_enum AS ENUM ('PENDING','VERIFIED','REJECTED');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE asset_type_enum AS ENUM ('LAPTOP','MOBILE','ACCESS_CARD','HEADSET','OTHER');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE asset_category_enum AS ENUM ('HARDWARE','SOFTWARE','PERIPHERAL');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE asset_condition_enum AS ENUM ('GOOD','DAMAGED','UNDER_REPAIR');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE acknowledgement_status_enum AS ENUM ('PENDING','ACKNOWLEDGED');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;


CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,

    employee_code TEXT UNIQUE NOT NULL,
    work_email TEXT UNIQUE NOT NULL,

    first_name TEXT NOT NULL,
    middle_name TEXT,
    last_name TEXT NOT NULL,
    display_name TEXT,

    gender gender_enum,
    dob DATE,
    marital_status marital_status_enum,
    blood_group TEXT,

    is_physically_handicapped BOOLEAN DEFAULT FALSE,
    disability_percentage FLOAT,

    nationality TEXT,

    personal_email TEXT UNIQUE,
    mobile_no TEXT,
    work_no TEXT,
    residence_no TEXT,

    designation TEXT,
    department_id INT,

    reporting_manager_id INT,
    dotted_line_manager_id INT,

    date_of_joining DATE,
    employment_type employment_type_enum,
    probation_end_date DATE,

    employment_context_role TEXT DEFAULT 'EMPLOYEE',

    relations JSONB,

    is_active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE IF NOT EXISTS employee_education (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    degree TEXT,
    branch TEXT,
    university TEXT,
    cgpa_or_pct FLOAT,

    year_of_joining INT,
    year_of_completion INT,

    certificate_url TEXT
);

CREATE INDEX IF NOT EXISTS idx_employee_education_employee_id
ON employee_education(employee_id);



CREATE TABLE IF NOT EXISTS employee_experience (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    company_name TEXT,
    designation TEXT,
    employment_type TEXT,

    start_date DATE,
    end_date DATE,

    is_current BOOLEAN DEFAULT FALSE,

    industry TEXT,
    description TEXT
);

CREATE INDEX IF NOT EXISTS idx_employee_experience_employee_id
ON employee_experience(employee_id);



CREATE TABLE IF NOT EXISTS employee_addresses (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    address_type address_type_enum,

    line1 TEXT,
    line2 TEXT,

    city TEXT,
    state TEXT,
    pin_code TEXT,
    country TEXT,

    ownership_type ownership_type_enum,

    UNIQUE(employee_id, address_type)
);

CREATE INDEX IF NOT EXISTS idx_employee_addresses_employee_id
ON employee_addresses(employee_id);


CREATE TABLE IF NOT EXISTS employee_identity_documents (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,

    doc_type doc_type_enum,
    doc_number TEXT,
    name_on_doc TEXT,

    issue_date DATE,
    expiry_date DATE,

    extra_info JSONB,

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

    file_size_kb INT,
    mime_type TEXT,

    verification_status verification_status_enum DEFAULT 'PENDING',

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

    asset_type asset_type_enum,
    asset_name TEXT,
    asset_category asset_category_enum,

    serial_no TEXT,

    assigned_on DATE,

    acknowledgement_status acknowledgement_status_enum DEFAULT 'PENDING',
    acknowledged_at TIMESTAMP,

    condition asset_condition_enum DEFAULT 'GOOD',

    assigned_by INT,
    notes TEXT,

    returned_on DATE,

    is_active BOOLEAN DEFAULT TRUE
);

CREATE INDEX IF NOT EXISTS idx_employee_assets_employee_id
ON employee_assets(employee_id);

CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,
    notification_type TEXT NOT NULL,
    title TEXT NOT NULL,
    message TEXT NOT NULL,
    reminder_day INT,
    metadata JSONB,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    read_at TIMESTAMP,
    UNIQUE(employee_id, notification_type, reminder_day)
);

CREATE INDEX IF NOT EXISTS idx_notifications_employee_id
ON notifications(employee_id);


ALTER TABLE employees ADD COLUMN IF NOT EXISTS password_hash TEXT;
ALTER TABLE employees ADD COLUMN IF NOT EXISTS refresh_token_hash TEXT;
ALTER TABLE employees ADD COLUMN IF NOT EXISTS refresh_token_expiry TIMESTAMP;
