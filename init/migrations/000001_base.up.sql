SET TIME ZONE 'UTC';

CREATE OR REPLACE FUNCTION update_edit_time()
RETURNS TRIGGER AS $$
BEGIN
   NEW.edited = now(); 
   RETURN NEW;
END;
$$ language 'plpgsql';
