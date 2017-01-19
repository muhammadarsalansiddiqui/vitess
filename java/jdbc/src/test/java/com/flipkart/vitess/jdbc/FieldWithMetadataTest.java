package com.flipkart.vitess.jdbc;

import com.flipkart.vitess.util.MysqlDefs;
import com.youtube.vitess.proto.Query;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.internal.verification.VerificationModeFactory;
import org.powermock.api.mockito.PowerMockito;
import org.powermock.core.classloader.annotations.PrepareForTest;
import org.powermock.modules.junit4.PowerMockRunner;

import java.sql.SQLException;
import java.sql.Types;

@PrepareForTest(FieldWithMetadata.class)
@RunWith(PowerMockRunner.class)
public class FieldWithMetadataTest extends BaseTest {

    @Test
    public void testTinyIntAsBitOrBool() throws SQLException {
        VitessConnection conn = getVitessConnection();

        Query.Field raw = Query.Field.newBuilder()
            .setTable("foo")
            .setColumnLength(3)
            .setType(Query.Type.INT8)
            .setName("foo")
            .setOrgName("foo")
            .build();

        conn.setTinyInt1isBit(true);

        FieldWithMetadata fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(Types.TINYINT, fieldWithMetadata.getJavaType());

        raw = raw.toBuilder()
            .setColumnLength(1)
            .build();
        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(Types.BIT, fieldWithMetadata.getJavaType());

        conn.setTransformedBitIsBoolean(true);
        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(Types.BOOLEAN, fieldWithMetadata.getJavaType());

        conn.setIncludedFields(Query.ExecuteOptions.IncludedFields.TYPE_AND_NAME);
        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(Types.TINYINT, fieldWithMetadata.getJavaType());
    }

    @Test
    public void testPrecisionAdjustFactor() throws SQLException {
        VitessConnection conn = getVitessConnection();

        assertPrecisionEquals(conn, Query.Type.FLOAT32, true, 0, 0);
        assertPrecisionEquals(conn, Query.Type.FLOAT64, true, 32, 0);
        assertPrecisionEquals(conn, Query.Type.BIT, true, 0, 0);
        assertPrecisionEquals(conn, Query.Type.DECIMAL, true, 0, -1);
        assertPrecisionEquals(conn, Query.Type.DECIMAL, true, 3, -2);
        assertPrecisionEquals(conn, Query.Type.INT32, true, /* this can't happen, but just checking */3, -2);
        assertPrecisionEquals(conn, Query.Type.INT32, true, 0, -1);

        assertPrecisionEquals(conn, Query.Type.FLOAT32, false, 0, 0);
        assertPrecisionEquals(conn, Query.Type.FLOAT64, false, 32, 0);
        assertPrecisionEquals(conn, Query.Type.BIT, false, 0, 0);
        assertPrecisionEquals(conn, Query.Type.DECIMAL, false, 0, -1);
        assertPrecisionEquals(conn, Query.Type.DECIMAL, false, 3, -1);
        assertPrecisionEquals(conn, Query.Type.UINT32, false, 0, 0);

        conn.setIncludedFields(Query.ExecuteOptions.IncludedFields.TYPE_AND_NAME);
        for (Query.Type type : Query.Type.values()) {
            if (type == Query.Type.UNRECOGNIZED) {
                continue;
            }

            // All should be 0
            assertPrecisionEquals(conn, type, true, 0, 0);
            assertPrecisionEquals(conn, type, false, 0, 0);
            assertPrecisionEquals(conn, type, true, 2, 0);
            assertPrecisionEquals(conn, type, false, 2, 0);
        }
    }

    private void assertPrecisionEquals(VitessConnection conn, Query.Type fieldType, boolean signed, int decimals, int expectedPrecisionAdjustFactor) throws SQLException {
        Query.Field raw = Query.Field.newBuilder()
            .setTable("foo")
            .setColumnLength(3)
            .setType(fieldType)
            .setDecimals(decimals)
            .setFlags(signed ? 0 : Query.MySqlFlag.UNSIGNED_FLAG_VALUE)
            .setName("foo")
            .setOrgName("foo")
            .build();

        FieldWithMetadata fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(expectedPrecisionAdjustFactor, fieldWithMetadata.getPrecisionAdjustFactor());
    }

