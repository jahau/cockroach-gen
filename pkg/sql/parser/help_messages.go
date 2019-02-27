// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1063
	`ALTER`: {
		//line sql.y: 1064
		Category: hGroup,
		//line sql.y: 1065
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1079
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1080
		Category: hDDL,
		//line sql.y: 1081
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
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1117
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1131
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1132
		Category: hDDL,
		//line sql.y: 1133
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1135
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1142
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1143
		Category: hDDL,
		//line sql.y: 1144
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
	//line sql.y: 1169
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1170
		Category: hPriv,
		//line sql.y: 1171
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1173
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1178
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1179
		Category: hDDL,
		//line sql.y: 1180
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1182
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1190
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1191
		Category: hDDL,
		//line sql.y: 1192
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1204
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1209
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1210
		Category: hDDL,
		//line sql.y: 1211
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1219
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1610
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1611
		Category: hCCL,
		//line sql.y: 1612
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
		//line sql.y: 1629
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1637
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1638
		Category: hCCL,
		//line sql.y: 1639
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
		//line sql.y: 1655
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1673
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1674
		Category: hCCL,
		//line sql.y: 1675
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
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
		//line sql.y: 1703
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1737
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1738
		Category: hCCL,
		//line sql.y: 1739
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1748
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1841
	`CANCEL`: {
		//line sql.y: 1842
		Category: hGroup,
		//line sql.y: 1843
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1850
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1851
		Category: hMisc,
		//line sql.y: 1852
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1855
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1873
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1874
		Category: hMisc,
		//line sql.y: 1875
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1878
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1909
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1910
		Category: hMisc,
		//line sql.y: 1911
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1914
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1987
	`CREATE`: {
		//line sql.y: 1988
		Category: hGroup,
		//line sql.y: 1989
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2070
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic (experimental)`,
		//line sql.y: 2071
		Category: hExperimental,
		//line sql.y: 2072
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2162
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2163
		Category: hDML,
		//line sql.y: 2164
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2168
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2183
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2184
		Category: hCfg,
		//line sql.y: 2185
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2197
	`DROP`: {
		//line sql.y: 2198
		Category: hGroup,
		//line sql.y: 2199
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2216
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2217
		Category: hDDL,
		//line sql.y: 2218
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2219
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2231
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2232
		Category: hDDL,
		//line sql.y: 2233
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2234
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2246
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2247
		Category: hDDL,
		//line sql.y: 2248
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2249
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2261
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2262
		Category: hDDL,
		//line sql.y: 2263
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2264
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2284
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2285
		Category: hDDL,
		//line sql.y: 2286
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2287
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2307
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2308
		Category: hPriv,
		//line sql.y: 2309
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2310
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2322
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2323
		Category: hPriv,
		//line sql.y: 2324
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2325
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2349
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2350
		Category: hMisc,
		//line sql.y: 2351
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2364
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2424
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2425
		Category: hMisc,
		//line sql.y: 2426
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2427
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2458
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2459
		Category: hMisc,
		//line sql.y: 2460
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2461
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2491
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2492
		Category: hMisc,
		//line sql.y: 2493
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2494
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2514
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2515
		Category: hPriv,
		//line sql.y: 2516
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
		//line sql.y: 2529
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2545
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2546
		Category: hPriv,
		//line sql.y: 2547
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
		//line sql.y: 2560
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2615
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2616
		Category: hCfg,
		//line sql.y: 2617
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2618
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2630
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2631
		Category: hCfg,
		//line sql.y: 2632
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2633
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2642
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2643
		Category: hCfg,
		//line sql.y: 2644
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2647
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2668
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2669
		Category: hExperimental,
		//line sql.y: 2670
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2678
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2684
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2685
		Category: hExperimental,
		//line sql.y: 2686
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2694
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2702
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2703
		Category: hExperimental,
		//line sql.y: 2704
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
		//line sql.y: 2715
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2771
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2772
		Category: hCfg,
		//line sql.y: 2773
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2774
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2795
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2796
		Category: hCfg,
		//line sql.y: 2797
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2803
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2820
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2821
		Category: hTxn,
		//line sql.y: 2822
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2829
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3012
	`SHOW`: {
		//line sql.y: 3013
		Category: hGroup,
		//line sql.y: 3014
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
JOBS, SHOW QUERIES, SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW
SESSION, SHOW SESSIONS, SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES,
SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3048
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3049
		Category: hCfg,
		//line sql.y: 3050
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3051
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3072
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3073
		Category: hExperimental,
		//line sql.y: 3074
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3081
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3096
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3097
		Category: hExperimental,
		//line sql.y: 3098
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3102
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3116
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3117
		Category: hCCL,
		//line sql.y: 3118
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 3119
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3146
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3147
		Category: hCfg,
		//line sql.y: 3148
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 3151
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3168
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3169
		Category: hDDL,
		//line sql.y: 3170
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3171
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3180
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3181
		Category: hDDL,
		//line sql.y: 3182
		Text: `SHOW DATABASES
`,
		//line sql.y: 3183
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3191
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3192
		Category: hPriv,
		//line sql.y: 3193
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3199
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3212
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3213
		Category: hDDL,
		//line sql.y: 3214
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 3215
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3236
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3237
		Category: hDDL,
		//line sql.y: 3238
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3239
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3254
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3255
		Category: hMisc,
		//line sql.y: 3256
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3257
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3273
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3274
		Category: hMisc,
		//line sql.y: 3275
		Text: `SHOW JOBS
`,
		//line sql.y: 3276
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3284
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3285
		Category: hMisc,
		//line sql.y: 3286
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3288
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3311
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3312
		Category: hMisc,
		//line sql.y: 3313
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3314
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3330
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3331
		Category: hDDL,
		//line sql.y: 3332
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3333
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3365
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3366
		Category: hDDL,
		//line sql.y: 3367
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3379
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3380
		Category: hDDL,
		//line sql.y: 3381
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3393
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3394
		Category: hMisc,
		//line sql.y: 3395
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3404
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3405
		Category: hCfg,
		//line sql.y: 3406
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3407
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3426
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3427
		Category: hDDL,
		//line sql.y: 3428
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3429
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3449
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3450
		Category: hPriv,
		//line sql.y: 3451
		Text: `SHOW USERS
`,
		//line sql.y: 3452
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3460
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3461
		Category: hPriv,
		//line sql.y: 3462
		Text: `SHOW ROLES
`,
		//line sql.y: 3463
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3510
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3511
		Category: hMisc,
		//line sql.y: 3512
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3750
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3751
		Category: hMisc,
		//line sql.y: 3752
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3755
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3772
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3773
		Category: hDDL,
		//line sql.y: 3774
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
		//line sql.y: 3801
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4389
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4390
		Category: hDDL,
		//line sql.y: 4391
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
		//line sql.y: 4401
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4448
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4449
		Category: hDML,
		//line sql.y: 4450
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4451
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4459
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4460
		Category: hPriv,
		//line sql.y: 4461
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4462
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4484
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4485
		Category: hPriv,
		//line sql.y: 4486
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4487
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4505
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4506
		Category: hDDL,
		//line sql.y: 4507
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4508
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4542
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4543
		Category: hDDL,
		//line sql.y: 4544
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4552
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4794
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4795
		Category: hTxn,
		//line sql.y: 4796
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4797
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4805
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4806
		Category: hMisc,
		//line sql.y: 4807
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4810
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4827
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4828
		Category: hTxn,
		//line sql.y: 4829
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4830
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4845
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4846
		Category: hTxn,
		//line sql.y: 4847
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4855
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4868
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4869
		Category: hTxn,
		//line sql.y: 4870
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4873
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4897
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4898
		Category: hTxn,
		//line sql.y: 4899
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4900
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5018
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5019
		Category: hDDL,
		//line sql.y: 5020
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5021
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5090
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5091
		Category: hDML,
		//line sql.y: 5092
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5097
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5116
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5117
		Category: hDML,
		//line sql.y: 5118
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5122
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5229
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5230
		Category: hDML,
		//line sql.y: 5231
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5238
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5412
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5413
		Category: hDML,
		//line sql.y: 5414
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5425
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5426
		Category: hDML,
		//line sql.y: 5427
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
		//line sql.y: 5439
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5514
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5515
		Category: hDML,
		//line sql.y: 5516
		Text: `TABLE <tablename>
`,
		//line sql.y: 5517
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5797
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5798
		Category: hDML,
		//line sql.y: 5799
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5800
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5905
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5906
		Category: hDML,
		//line sql.y: 5907
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexflags> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> [ <jointype> ] JOIN <source> ON <expr>
  <source> [ <jointype> ] JOIN <source> USING ( <colnames...> )
  <source> NATURAL [ <jointype> ] JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

Join types:
  { INNER | { LEFT | RIGHT | FULL } [OUTER] } [ { HASH | MERGE | LOOKUP } ]

`,
		//line sql.y: 5928
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
