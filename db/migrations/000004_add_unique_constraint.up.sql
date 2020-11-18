ALTER TABLE accounts
    ADD CONSTRAINT unique_login UNIQUE (username, email);