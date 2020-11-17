
CREATE OR REPLACE FUNCTION calculate_lineage()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
AS
$$
DECLARE
    p_depth   integer;
    p_lineage varchar;
BEGIN
    IF NEW.parent_id <> OLD.parent_id THEN

        SELECT depth, lineage INTO p_depth, p_lineage FROM nodes WHERE id = NEW.parent_id;

        UPDATE nodes SET depth= p_depth + 1, lineage = p_lineage || NEW.id::varchar || '/' WHERE id = NEW.id;

    END IF;

    IF (TG_OP = 'INSERT') THEN
        SELECT depth, lineage INTO p_depth, p_lineage FROM nodes WHERE id = NEW.parent_id;

        UPDATE nodes SET depth= p_depth + 1, lineage = p_lineage || NEW.id::varchar || '/' WHERE id = NEW.id;

    END IF;

    RETURN NEW;
END
$$;


CREATE TRIGGER lineage_trigger
    AFTER UPDATE OR INSERT
    ON nodes
    FOR EACH ROW
EXECUTE PROCEDURE calculate_lineage();

