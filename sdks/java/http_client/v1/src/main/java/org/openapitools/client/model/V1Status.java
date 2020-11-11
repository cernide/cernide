// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.3.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.V1StatusCondition;
import org.openapitools.client.model.V1Statuses;

/**
 * V1Status
 */

public class V1Status {
  public static final String SERIALIZED_NAME_UUID = "uuid";
  @SerializedName(SERIALIZED_NAME_UUID)
  private String uuid;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private V1Statuses status = V1Statuses.CREATED;

  public static final String SERIALIZED_NAME_STATUS_CONDITIONS = "status_conditions";
  @SerializedName(SERIALIZED_NAME_STATUS_CONDITIONS)
  private List<V1StatusCondition> statusConditions = null;


  public V1Status uuid(String uuid) {
    
    this.uuid = uuid;
    return this;
  }

   /**
   * Get uuid
   * @return uuid
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUuid() {
    return uuid;
  }


  public void setUuid(String uuid) {
    this.uuid = uuid;
  }


  public V1Status status(V1Statuses status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1Statuses getStatus() {
    return status;
  }


  public void setStatus(V1Statuses status) {
    this.status = status;
  }


  public V1Status statusConditions(List<V1StatusCondition> statusConditions) {
    
    this.statusConditions = statusConditions;
    return this;
  }

  public V1Status addStatusConditionsItem(V1StatusCondition statusConditionsItem) {
    if (this.statusConditions == null) {
      this.statusConditions = new ArrayList<V1StatusCondition>();
    }
    this.statusConditions.add(statusConditionsItem);
    return this;
  }

   /**
   * Get statusConditions
   * @return statusConditions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1StatusCondition> getStatusConditions() {
    return statusConditions;
  }


  public void setStatusConditions(List<V1StatusCondition> statusConditions) {
    this.statusConditions = statusConditions;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Status v1Status = (V1Status) o;
    return Objects.equals(this.uuid, v1Status.uuid) &&
        Objects.equals(this.status, v1Status.status) &&
        Objects.equals(this.statusConditions, v1Status.statusConditions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(uuid, status, statusConditions);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Status {\n");
    sb.append("    uuid: ").append(toIndentedString(uuid)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    statusConditions: ").append(toIndentedString(statusConditions)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

