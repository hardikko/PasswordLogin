BEGIN;

-- Generic "updated_at" trigger function
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$

BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

--Organization
CREATE TABLE "organizations" (
    "id" bigserial UNIQUE NOT NULL PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "website" varchar,
    "sector" varchar NOT NULL,
    "status" varchar NOT NULL DEFAULT '',
    "is_final" boolean NOT NULL DEFAULT FALSE,
    "is_archived" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON organizations
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

--Department
CREATE TABLE "departments" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "code" varchar UNIQUE NOT NULL,
    "org_id" bigint NOT NULL REFERENCES organizations (id),
    "name" varchar NOT NULL,
    "status" varchar NOT NULL DEFAULT '',
    "is_final" boolean NOT NULL DEFAULT FALSE,
    "is_archived" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON departments
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();
COMMIT;

--Roles
CREATE TABLE "roles" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "code" varchar UNIQUE NOT NULL,
    "org_id" bigint NOT NULL REFERENCES organizations (id),
    "depart_id" bigint NOT NULL REFERENCES departments (id),
    "name" varchar NOT NULL,
    "permissions" text[],
    "is_management" boolean NOT NULL DEFAULT FALSE,
    "status" varchar NOT NULL DEFAULT '',
    "is_final" boolean NOT NULL DEFAULT FALSE,
    "is_archived" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON roles
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

-- User
CREATE TABLE "users" (
    "id" bigserial UNIQUE NOT NULL PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "phone" varchar UNIQUE NOT NULL,
    "password_hash" varchar NOT NULL,
    "is_admin" boolean NOT NULL DEFAULT FALSE,
    "org_id" bigint REFERENCES organizations (id),
    "role_id" bigint REFERENCES roles (id),
    "status" varchar NOT NULL DEFAULT '',
    "is_final" boolean NOT NULL DEFAULT FALSE,
    "is_archived" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON  users
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

--Otp Session
CREATE TABLE "otp_sessions" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "user_id" bigint NOT NULL REFERENCES users (id),
    "token" varchar NOT NULL,
    "is_valid" boolean NOT NULL DEFAULT FALSE,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON otp_sessions
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

-- Authtoken
CREATE TABLE "auth_sessions" (
    "id" bigserial UNIQUE NOT NULL PRIMARY KEY,
    "user_id" bigint NOT NULL REFERENCES users (id),
    "token" uuid UNIQUE NOT NULL,
    "is_valid" boolean NOT NULL DEFAULT FALSE,
    "expire_at" timestamptz NOT NULL DEFAULT Now(),
    "created_at" timestamptz NOT NULL DEFAULT Now(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON  auth_sessions
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

--Leads
CREATE TABLE "leads" (
    "id" bigserial UNIQUE NOT NULL PRIMARY KEY,
    "user_id" bigint NOT NULL REFERENCES users (id),
    "firstname" varchar NOT NULL,
    "lastname" varchar NOT NULL,
    "address" varchar NOT NULL,
    "phone" varchar NOT NULL,
    "email" varchar NOT NULL,
    "occupation" varchar NOT NULL,
    "company" varchar NOT NULL,
    "status" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON leads
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();


END;