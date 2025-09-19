-- 001_init.sql
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) UNIQUE NOT NULL,
  full_name VARCHAR(200),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS admins (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) UNIQUE NOT NULL,
  role VARCHAR(50),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- contoh data
INSERT INTO users (username, full_name) VALUES ('alice', 'Alice Example') ON CONFLICT DO NOTHING;
INSERT INTO admins (username, role) VALUES ('superadmin', 'super') ON CONFLICT DO NOTHING;
