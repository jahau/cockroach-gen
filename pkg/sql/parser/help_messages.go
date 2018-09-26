// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1073
	`ALTER`: {
		//line sql.y: 1074
		Category: hGroup,
		//line sql.y: 1075
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1089
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1090
		Category: hDDL,
		//line sql.y: 1091
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>
  ALTER PARTITION ... OF TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1126
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1140
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1141
		Category: hDDL,
		//line sql.y: 1142
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1144
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1151
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1152
		Category: hDDL,
		//line sql.y: 1153
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1176
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1177
		Category: hPriv,
		//line sql.y: 1178
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1180
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1185
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1186
		Category: hDDL,
		//line sql.y: 1187
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1189
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1197
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1198
		Category: hDDL,
		//line sql.y: 1199
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1210
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1215
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1216
		Category: hDDL,
		//line sql.y: 1217
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1225
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1599
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1600
		Category: hCCL,
		//line sql.y: 1601
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1618
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1626
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1627
		Category: hCCL,
		//line sql.y: 1628
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1644
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1662
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1663
		Category: hCCL,
		//line sql.y: 1664
		Text: `
IMPORT [ TABLE <tablename> FROM ]
       <format> ( <datafile> )
       [ WITH <option> [= <value>] [, ...] ]

IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   MYSQLOUTFILE
   MYSQLDUMP (mysqldump's SQL output)
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1690
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1719
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1720
		Category: hCCL,
		//line sql.y: 1721
		Text: `
EXPORT INTO <format> (<datafile> [WITH <option> [= value] [,...]]) FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1730
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1820
	`CANCEL`: {
		//line sql.y: 1821
		Category: hGroup,
		//line sql.y: 1822
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1829
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1830
		Category: hMisc,
		//line sql.y: 1831
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1834
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1852
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1853
		Category: hMisc,
		//line sql.y: 1854
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1857
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1888
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1889
		Category: hMisc,
		//line sql.y: 1890
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1893
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1940
	`CREATE`: {
		//line sql.y: 1941
		Category: hGroup,
		//line sql.y: 1942
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 1965
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic (experimental)`,
		//line sql.y: 1966
		Category: hExperimental,
		//line sql.y: 1967
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename>
`,
	},
	//line sql.y: 2025
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2026
		Category: hDML,
		//line sql.y: 2027
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2031
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2046
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2047
		Category: hCfg,
		//line sql.y: 2048
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2060
	`DROP`: {
		//line sql.y: 2061
		Category: hGroup,
		//line sql.y: 2062
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2078
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2079
		Category: hDDL,
		//line sql.y: 2080
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2081
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2093
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2094
		Category: hDDL,
		//line sql.y: 2095
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2096
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2108
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2109
		Category: hDDL,
		//line sql.y: 2110
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2111
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2123
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2124
		Category: hDDL,
		//line sql.y: 2125
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2126
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2146
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2147
		Category: hDDL,
		//line sql.y: 2148
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2149
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2169
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2170
		Category: hPriv,
		//line sql.y: 2171
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2172
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2184
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2185
		Category: hPriv,
		//line sql.y: 2186
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2187
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2209
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2210
		Category: hMisc,
		//line sql.y: 2211
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, VERBOSE

`,
		//line sql.y: 2223
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2285
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2286
		Category: hMisc,
		//line sql.y: 2287
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2288
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2310
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2311
		Category: hMisc,
		//line sql.y: 2312
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2313
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2336
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2337
		Category: hMisc,
		//line sql.y: 2338
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2339
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2359
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2360
		Category: hPriv,
		//line sql.y: 2361
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2374
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2390
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2391
		Category: hPriv,
		//line sql.y: 2392
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2405
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2460
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2461
		Category: hCfg,
		//line sql.y: 2462
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2463
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2475
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2476
		Category: hCfg,
		//line sql.y: 2477
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2478
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2487
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2488
		Category: hCfg,
		//line sql.y: 2489
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2492
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2512
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2513
		Category: hExperimental,
		//line sql.y: 2514
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2522
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2528
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2529
		Category: hExperimental,
		//line sql.y: 2530
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2538
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2546
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2547
		Category: hExperimental,
		//line sql.y: 2548
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 2559
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2614
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2615
		Category: hCfg,
		//line sql.y: 2616
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2617
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2638
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2639
		Category: hCfg,
		//line sql.y: 2640
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2646
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2663
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2664
		Category: hTxn,
		//line sql.y: 2665
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2672
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2840
	`SHOW`: {
		//line sql.y: 2841
		Category: hGroup,
		//line sql.y: 2842
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW JOBS,
SHOW QUERIES, SHOW ROLES, SHOW SESSION, SHOW SESSIONS, SHOW STATISTICS,
SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 2874
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2875
		Category: hCfg,
		//line sql.y: 2876
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2877
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2898
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 2899
		Category: hExperimental,
		//line sql.y: 2900
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 2907
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 2921
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 2922
		Category: hExperimental,
		//line sql.y: 2923
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 2927
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 2941
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2942
		Category: hCCL,
		//line sql.y: 2943
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 2944
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2971
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2972
		Category: hCfg,
		//line sql.y: 2973
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2976
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2993
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2994
		Category: hDDL,
		//line sql.y: 2995
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2996
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3004
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3005
		Category: hDDL,
		//line sql.y: 3006
		Text: `SHOW DATABASES
`,
		//line sql.y: 3007
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3015
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3016
		Category: hPriv,
		//line sql.y: 3017
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3023
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3036
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3037
		Category: hDDL,
		//line sql.y: 3038
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 3039
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3057
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3058
		Category: hDDL,
		//line sql.y: 3059
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3060
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3073
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3074
		Category: hMisc,
		//line sql.y: 3075
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3076
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3092
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3093
		Category: hMisc,
		//line sql.y: 3094
		Text: `SHOW JOBS
`,
		//line sql.y: 3095
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3103
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3104
		Category: hMisc,
		//line sql.y: 3105
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3107
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3130
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3131
		Category: hMisc,
		//line sql.y: 3132
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3133
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3149
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3150
		Category: hDDL,
		//line sql.y: 3151
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3152
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3178
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3179
		Category: hDDL,
		//line sql.y: 3180
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3192
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3193
		Category: hMisc,
		//line sql.y: 3194
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3203
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3204
		Category: hCfg,
		//line sql.y: 3205
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3206
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3225
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3226
		Category: hDDL,
		//line sql.y: 3227
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3228
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3246
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3247
		Category: hPriv,
		//line sql.y: 3248
		Text: `SHOW USERS
`,
		//line sql.y: 3249
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3257
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3258
		Category: hPriv,
		//line sql.y: 3259
		Text: `SHOW ROLES
`,
		//line sql.y: 3260
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3305
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3306
		Category: hMisc,
		//line sql.y: 3307
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3543
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3544
		Category: hMisc,
		//line sql.y: 3545
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3548
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3565
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3566
		Category: hDDL,
		//line sql.y: 3567
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3594
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4107
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4108
		Category: hDDL,
		//line sql.y: 4109
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4119
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4173
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4174
		Category: hDML,
		//line sql.y: 4175
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4176
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4184
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4185
		Category: hPriv,
		//line sql.y: 4186
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4187
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4209
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4210
		Category: hPriv,
		//line sql.y: 4211
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4212
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4224
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4225
		Category: hDDL,
		//line sql.y: 4226
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4227
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4257
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4258
		Category: hDDL,
		//line sql.y: 4259
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4267
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4471
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4472
		Category: hTxn,
		//line sql.y: 4473
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4474
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4482
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4483
		Category: hMisc,
		//line sql.y: 4484
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4487
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4504
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4505
		Category: hTxn,
		//line sql.y: 4506
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4507
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4522
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4523
		Category: hTxn,
		//line sql.y: 4524
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4532
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4545
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4546
		Category: hTxn,
		//line sql.y: 4547
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4550
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4574
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4575
		Category: hTxn,
		//line sql.y: 4576
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4577
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 4690
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 4691
		Category: hDDL,
		//line sql.y: 4692
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4693
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 4762
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 4763
		Category: hDML,
		//line sql.y: 4764
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4769
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 4788
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 4789
		Category: hDML,
		//line sql.y: 4790
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 4794
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 4899
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 4900
		Category: hDML,
		//line sql.y: 4901
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4908
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5078
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5079
		Category: hDML,
		//line sql.y: 5080
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5091
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5092
		Category: hDML,
		//line sql.y: 5093
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 5105
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5180
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5181
		Category: hDML,
		//line sql.y: 5182
		Text: `TABLE <tablename>
`,
		//line sql.y: 5183
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5452
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5453
		Category: hDML,
		//line sql.y: 5454
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5455
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5545
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5546
		Category: hDML,
		//line sql.y: 5547
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 5565
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
