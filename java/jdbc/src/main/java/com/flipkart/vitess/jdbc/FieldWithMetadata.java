package com.flipkart.vitess.jdbc;

import java.sql.SQLException;
import java.sql.Types;
import java.util.regex.PatternSyntaxException;

import com.flipkart.vitess.util.CharsetMapping;
import com.flipkart.vitess.util.Constants;
import com.flipkart.vitess.util.MysqlDefs;
import com.flipkart.vitess.util.StringUtils;
import com.google.protobuf.ByteString;
import com.youtube.vitess.proto.Query;

public class FieldWithMetadata {

  private final VitessConnection connection;
  private final Query.Field field;
  private final boolean isImplicitTempTable;

  private int javaType;
  private int mysqlType;
  private int precisionAdjustFactor;
  private int colFlag;
  private String encoding;
  private String collationName;
  private int collationIndex;
  private int maxBytesPerChar;

  public FieldWithMetadata(VitessConnection connection, Query.Field field) throws SQLException {
    this.connection = connection;
    this.field = field;
    this.colFlag = field.getFlags();
    this.mysqlType = field.getMysqlType();
    this.collationIndex = field.getCharset();

    // Map MySqlTypes to java.sql Types
    if (MysqlDefs.vitessToJavaType.containsKey(field.getType())) {
      this.javaType = MysqlDefs.vitessToJavaType.get(field.getType());
    } else if (field.getType().equals(Query.Type.TUPLE)) {
      throw new SQLException(Constants.SQLExceptionMessages.INVALID_COLUMN_TYPE);
    } else {
      throw new SQLException(Constants.SQLExceptionMessages.UNKNOWN_COLUMN_TYPE);
    }

    this.isImplicitTempTable = field.getTable().length() > 5 && field.getTable().startsWith("#sql_");

    // Re-map to 'real' blob type, if we're a BLOB
    boolean isFromFunction = field.getOrgTable().isEmpty();

    if (mysqlType == MysqlDefs.FIELD_TYPE_BLOB) {
      if (connection.getBlobsAreStrings() || (connection.getFunctionsNeverReturnBlobs() && isFromFunction)) {
        this.javaType = Types.VARCHAR;
        this.mysqlType = MysqlDefs.FIELD_TYPE_VARCHAR;
      } else if (collationIndex == CharsetMapping.MYSQL_COLLATION_INDEX_binary) {
        if (connection.getUseBlobToStoreUTF8OutsideBMP() && shouldSetupForUtf8StringInBlob()) {
          if (this.getColumnLength() == MysqlDefs.LENGTH_TINYBLOB || this.getColumnLength() == MysqlDefs.LENGTH_BLOB) {
            this.mysqlType = MysqlDefs.FIELD_TYPE_VARCHAR;
            this.javaType = Types.VARCHAR;
          } else {
            this.mysqlType = MysqlDefs.FIELD_TYPE_VAR_STRING;
            this.javaType = Types.LONGVARCHAR;
          }

          this.collationIndex = CharsetMapping.MYSQL_COLLATION_INDEX_utf8;
        } else {
          if (this.getColumnLength() == MysqlDefs.LENGTH_TINYBLOB) {
            this.mysqlType = MysqlDefs.FIELD_TYPE_TINY_BLOB;
          } else if (this.getColumnLength() == MysqlDefs.LENGTH_BLOB) {
            this.mysqlType = MysqlDefs.FIELD_TYPE_BLOB;
          } else if (this.getColumnLength() == MysqlDefs.LENGTH_MEDIUMBLOB) {
            this.mysqlType = MysqlDefs.FIELD_TYPE_MEDIUM_BLOB;
          } else if (this.getColumnLength() == MysqlDefs.LENGTH_LONGBLOB) {
            this.mysqlType = MysqlDefs.FIELD_TYPE_LONG_BLOB;
          }
          this.javaType = MysqlDefs.mysqlToJavaType(mysqlType);
        }
      } else {
        // *TEXT masquerading as blob
        this.mysqlType = MysqlDefs.FIELD_TYPE_VAR_STRING;
        this.javaType = Types.LONGVARCHAR;
      }
    }

    if (this.javaType == Types.TINYINT && this.field.getColumnLength() == 1 && connection.getTinyInt1isBit()) {
      // Adjust for pseudo-boolean
      if (connection.getTinyInt1isBit()) {
        if (connection.getTransformedBitIsBoolean()) {
          this.javaType = Types.BOOLEAN;
        } else {
          this.javaType = Types.BIT;
        }
      }
    }

    if (!isNativeNumericType() && !isNativeDateTimeType()) {
      this.encoding = connection.getEncodingForIndex(this.collationIndex);

      // ucs2, utf16, and utf32 cannot be used as a client character set, but if it was received from server under some circumstances we can parse them as
      // utf16
      if ("UnicodeBig".equals(this.encoding)) {
        this.encoding = "UTF-16";
      }

      // MySQL encodes JSON data with utf8mb4.
      if (mysqlType == MysqlDefs.FIELD_TYPE_JSON) {
        this.encoding = "UTF-8";
      }

      // Handle VARBINARY/BINARY (server doesn't have a different type for this

      boolean isBinary = isBinary();

      if (mysqlType == MysqlDefs.FIELD_TYPE_VAR_STRING && isBinary && collationIndex == CharsetMapping.MYSQL_COLLATION_INDEX_binary) {
        if (connection.getFunctionsNeverReturnBlobs() && isFromFunction) {
          this.javaType = Types.VARCHAR;
          this.mysqlType = MysqlDefs.FIELD_TYPE_VARCHAR;
        } else if (this.isOpaqueBinary()) {
          this.javaType = Types.VARBINARY;
        }
      }

      if (mysqlType == MysqlDefs.FIELD_TYPE_STRING && isBinary && collationIndex == CharsetMapping.MYSQL_COLLATION_INDEX_binary) {
        //
        // Okay, this is a hack, but there's currently no way to easily distinguish something like DATE_FORMAT( ..) from the "BINARY" column type, other
        // than looking at the original column name.
        //

        if (isOpaqueBinary() && !connection.getBlobsAreStrings()) {
          this.javaType = Types.BINARY;
        }
      }

      if (mysqlType == MysqlDefs.FIELD_TYPE_BIT) {
        boolean isSingleBit = field.getColumnLength() == 0 || field.getColumnLength() == 1;

        if (!isSingleBit) {
          this.colFlag |= 128; // Pretend this is a full binary(128) and blob(16) so that this field is de-serializable.
          this.colFlag |= 16;
          isBinary = true;
        }
      }

      //
      // Handle TEXT type (special case)
      //
      if (javaType == java.sql.Types.LONGVARBINARY && !isBinary) {
        this.javaType = java.sql.Types.LONGVARCHAR;
      } else if (javaType == java.sql.Types.VARBINARY && !isBinary) {
        this.javaType = java.sql.Types.VARCHAR;
      }
    } else {
      this.encoding = "US-ASCII";
    }

    //
    // Handle odd values for 'M' for floating point/decimal numbers
    //
    this.precisionAdjustFactor = 0;
    if (!isUnsigned()) {
      switch (mysqlType) {
        case MysqlDefs.FIELD_TYPE_DECIMAL:
        case MysqlDefs.FIELD_TYPE_NEW_DECIMAL:
          this.precisionAdjustFactor = -1;

          break;
        case MysqlDefs.FIELD_TYPE_DOUBLE:
        case MysqlDefs.FIELD_TYPE_FLOAT:
          this.precisionAdjustFactor = 1;

          break;
      }
    } else {
      switch (mysqlType) {
        case MysqlDefs.FIELD_TYPE_DOUBLE:
        case MysqlDefs.FIELD_TYPE_FLOAT:
          this.precisionAdjustFactor = 1;

          break;
      }
    }

    //this.mysqlType = mysqlType;
    //this.javaType = javaType;
    //this.colFlag = colFlag;
    //this.encoding = encoding;
    //this.collationIndex = collationIndex;
    //this.isImplicitTempTable = isImplicitTempTable;
    //this.precisionAdjustFactor = precisionAdjustFactor;
  }

