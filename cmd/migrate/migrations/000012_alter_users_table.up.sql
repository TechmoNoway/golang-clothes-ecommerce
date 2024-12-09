ALTER TABLE
    users
ADD CONSTRAINT fk_user_roles FOREIGN KEY (role_id) REFERENCES roles(id);