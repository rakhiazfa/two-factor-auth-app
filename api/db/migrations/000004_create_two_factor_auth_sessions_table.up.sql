CREATE TABLE IF NOT EXISTS two_factor_auth_sessions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL,
  user_device_id UUID NOT NULL,
  approved_by UUID NULL,
  correct_number VARCHAR(100) NULL,
  verified BOOLEAN NOT NULL DEFAULT false,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  CONSTRAINT fk_user_device
    FOREIGN KEY (user_device_id) REFERENCES user_devices (id) ON DELETE CASCADE,
  CONSTRAINT fk_approver
    FOREIGN KEY (approved_by) REFERENCES user_devices (id) ON DELETE SET NULL
);