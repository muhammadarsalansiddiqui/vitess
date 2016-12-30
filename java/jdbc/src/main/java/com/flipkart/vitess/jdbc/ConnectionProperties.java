package com.flipkart.vitess.jdbc;

import java.io.UnsupportedEncodingException;
import java.lang.reflect.Field;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.Properties;

import com.flipkart.vitess.util.StringUtils;

public class ConnectionProperties {

  private static final ArrayList<java.lang.reflect.Field> PROPERTY_LIST = new ArrayList<>();
  static {
    try {
      java.lang.reflect.Field[] declaredFields = ConnectionProperties.class.getDeclaredFields();

      for (int i = 0; i < declaredFields.length; i++) {
        if (ConnectionProperties.ConnectionProperty.class.isAssignableFrom(declaredFields[i].getType())) {
          PROPERTY_LIST.add(declaredFields[i]);
        }
      }
    } catch (Exception ex) {
      RuntimeException rtEx = new RuntimeException();
      rtEx.initCause(ex);

      throw rtEx;
    }
  }

  private BooleanConnectionProperty blobsAreStrings = new BooleanConnectionProperty("blobsAreStrings", false);
  private BooleanConnectionProperty useBlobToStoreUTF8OutsideBMP = new BooleanConnectionProperty("useBlobToStoreUTF8OutsideBMP", false);
  private BooleanConnectionProperty tinyInt1isBit = new BooleanConnectionProperty("tinyInt1isBit", true);
  private BooleanConnectionProperty transformedBitIsBoolean = new BooleanConnectionProperty("transformedBitIsBoolean", false);
  private BooleanConnectionProperty functionsNeverReturnBlobs = new BooleanConnectionProperty("functionsNeverReturnBlobs", false);
  private BooleanConnectionProperty paranoid = new BooleanConnectionProperty("paranoid", false);
  private BooleanConnectionProperty cacheServerConfiguration = new BooleanConnectionProperty("cacheServerConfiguration", false);
  private BooleanConnectionProperty yearIsDateType = new BooleanConnectionProperty("yearIsDateType", true);

  private StringConnectionProperty utf8OutsideBmpIncludedColumnNamePattern = new StringConnectionProperty("utf8OutsideBmpIncludedColumnNamePattern", null);
  private StringConnectionProperty utf8OutsideBmpExcludedColumnNamePattern = new StringConnectionProperty("utf8OutsideBmpExcludedColumnNamePattern", null);
  private StringConnectionProperty characterEncoding = new StringConnectionProperty("characterEncoding", null);

  protected void initializeProperties(Properties props) throws SQLException {
    Properties propsCopy = (Properties) props.clone();

    for (Field propertyField : PROPERTY_LIST) {
      try {
        ConnectionProperty propToSet = (ConnectionProperty) propertyField.get(this);

        propToSet.initializeFrom(propsCopy);
      } catch (IllegalAccessException iae) {
        throw new SQLException("ConnectionProperties.unableToInitDriverProperties" + iae.toString());
      }
    }

    postInitialization();
  }

  private void postInitialization() throws SQLException {
    String testEncoding = ((String) this.characterEncoding.getValueAsObject());

    if (testEncoding != null) {
      // Attempt to use the encoding, and bail out if it can't be used
      try {
        String testString = "abc";
        StringUtils.getBytes(testString, testEncoding);
      } catch (UnsupportedEncodingException UE) {
        throw new SQLException("ConnectionProperties.unsupportedCharacterEncoding: "  + testEncoding);
      }
    }
  }

  public boolean getBlobsAreStrings() {
    return blobsAreStrings.getValueAsBoolean();
  }

  public boolean getUseBlobToStoreUTF8OutsideBMP() {
    return useBlobToStoreUTF8OutsideBMP.getValueAsBoolean();
  }

  public boolean getTinyInt1isBit() {
    return tinyInt1isBit.getValueAsBoolean();
  }

  public boolean getTransformedBitIsBoolean() {
    return transformedBitIsBoolean.getValueAsBoolean();
  }

  public boolean getFunctionsNeverReturnBlobs() {
    return functionsNeverReturnBlobs.getValueAsBoolean();
  }

  public String getUtf8OutsideBmpIncludedColumnNamePattern() {
    return utf8OutsideBmpIncludedColumnNamePattern.getValueAsString();
  }

  public String getUtf8OutsideBmpExcludedColumnNamePattern() {
    return utf8OutsideBmpExcludedColumnNamePattern.getValueAsString();
  }

  public boolean getParanoid() {
    return paranoid.getValueAsBoolean();
  }

  public String getCharacterEncoding() {
    return characterEncoding.getValueAsString();
  }

  public boolean getCacheServerConfiguration() {
    return cacheServerConfiguration.getValueAsBoolean();
  }

  public boolean getYearIsDateType() {
    return yearIsDateType.getValueAsBoolean();
  }
  private abstract static class ConnectionProperty {

    private final String name;
    protected final Object defaultValue;
    protected Object valueAsObject;

    private ConnectionProperty(String name, Object defaultValue) {
      this.name = name;
      this.defaultValue = defaultValue;
    }

    void initializeFrom(Properties extractFrom) {
      String extractedValue = extractFrom.getProperty(getPropertyName());
      extractFrom.remove(getPropertyName());
      initializeFrom(extractedValue);
    }

    abstract void initializeFrom(String extractedValue);

    public String getPropertyName() {
      return name;
    }

    public Object getDefaultValue() {
      return defaultValue;
    }

    public Object getValueAsObject() {
      return valueAsObject;
    }
  }

  private static class BooleanConnectionProperty extends ConnectionProperty {

    private BooleanConnectionProperty(String name, Object defaultValue) {
      super(name, defaultValue);
    }

    @Override
    void initializeFrom(String extractedValue) {
      if (extractedValue != null) {
        this.valueAsObject = Boolean.valueOf(extractedValue.equalsIgnoreCase("TRUE") || extractedValue.equalsIgnoreCase("YES"));
      } else {
        this.valueAsObject = this.defaultValue;
      }
    }

    boolean getValueAsBoolean() {
      return (boolean) valueAsObject;
    }
  }

  private static class StringConnectionProperty extends ConnectionProperty {

    private StringConnectionProperty(String name, Object defaultValue) {
      super(name, defaultValue);
    }

    @Override
    void initializeFrom(String extractedValue) {
      if (extractedValue != null) {
        this.valueAsObject = extractedValue;
      } else {
        this.valueAsObject = this.defaultValue;
      }
    }

    String getValueAsString() {
      return (String) valueAsObject;
    }
  }


}
