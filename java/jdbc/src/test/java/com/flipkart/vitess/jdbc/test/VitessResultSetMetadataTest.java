package com.flipkart.vitess.jdbc.test;

import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;
import java.util.Properties;

import org.junit.Assert;
import org.junit.BeforeClass;
import org.junit.Test;

import com.flipkart.vitess.jdbc.VitessConnection;
import com.flipkart.vitess.jdbc.VitessResultSetMetaData;
import com.flipkart.vitess.util.Constants;
import com.flipkart.vitess.util.MysqlDefs;
import com.youtube.vitess.proto.Query;

/**
 * Created by ashudeep.sharma on 08/02/16.
 */
public class VitessResultSetMetadataTest {

    private List<Query.Field> fieldList;

    String dbURL = "jdbc:vitess://locahost:9000/vt_shipment/shipment";

    @BeforeClass
    public static void setUp() {
        // load Vitess driver
        try {
            Class.forName("com.flipkart.vitess.jdbc.VitessDriver");
        } catch (ClassNotFoundException e) {
            Assert.fail("Driver is not in the CLASSPATH -> " + e.getMessage());
        }
    }

    private VitessConnection getVitessConnection() throws SQLException {
        return new VitessConnection(dbURL, new Properties());
    }

    public void initializeFieldList() {

        fieldList = new ArrayList<>();
        fieldList.add(Query.Field.newBuilder().setName("col1").setTable("tbl").setType(Query.Type.INT8).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col2").setTable("tbl").setType(Query.Type.UINT8).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setFlags(Query.MySqlFlag.UNSIGNED_FLAG_VALUE).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col3").setTable("tbl").setType(Query.Type.INT16).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col4").setTable("tbl").setType(Query.Type.UINT16).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setFlags(Query.MySqlFlag.UNSIGNED_FLAG_VALUE).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col5").setTable("tbl").setType(Query.Type.INT24).setMysqlType(MysqlDefs.FIELD_TYPE_INT24).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col6").setTable("tbl").setType(Query.Type.UINT24).setMysqlType(MysqlDefs.FIELD_TYPE_INT24).setFlags(Query.MySqlFlag.UNSIGNED_FLAG_VALUE).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col7").setTable("tbl").setType(Query.Type.INT32).setMysqlType(MysqlDefs.FIELD_TYPE_INT24).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col8").setTable("tbl").setType(Query.Type.UINT32).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setFlags(Query.MySqlFlag.UNSIGNED_FLAG_VALUE).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col9").setTable("tbl").setType(Query.Type.INT64).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setColumnLength(3).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col10").setTable("tbl").setType(Query.Type.UINT64).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setFlags(Query.MySqlFlag.UNSIGNED_FLAG_VALUE).setColumnLength(3).setCharset(33).build());
        fieldList
            .add(Query.Field.newBuilder().setName("col11").setTable("tbl").setType(Query.Type.FLOAT32).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setColumnLength(3).setCharset(33).build());
        fieldList
            .add(Query.Field.newBuilder().setName("col12").setTable("tbl").setType(Query.Type.FLOAT64).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setColumnLength(3).setCharset(33).build());
        fieldList
            .add(Query.Field.newBuilder().setName("col13").setTable("tbl").setType(Query.Type.TIMESTAMP).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col14").setTable("tbl").setType(Query.Type.DATE).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col15").setTable("tbl").setType(Query.Type.TIME).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList
            .add(Query.Field.newBuilder().setName("col16").setTable("tbl").setType(Query.Type.DATETIME).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col17").setTable("tbl").setType(Query.Type.YEAR).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList
            .add(Query.Field.newBuilder().setName("col18").setTable("tbl").setType(Query.Type.DECIMAL).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setColumnLength(3).setDecimals(2).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col19").setTable("tbl").setType(Query.Type.TEXT).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col20").setTable("tbl").setType(Query.Type.BLOB).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList
            .add(Query.Field.newBuilder().setName("col21").setTable("tbl").setType(Query.Type.VARCHAR).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList
            .add(Query.Field.newBuilder().setName("col22").setTable("tbl").setType(Query.Type.VARBINARY).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setFlags(Query.MySqlFlag.BINARY_FLAG_VALUE).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col23").setTable("tbl").setType(Query.Type.CHAR).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col24").setTable("tbl").setType(Query.Type.BINARY).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col25").setTable("tbl").setType(Query.Type.BIT).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col26").setTable("tbl").setType(Query.Type.ENUM).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col27").setTable("tbl").setType(Query.Type.SET).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList.add(Query.Field.newBuilder().setName("col28").setTable("tbl").setType(Query.Type.TUPLE).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
        fieldList
            .add(Query.Field.newBuilder().setName("col29").setTable("tbl").setType(Query.Type.VARBINARY).setMysqlType(MysqlDefs.FIELD_TYPE_TINY).setCharset(33).build());
    }

    public List<Query.Field> getFieldList() {

        initializeFieldList();
        return this.fieldList;
    }

    @Test public void testgetColumnCount() throws SQLException {

        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetadata = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        Assert.assertEquals(29, vitessResultSetMetadata.getColumnCount());
    }

    @Test public void testgetColumnName() throws SQLException {

        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetadata = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        for (int i = 1; i <= vitessResultSetMetadata.getColumnCount(); i++) {
            Assert.assertEquals("col" + i, vitessResultSetMetadata.getColumnName(i));
        }
    }

    @Test public void testgetColumnTypeName() throws SQLException {

        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetadata = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        Assert.assertEquals("TINYINT", vitessResultSetMetadata.getColumnTypeName(1));
        Assert.assertEquals("TINYINT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(2));
        Assert.assertEquals("SMALLINT", vitessResultSetMetadata.getColumnTypeName(3));
        Assert.assertEquals("SMALLINT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(4));
        Assert.assertEquals("MEDIUMINT", vitessResultSetMetadata.getColumnTypeName(5));
        Assert.assertEquals("MEDIUMINT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(6));
        Assert.assertEquals("INT", vitessResultSetMetadata.getColumnTypeName(7));
        Assert.assertEquals("INT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(8));
        Assert.assertEquals("BIGINT", vitessResultSetMetadata.getColumnTypeName(9));
        Assert.assertEquals("BIGINT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(10));
        Assert.assertEquals("FLOAT", vitessResultSetMetadata.getColumnTypeName(11));
        Assert.assertEquals("DOUBLE", vitessResultSetMetadata.getColumnTypeName(12));
        Assert.assertEquals("TIMESTAMP", vitessResultSetMetadata.getColumnTypeName(13));
        Assert.assertEquals("DATE", vitessResultSetMetadata.getColumnTypeName(14));
        Assert.assertEquals("TIME", vitessResultSetMetadata.getColumnTypeName(15));
        Assert.assertEquals("DATETIME", vitessResultSetMetadata.getColumnTypeName(16));
        Assert.assertEquals("YEAR", vitessResultSetMetadata.getColumnTypeName(17));
        Assert.assertEquals("DECIMAL", vitessResultSetMetadata.getColumnTypeName(18));
        Assert.assertEquals("TEXT", vitessResultSetMetadata.getColumnTypeName(19));
        Assert.assertEquals("BLOB", vitessResultSetMetadata.getColumnTypeName(20));
        Assert.assertEquals("VARCHAR", vitessResultSetMetadata.getColumnTypeName(21));
        Assert.assertEquals("VARBINARY", vitessResultSetMetadata.getColumnTypeName(22));
        Assert.assertEquals("CHAR", vitessResultSetMetadata.getColumnTypeName(23));
        Assert.assertEquals("BINARY", vitessResultSetMetadata.getColumnTypeName(24));
        Assert.assertEquals("BIT", vitessResultSetMetadata.getColumnTypeName(25));
        Assert.assertEquals("ENUM", vitessResultSetMetadata.getColumnTypeName(26));
        Assert.assertEquals("SET", vitessResultSetMetadata.getColumnTypeName(27));
        Assert.assertEquals("TUPLE", vitessResultSetMetadata.getColumnTypeName(28));
    }

    @Test public void testgetColumnType() throws SQLException {

        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetadata = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        Assert.assertEquals(-6, vitessResultSetMetadata.getColumnType(1));
        Assert.assertEquals(-6, vitessResultSetMetadata.getColumnType(2));
        Assert.assertEquals(5, vitessResultSetMetadata.getColumnType(3));
        Assert.assertEquals(5, vitessResultSetMetadata.getColumnType(4));
        Assert.assertEquals(4, vitessResultSetMetadata.getColumnType(5));
        Assert.assertEquals(4, vitessResultSetMetadata.getColumnType(6));
        Assert.assertEquals(4, vitessResultSetMetadata.getColumnType(7));
        Assert.assertEquals(4, vitessResultSetMetadata.getColumnType(8));
        Assert.assertEquals(-5, vitessResultSetMetadata.getColumnType(9));
        Assert.assertEquals(-5, vitessResultSetMetadata.getColumnType(10));
        Assert.assertEquals(6, vitessResultSetMetadata.getColumnType(11));
        Assert.assertEquals(8, vitessResultSetMetadata.getColumnType(12));
        Assert.assertEquals(93, vitessResultSetMetadata.getColumnType(13));
        Assert.assertEquals(91, vitessResultSetMetadata.getColumnType(14));
        Assert.assertEquals(92, vitessResultSetMetadata.getColumnType(15));
        Assert.assertEquals(93, vitessResultSetMetadata.getColumnType(16));
        Assert.assertEquals(5, vitessResultSetMetadata.getColumnType(17));
        Assert.assertEquals(3, vitessResultSetMetadata.getColumnType(18));
        Assert.assertEquals(12, vitessResultSetMetadata.getColumnType(19));
        Assert.assertEquals(2004, vitessResultSetMetadata.getColumnType(20));
        Assert.assertEquals(12, vitessResultSetMetadata.getColumnType(21));
        Assert.assertEquals(-3, vitessResultSetMetadata.getColumnType(22));
        Assert.assertEquals(1, vitessResultSetMetadata.getColumnType(23));
        Assert.assertEquals(-2, vitessResultSetMetadata.getColumnType(24));
        Assert.assertEquals(-7, vitessResultSetMetadata.getColumnType(25));
        Assert.assertEquals(1, vitessResultSetMetadata.getColumnType(26));
        Assert.assertEquals(1, vitessResultSetMetadata.getColumnType(27));
        Assert.assertEquals(12, vitessResultSetMetadata.getColumnType(29));
        try {
            int type = vitessResultSetMetadata.getColumnType(28);
        } catch (SQLException ex) {
            Assert
                .assertEquals(Constants.SQLExceptionMessages.INVALID_COLUMN_TYPE, ex.getMessage());
        }

        try {
            int type = vitessResultSetMetadata.getColumnType(0);
        } catch (SQLException ex) {
            Assert.assertEquals(Constants.SQLExceptionMessages.INVALID_COLUMN_INDEX + ": " + 0,
                ex.getMessage());
        }
    }

    @Test public void testgetColumnLabel() throws SQLException {
        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetaData = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        for (int i = 1; i <= vitessResultSetMetaData.getColumnCount(); i++) {
            Assert.assertEquals("col" + i, vitessResultSetMetaData.getColumnLabel(i));
        }
    }

    @Test public void testgetTableName() throws SQLException {
        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetadata = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        for (int i = 1; i <= vitessResultSetMetadata.getColumnCount(); i++) {
            Assert.assertEquals(vitessResultSetMetadata.getTableName(i), "tbl");
        }
    }

    @Test public void isReadOnlyTest() throws SQLException {
        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetadata = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        for (int i = 1; i <= vitessResultSetMetadata.getColumnCount(); i++) {
            Assert.assertEquals(vitessResultSetMetadata.isReadOnly(i), false);
            Assert.assertEquals(vitessResultSetMetadata.isWritable(i), true);
            Assert.assertEquals(vitessResultSetMetadata.isDefinitelyWritable(i), true);
        }
    }

    @Test public void getColumnClassNameTest() throws SQLException {
        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetadata = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        Assert.assertEquals("TINYINT", vitessResultSetMetadata.getColumnTypeName(1));
        Assert.assertEquals("TINYINT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(2));
        Assert.assertEquals("SMALLINT", vitessResultSetMetadata.getColumnTypeName(3));
        Assert.assertEquals("SMALLINT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(4));
        Assert.assertEquals("MEDIUMINT", vitessResultSetMetadata.getColumnTypeName(5));
        Assert.assertEquals("MEDIUMINT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(6));
        Assert.assertEquals("INT", vitessResultSetMetadata.getColumnTypeName(7));
        Assert.assertEquals("INT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(8));
        Assert.assertEquals("BIGINT", vitessResultSetMetadata.getColumnTypeName(9));
        Assert.assertEquals("BIGINT UNSIGNED", vitessResultSetMetadata.getColumnTypeName(10));
        Assert.assertEquals("FLOAT", vitessResultSetMetadata.getColumnTypeName(11));
        Assert.assertEquals("DOUBLE", vitessResultSetMetadata.getColumnTypeName(12));
        Assert.assertEquals("TIMESTAMP", vitessResultSetMetadata.getColumnTypeName(13));
        Assert.assertEquals("DATE", vitessResultSetMetadata.getColumnTypeName(14));
        Assert.assertEquals("TIME", vitessResultSetMetadata.getColumnTypeName(15));
        Assert.assertEquals("DATETIME", vitessResultSetMetadata.getColumnTypeName(16));
        Assert.assertEquals("YEAR", vitessResultSetMetadata.getColumnTypeName(17));
        Assert.assertEquals("DECIMAL", vitessResultSetMetadata.getColumnTypeName(18));
        Assert.assertEquals("TEXT", vitessResultSetMetadata.getColumnTypeName(19));
        Assert.assertEquals("BLOB", vitessResultSetMetadata.getColumnTypeName(20));
        Assert.assertEquals("VARCHAR", vitessResultSetMetadata.getColumnTypeName(21));
        Assert.assertEquals("VARBINARY", vitessResultSetMetadata.getColumnTypeName(22));
        Assert.assertEquals("CHAR", vitessResultSetMetadata.getColumnTypeName(23));
        Assert.assertEquals("BINARY", vitessResultSetMetadata.getColumnTypeName(24));
        Assert.assertEquals("BIT", vitessResultSetMetadata.getColumnTypeName(25));
        Assert.assertEquals("ENUM", vitessResultSetMetadata.getColumnTypeName(26));
        Assert.assertEquals("SET", vitessResultSetMetadata.getColumnTypeName(27));
        Assert.assertEquals("TUPLE", vitessResultSetMetadata.getColumnTypeName(28));
    }

    @Test public void getSchemaNameTest() throws SQLException {
        List<Query.Field> fieldList = getFieldList();
        VitessResultSetMetaData vitessResultSetMetaData = new VitessResultSetMetaData(getVitessConnection(), fieldList);
        Assert.assertEquals(vitessResultSetMetaData.getSchemaName(1), "");
        Assert.assertEquals(vitessResultSetMetaData.getCatalogName(1), "");
        Assert.assertEquals(vitessResultSetMetaData.getPrecision(1), 2);
        Assert.assertEquals(vitessResultSetMetaData.getScale(1), 0);
        Assert.assertEquals(vitessResultSetMetaData.getColumnDisplaySize(1), 1);
        Assert.assertEquals(vitessResultSetMetaData.isCurrency(1), false);
    }
}

