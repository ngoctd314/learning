# Limits on Table Column Count and Row Size

## Column Count Limits

MySQL has hard limit of 4096 columns per table, but the effective maximum may be less for a given table. The exact column limit depends on several factors:

- The maximum row size for a table constrains the number (and possibly size) of columns because the total length of all columns cannot exceed this size.
- The storage requirements of individual columns constrain the number of columns that fit within a given maximum row size.
- Storage engines may impose additional restrictions that limit table column count. For example, InnoDB has a limit of 1017 columns per table.
- Each table has an .frm file that contains the table definition. The definition affects the content of this file in ways that may effect the number of columns permitted in the table.

## Row Size Limits

The maximum row size for a given table is determined by several factors:

- The internal representation of a MySQL table has a maximum row size limit of 65,535 bytes, even if the storage engine is capable of supporting larger rows. BLOB and TEXT columns only contribute 9 to 12 bytes toward the row size limit because their contents are stored separately from the rest of the row.
- The maximum row size for an InnoDB table, which applies to data stored locally within a database page, is slightly less than half a page for 4KB, 8KB, 16KB, and 32KB innodb_page_size settings. 
