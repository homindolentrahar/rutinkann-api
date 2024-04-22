CREATE TABLE IF NOT EXISTS logs (
  id SERIAL PRIMARY KEY,
  description TEXT,
  count INT DEFAULT 0,
  completed_at TIMESTAMP DEFAULT NOW(),
  activity_id INT, 
  CONSTRAINT fk_actvity FOREIGN KEY(activity_id) REFERENCES activities(id)
);