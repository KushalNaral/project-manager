-- Create oauth_providers table to store information about the OAuth providers
CREATE TABLE oauth_providers (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE, -- Example: 'google', 'github', etc.
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create oauth_tokens table to store OAuth tokens for users
CREATE TABLE oauth_tokens (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL, -- Foreign key to the users table
    provider_id BIGINT UNSIGNED NOT NULL, -- Foreign key to oauth_providers
    access_token TEXT NOT NULL, -- The OAuth access token
    refresh_token TEXT, -- The OAuth refresh token
    expiry DATETIME NOT NULL, -- Expiry time of the access token
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (provider_id) REFERENCES oauth_providers(id) ON DELETE CASCADE
);