  private boolean isNativeNumericType() {
    return ((this.mysqlType >= MysqlDefs.FIELD_TYPE_TINY && this.mysqlType <= MysqlDefs.FIELD_TYPE_DOUBLE)
        || this.mysqlType == MysqlDefs.FIELD_TYPE_LONGLONG || this.mysqlType == MysqlDefs.FIELD_TYPE_YEAR);
  }

  private boolean isNativeDateTimeType() {
    return (this.mysqlType == MysqlDefs.FIELD_TYPE_DATE || this.mysqlType == MysqlDefs.FIELD_TYPE_NEWDATE || this.mysqlType == MysqlDefs.FIELD_TYPE_DATETIME);
  }

  private boolean shouldSetupForUtf8StringInBlob() throws SQLException {
    String includePattern = connection.getUtf8OutsideBmpIncludedColumnNamePattern();
    String excludePattern = connection.getUtf8OutsideBmpExcludedColumnNamePattern();

    if (excludePattern != null && !StringUtils.isNullOrEmptyWithoutWS(excludePattern)) {
      try {
        if (getOrgName().matches(excludePattern)) {
          if (includePattern != null && !StringUtils.isNullOrEmptyWithoutWS(includePattern)) {
            try {
              if (getOrgName().matches(includePattern)) {
                return true;
              }
            } catch (PatternSyntaxException pse) {
              SQLException sqlEx = new SQLException("Illegal regex specified for \"utf8OutsideBmpIncludedColumnNamePattern\"",
                  Constants.SQLExceptionMessages.ILLEGAL_VALUE_FOR);

              if (!connection.getParanoid()) {
                sqlEx.initCause(pse);
              }

              throw sqlEx;
            }
          }

          return false;
        }
      } catch (PatternSyntaxException pse) {
        SQLException sqlEx = new SQLException("Illegal regex specified for \"utf8OutsideBmpExcludedColumnNamePattern\"",
            Constants.SQLExceptionMessages.ILLEGAL_VALUE_FOR);

        if (!connection.getParanoid()) {
          sqlEx.initCause(pse);
        }

        throw sqlEx;
      }
    }

    return true;
  }

