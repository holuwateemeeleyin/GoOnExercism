-- Schema:
-- CREATE TABLE "bottle-song" (
--         start_bottles INTEGER NOT NULL,
--         take_down     INTEGER NOT NULL,
--         result        TEXT
-- );
-- Task: update bottle-song table and set the result based on the
-- start_bottles and take_down.
UPDATE "bottle-song"
SET result = (
    WITH RECURSIVE verses(n, text) AS (
        -- initial row
        SELECT start_bottles,
        CASE 
            WHEN start_bottles = 1 THEN 'One green bottle hanging on the wall,' || CHAR(10) ||
                                      'One green bottle hanging on the wall,' || CHAR(10) ||
                                      'And if one green bottle should accidentally fall,' || CHAR(10) ||
                                      'There''ll be no green bottles hanging on the wall.'
            ELSE
                (CASE start_bottles
                    WHEN 10 THEN 'Ten' WHEN 9 THEN 'Nine' WHEN 8 THEN 'Eight'
                    WHEN 7 THEN 'Seven' WHEN 6 THEN 'Six' WHEN 5 THEN 'Five'
                    WHEN 4 THEN 'Four'  WHEN 3 THEN 'Three' WHEN 2 THEN 'Two'
                END) || ' green bottles hanging on the wall,' || CHAR(10) ||

                (CASE start_bottles
                    WHEN 10 THEN 'Ten' WHEN 9 THEN 'Nine' WHEN 8 THEN 'Eight'
                    WHEN 7 THEN 'Seven' WHEN 6 THEN 'Six' WHEN 5 THEN 'Five'
                    WHEN 4 THEN 'Four'  WHEN 3 THEN 'Three' WHEN 2 THEN 'Two'
                END) || ' green bottles hanging on the wall,' || CHAR(10) ||

                'And if one green bottle should accidentally fall,' || CHAR(10) ||

                'There''ll be ' ||
                CASE start_bottles - 1
                    WHEN 10 THEN 'ten green bottles hanging on the wall.'
                    WHEN 9 THEN 'nine green bottles hanging on the wall.'
                    WHEN 8 THEN 'eight green bottles hanging on the wall.'
                    WHEN 7 THEN 'seven green bottles hanging on the wall.'
                    WHEN 6 THEN 'six green bottles hanging on the wall.'
                    WHEN 5 THEN 'five green bottles hanging on the wall.'
                    WHEN 4 THEN 'four green bottles hanging on the wall.'
                    WHEN 3 THEN 'three green bottles hanging on the wall.'
                    WHEN 2 THEN 'two green bottles hanging on the wall.'
                    WHEN 1 THEN 'one green bottle hanging on the wall.'
                    WHEN 0 THEN 'no green bottles hanging on the wall.'
                END
        END

        UNION ALL

        -- recursive row: generate next verse
        SELECT n - 1,
        CHAR(10) || CHAR(10) ||  -- blank line between verses
        CASE 
            WHEN n - 1 = 1 THEN 'One green bottle hanging on the wall,' || CHAR(10) ||
                              'One green bottle hanging on the wall,' || CHAR(10) ||
                              'And if one green bottle should accidentally fall,' || CHAR(10) ||
                              'There''ll be no green bottles hanging on the wall.'
            ELSE
                (CASE n - 1
                    WHEN 10 THEN 'Ten' WHEN 9 THEN 'Nine' WHEN 8 THEN 'Eight'
                    WHEN 7 THEN 'Seven' WHEN 6 THEN 'Six' WHEN 5 THEN 'Five'
                    WHEN 4 THEN 'Four'  WHEN 3 THEN 'Three' WHEN 2 THEN 'Two'
                END) 
                || ' green bottles hanging on the wall,' || CHAR(10) ||
                (CASE n - 1
                    WHEN 10 THEN 'Ten' WHEN 9 THEN 'Nine' WHEN 8 THEN 'Eight'
                    WHEN 7 THEN 'Seven' WHEN 6 THEN 'Six' WHEN 5 THEN 'Five'
                    WHEN 4 THEN 'Four'  WHEN 3 THEN 'Three' WHEN 2 THEN 'Two'
                END)
                || ' green bottles hanging on the wall,' || CHAR(10) ||
                'And if one green bottle should accidentally fall,' || CHAR(10) ||
                'There''ll be ' ||
                CASE n - 2
                    WHEN 10 THEN 'ten green bottles hanging on the wall.'
                    WHEN 9 THEN 'nine green bottles hanging on the wall.'
                    WHEN 8 THEN 'eight green bottles hanging on the wall.'
                    WHEN 7 THEN 'seven green bottles hanging on the wall.'
                    WHEN 6 THEN 'six green bottles hanging on the wall.'
                    WHEN 5 THEN 'five green bottles hanging on the wall.'
                    WHEN 4 THEN 'four green bottles hanging on the wall.'
                    WHEN 3 THEN 'three green bottles hanging on the wall.'
                    WHEN 2 THEN 'two green bottles hanging on the wall.'
                    WHEN 1 THEN 'one green bottle hanging on the wall.'
                    WHEN 0 THEN 'no green bottles hanging on the wall.'
                END
        END
        FROM verses
        WHERE start_bottles - n + 1 < take_down
    )
    SELECT group_concat(text, '')
    FROM verses
);
