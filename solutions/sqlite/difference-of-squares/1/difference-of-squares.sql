-- Schema: CREATE TABLE "difference-of-squares" ("number" INT, "property" TEXT, "result" INT);
-- Task: update the difference-of-squares table and set the result based on the number and property fields.
UPDATE "difference-of-squares"
SET result = CASE
    WHEN LOWER(property) IN ('squareofsum', 'square_of_sum') THEN
        (number * (number + 1) / 2) * (number * (number + 1) / 2)

    WHEN LOWER(property) IN ('sumofsquares', 'sum_of_squares') THEN
        (number * (number + 1) * (2 * number + 1) / 6)

    WHEN LOWER(property) IN ('difference', 'diff') THEN
        (number * (number + 1) / 2) * (number * (number + 1) / 2)
        - (number * (number + 1) * (2 * number + 1) / 6)

    -- Safety fallback: compute the difference
    ELSE
        (number * (number + 1) / 2) * (number * (number + 1) / 2)
        - (number * (number + 1) * (2 * number + 1) / 6)
END;
