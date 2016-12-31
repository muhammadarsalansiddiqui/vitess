package com.flipkart.vitess.util;

import org.joda.time.Duration;

import com.youtube.vitess.client.Context;
import com.youtube.vitess.proto.Vtrpc;

/**
 * Created by naveen.nahata on 24/02/16.
 */
public class CommonUtils {

    /**
     * Create context used to create grpc client and executing query.
     *
     * @param username
     * @param connectionTimeout
     * @return
     */
    public static Context createContext(String username, boolean includeFieldMetadata, long connectionTimeout) {
        Context context;
        Vtrpc.CallerID callerID = null;
        if (null != username) {
            callerID = Vtrpc.CallerID.newBuilder().setPrincipal(username).build();
        }
        if (null != callerID) {
            context = Context.getDefault()
                .withDeadlineAfter(Duration.millis(connectionTimeout))
                .withCallerId(callerID)
                .withIncludeFieldMetadata(includeFieldMetadata);
        } else {
            context = Context.getDefault()
                .withDeadlineAfter(Duration.millis(connectionTimeout))
                .withIncludeFieldMetadata(includeFieldMetadata);
        }
        return context;
    }

}

