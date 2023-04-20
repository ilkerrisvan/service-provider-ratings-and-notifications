CREATE TABLE Rating (
    id SERIAL PRIMARY KEY,
    service_provider_id INT NOT NULL,
    service_provier_rating INT NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_service_provider_id
ON Rating(service_provider_id);