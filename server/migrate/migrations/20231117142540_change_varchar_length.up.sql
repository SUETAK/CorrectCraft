SET statement_timeout = 0;

--bun:split

ALTER TABLE gpt_describes ALTER COLUMN describe type varchar(1024);
