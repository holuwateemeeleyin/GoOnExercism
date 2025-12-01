UPDATE grains
SET result = CASE
    WHEN task = 'single-square' THEN CAST(POWER(2, square - 1) AS NUMERIC)
    WHEN task = 'total' THEN CAST(POWER(2, 64) - 1 AS NUMERIC)
END;