  public boolean isAutoIncrement() throws SQLException {
    return ((this.colFlag & Query.MySqlFlag.AUTO_INCREMENT_FLAG_VALUE) > 0);
  }

  public boolean isBinary() {
    return ((this.colFlag & Query.MySqlFlag.BINARY_FLAG_VALUE) > 0);
  }

  public boolean isBlob() {
    return ((this.colFlag & Query.MySqlFlag.BLOB_FLAG_VALUE) > 0);
  }

  public boolean isMultipleKey() {
    return ((this.colFlag & Query.MySqlFlag.MULTIPLE_KEY_FLAG_VALUE) > 0);
  }

  boolean isNotNull() {
    return ((this.colFlag & Query.MySqlFlag.NOT_NULL_FLAG_VALUE) > 0);
  }

  public boolean isZeroFill() {
    return ((this.colFlag & Query.MySqlFlag.ZEROFILL_FLAG_VALUE) > 0);
  }

  boolean isOpaqueBinary() throws SQLException {

    //
    // Detect CHAR(n) CHARACTER SET BINARY which is a synonym for fixed-length binary types
    //

    if (this.collationIndex == CharsetMapping.MYSQL_COLLATION_INDEX_binary && isBinary()
        && (this.getMysqlType() == MysqlDefs.FIELD_TYPE_STRING || this.getMysqlType() == MysqlDefs.FIELD_TYPE_VAR_STRING)) {
      // Okay, queries resolved by temp tables also have this 'signature', check for that

      return !isImplicitTemporaryTable();
    }

    return "binary".equalsIgnoreCase(getEncoding());

  }

  public boolean isPrimaryKey() {
    return ((this.colFlag & 2) > 0);
  }

  /**
   * Is this field _definitely_ not writable?
   *
   * @return true if this field can not be written to in an INSERT/UPDATE
   *         statement.
   */
  boolean isReadOnly() throws SQLException {
    String orgColumnName = getOrgName();
    String orgTableName = getOrgTable();

    return !(orgColumnName != null && orgColumnName.length() > 0 && orgTableName != null && orgTableName.length() > 0);
  }

