CREATE TABLE IF NOT EXISTS grievant_categories (
    id SERIAL,
    name VARCHAR (200) NULL,
    description TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievants (
    id SERIAL,
    grievant_category_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievants_groups (
    id SERIAL,
    name VARCHAR (200) NULL,
    description TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    grievant_category_id INT REFERENCES grievant_categories (id),
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievants_groups_has_grievants (
    grievant_group_id INT REFERENCES grievants_groups (id),
    grievant_id INT REFERENCES grievants (id),
    role VARCHAR (45) NULL,
    PRIMARY KEY (grievant_group_id, grievant_id)
);
CREATE TABLE IF NOT EXISTS grievance_filing_modes (
    id SERIAL,
    name VARCHAR (200) NULL,
    code_name VARCHAR (45) NULL,
    description TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievance_categories (
    id SERIAL,
    name VARCHAR (200) NULL,
    code_name VARCHAR (45) NULL,
    description TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievance_sub_categories (
    id SERIAL,
    name VARCHAR (200) NULL,
    code_name VARCHAR (45) NULL,
    description TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    grievance_category_id INT REFERENCES grievance_categories (id),
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievances (
    id SERIAL,
    location_occured VARCHAR (255) NULL,
    description TEXT NULL,
    comments TEXT NULL,
    state VARCHAR (200) NULL,
    grievance_filing_mode_id INT REFERENCES grievance_filing_modes (id),
    grievance_sub_category_id INT REFERENCES grievance_sub_categories (id),
    grievant_id INT REFERENCES grievants (id),
    grievant_group_id INT REFERENCES grievants_groups (id),
    reference_number VARCHAR (200) NOT NULL UNIQUE,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievance_appeals (
    id SERIAL,
    reference_number VARCHAR (200) NOT NULL UNIQUE,
    home_email VARCHAR (255) NULL,
    workplace_email VARCHAR (255) NULL,
    description TEXT NULL,
    grievance_reference_number VARCHAR (200) NOT NULL,
    desired_outcome TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS appeal_reasons (
    id SERIAL,
    name VARCHAR (200) NULL,
    description TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievance_appeals_has_appeal_reasons (
    grievance_appeal_id INT REFERENCES grievance_appeals (id),
    appeal_reason_id INT REFERENCES appeal_reasons (id),
    PRIMARY KEY (grievance_appeal_id, appeal_reason_id)
);
CREATE TABLE IF NOT EXISTS grievances_attachments (
    id SERIAL,
    attachable_id INT NULL,
    attachable_type VARCHAR (200) NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievance_states (
    id SERIAL,
    name VARCHAR (200) NULL,
    description TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    code_name VARCHAR (45) NULL,
    days INT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievances_has_states (
    id SERIAL,
    grievance_id INT REFERENCES grievances (id),
    grievance_state_id INT REFERENCES grievance_states (id),
    status VARCHAR (45) NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS grievance_state_actions (
    id SERIAL,
    name VARCHAR (45) NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    role_perform_action VARCHAR (45) NULL,
    grievance_state_id INT REFERENCES grievance_states (id),
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS state_transitions (
    id SERIAL,
    from_state_id INT NULL,
    to_state_id INT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS state_transtion_actions (
    id SERIAL,
    description TEXT NULL,
    created_at timestamp (0) NULL,
    updated_at timestamp (0) NULL,
    data text NULL,
    grievances_states_grievances_id INT REFERENCES grievances_has_states (id),
    user_id INT,
    PRIMARY KEY (id)
);