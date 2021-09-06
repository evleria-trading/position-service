CREATE OR REPLACE FUNCTION notify_position ()
    RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$$
DECLARE
BEGIN
    IF (TG_OP = 'INSERT') THEN
        PERFORM pg_notify('notify_position_opened', row_to_json(NEW)::text);
        RETURN NEW;
    ELSIF (TG_OP = 'UPDATE' AND OLD.close_price IS NULL AND NEW.close_price IS NOT NULL) THEN
        PERFORM pg_notify('notify_position_closed', row_to_json(NEW)::text);
        RETURN NEW;
    ELSIF (TG_OP = 'UPDATE') THEN
        PERFORM pg_notify('notify_position_updated', row_to_json(NEW)::text);
        RETURN NEW;
    END IF;
END;
$$;

CREATE TRIGGER notify_position_trigger
    AFTER INSERT OR UPDATE
    ON positions
    FOR EACH ROW EXECUTE PROCEDURE notify_position();