  public String getCollation() throws SQLException {
    if (this.collationName == null) {
      try {
        this.collationName = CharsetMapping.COLLATION_INDEX_TO_COLLATION_NAME[this.collationIndex];
      } catch (RuntimeException ex) {
        throw new SQLException(ex.toString(), Constants.SQLExceptionMessages.ILLEGAL_VALUE_FOR, ex);
      }
    }

    return this.collationName;
  }

  public synchronized int getMaxBytesPerCharacter() throws SQLException {
    if (this.maxBytesPerChar == 0) {
      this.maxBytesPerChar = this.connection.getMaxBytesPerChar(this.collationIndex, getEncoding());
    }
    return this.maxBytesPerChar;
  }

  public boolean isUniqueKey() {
    return ((this.colFlag & 4) > 0);
  }

  public boolean isUnsigned() {
    return ((this.colFlag & 32) > 0);
  }

  public String getName() {
    return field.getName();
  }

  public ByteString getNameBytes() {
    return field.getNameBytes();
  }

  public int getTypeValue() {
    return field.getTypeValue();
  }

  public Query.Type getType() {
    return field.getType();
  }

  public String getTable() {
    return field.getTable();
  }

  public ByteString getTableBytes() {
    return field.getTableBytes();
  }

  public String getSchema() {
    return field.getSchema();
  }

  public ByteString getSchemaBytes() {
    return field.getSchemaBytes();
  }

  public String getOrgTable() {
    return field.getOrgTable();
  }

  public ByteString getOrgTableBytes() {
    return field.getOrgTableBytes();
  }

  public String getDatabase() {
    return field.getDatabase();
  }

  public ByteString getDatabaseBytes() {
    return field.getDatabaseBytes();
  }

  public String getOrgName() {
    return field.getOrgName();
  }

  public ByteString getOrgNameBytes() {
    return field.getOrgNameBytes();
  }

  public int getColumnLength() {
    return field.getColumnLength();
  }

  public long getCharset() {
    return field.getCharset();
  }

  public int getDecimals() {
    return field.getDecimals();
  }

  public int getFlags() {
    return field.getFlags();
  }

  @Override
  public String toString() {
    try {
      StringBuilder asString = new StringBuilder();
      asString.append(super.toString());
      asString.append("[");
      asString.append("catalog=");
      asString.append(this.getDatabase());
      asString.append(",tableName=");
      asString.append(this.getTable());
      asString.append(",originalTableName=");
      asString.append(this.getOrgTable());
      asString.append(",columnName=");
      asString.append(this.getName());
      asString.append(",originalColumnName=");
      asString.append(this.getOrgName());
      asString.append(",mysqlType=");
      asString.append(getMysqlType());
      asString.append("(");
      asString.append(MysqlDefs.typeToName(getMysqlType()));
      asString.append(")");
      asString.append(",flags=");

      if (isAutoIncrement()) {
        asString.append(" AUTO_INCREMENT");
      }

      if (isPrimaryKey()) {
        asString.append(" PRIMARY_KEY");
      }

      if (isUniqueKey()) {
        asString.append(" UNIQUE_KEY");
      }

      if (isBinary()) {
        asString.append(" BINARY");
      }

      if (isBlob()) {
        asString.append(" BLOB");
      }

      if (isMultipleKey()) {
        asString.append(" MULTI_KEY");
      }

      if (isUnsigned()) {
        asString.append(" UNSIGNED");
      }

      if (isZeroFill()) {
        asString.append(" ZEROFILL");
      }

      asString.append(", charsetIndex=");
      asString.append(this.collationIndex);
      asString.append(", charsetName=");
      asString.append(this.encoding);

      //if (this.buffer != null) {
      //	asString.append("\n\nData as received from server:\n\n");
      //	asString.append(StringUtils.dumpAsHex(this.buffer,
      //			this.buffer.length));
      //}

      asString.append("]");

      return asString.toString();
    } catch (Throwable t) {
      return super.toString();
    }
  }

  public int getJavaType() {
    return javaType;
  }

  public int getMysqlType() {
    return mysqlType;
  }

  public boolean isImplicitTemporaryTable() {
    return isImplicitTempTable;
  }

  public String getEncoding() {
    return encoding;
  }

  public int getPrecisionAdjustFactor() {
    return precisionAdjustFactor;
  }
}