    @Test
    public void testFlags() throws SQLException {
        VitessConnection conn = getVitessConnection();

        Query.Field raw = Query.Field.newBuilder()
            .setTable("foo")
            .setColumnLength(3)
            .setType(Query.Type.VARBINARY)
            .setName("foo")
            .setOrgName("foo")
            .build();

        FieldWithMetadata fieldWithMetadata = new FieldWithMetadata(conn, raw);

        Assert.assertEquals(false, fieldWithMetadata.isBinary());
        Assert.assertEquals(false, fieldWithMetadata.isBlob());
        Assert.assertEquals(false, fieldWithMetadata.isAutoIncrement());
        Assert.assertEquals(false, fieldWithMetadata.isMultipleKey());
        Assert.assertEquals(false, fieldWithMetadata.isNotNull());
        Assert.assertEquals(false, fieldWithMetadata.isPrimaryKey());
        Assert.assertEquals(false, fieldWithMetadata.isUniqueKey());
        Assert.assertEquals(false, fieldWithMetadata.isUnsigned());
        Assert.assertEquals(/* just inverses isUnsigned */true, fieldWithMetadata.isSigned());
        Assert.assertEquals(false, fieldWithMetadata.isZeroFill());

        int value = 0;
        for (Query.MySqlFlag flag : Query.MySqlFlag.values()) {
            if (flag == Query.MySqlFlag.UNRECOGNIZED) {
                continue;
            }
            value |= flag.getNumber();
        }

        raw = raw.toBuilder()
            .setFlags(value)
            .build();

        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(true, fieldWithMetadata.isBinary());
        Assert.assertEquals(true, fieldWithMetadata.isBlob());
        Assert.assertEquals(true, fieldWithMetadata.isAutoIncrement());
        Assert.assertEquals(true, fieldWithMetadata.isMultipleKey());
        Assert.assertEquals(true, fieldWithMetadata.isNotNull());
        Assert.assertEquals(true, fieldWithMetadata.isPrimaryKey());
        Assert.assertEquals(true, fieldWithMetadata.isUniqueKey());
        Assert.assertEquals(true, fieldWithMetadata.isUnsigned());
        Assert.assertEquals(/* just inverses isUnsigned */false, fieldWithMetadata.isSigned());
        Assert.assertEquals(true, fieldWithMetadata.isZeroFill());

        conn.setIncludedFields(Query.ExecuteOptions.IncludedFields.TYPE_AND_NAME);
        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(false, fieldWithMetadata.isBinary());
        Assert.assertEquals(false, fieldWithMetadata.isBlob());
        Assert.assertEquals(false, fieldWithMetadata.isAutoIncrement());
        Assert.assertEquals(false, fieldWithMetadata.isMultipleKey());
        Assert.assertEquals(true, fieldWithMetadata.isNotNull());
        Assert.assertEquals(false, fieldWithMetadata.isPrimaryKey());
        Assert.assertEquals(false, fieldWithMetadata.isUniqueKey());
        Assert.assertEquals(true, fieldWithMetadata.isUnsigned());
        Assert.assertEquals(/* just inverses isUnsigned */false, fieldWithMetadata.isSigned());
        Assert.assertEquals(false, fieldWithMetadata.isZeroFill());

    }

    @Test
    public void testReadOnly() throws SQLException {
        VitessConnection conn = getVitessConnection();

        Query.Field raw = Query.Field.newBuilder()
            .setTable("foo")
            .setType(Query.Type.CHAR)
            .setName("foo")
            .setOrgName("foo")
            .setOrgTable("foo")
            .build();

        FieldWithMetadata fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(false, fieldWithMetadata.isReadOnly());

        raw = raw.toBuilder()
            .setOrgName("")
            .setOrgTable("foo")
            .build();

        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(true, fieldWithMetadata.isReadOnly());

        raw = raw.toBuilder()
            .setOrgName("foo")
            .setOrgTable("")
            .build();

        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(true, fieldWithMetadata.isReadOnly());

        raw = raw.toBuilder()
            .setOrgTable("")
            .setOrgName("")
            .build();

        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(true, fieldWithMetadata.isReadOnly());

        conn.setIncludedFields(Query.ExecuteOptions.IncludedFields.TYPE_AND_NAME);
        fieldWithMetadata = new FieldWithMetadata(conn, raw);
        Assert.assertEquals(false, fieldWithMetadata.isReadOnly());
    }
}
