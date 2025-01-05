CREATE TABLE IF NOT EXISTS two_factor_auth_number_options (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  two_factor_auth_session_id UUID NOT NULL,
  number VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_two_factor_auth_session
    FOREIGN KEY (two_factor_auth_session_id) REFERENCES two_factor_auth_sessions (id) ON DELETE CASCADE
);