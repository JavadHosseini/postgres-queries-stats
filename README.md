## to use pg_stat_statements you should follow steps below:

1. CREATE EXTENSION pg_stat_statements;

in the database you want to monitor.

2. Additionally you will have to edit the following line in postgresql.conf

shared_preload_libraries = ''

To

shared_preload_libraries = 'pg_stat_statements'

3. A restart of the PSQL service is required.