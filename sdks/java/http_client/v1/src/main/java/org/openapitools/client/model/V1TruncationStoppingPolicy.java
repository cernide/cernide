// Copyright 2018-2021 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.12.0
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

/**
 * Early stopping with truncation stopping, this policy stops a percentage of all running runs at every evaluation.
 */
@ApiModel(description = "Early stopping with truncation stopping, this policy stops a percentage of all running runs at every evaluation.")

public class V1TruncationStoppingPolicy {
  public static final String SERIALIZED_NAME_KIND = "kind";
  @SerializedName(SERIALIZED_NAME_KIND)
  private String kind = "truncation";

  public static final String SERIALIZED_NAME_PERCENT = "percent";
  @SerializedName(SERIALIZED_NAME_PERCENT)
  private Integer percent;

  public static final String SERIALIZED_NAME_EVALUATION_INTERVAL = "evaluationInterval";
  @SerializedName(SERIALIZED_NAME_EVALUATION_INTERVAL)
  private Integer evaluationInterval;

  public static final String SERIALIZED_NAME_MIN_INTERVAL = "minInterval";
  @SerializedName(SERIALIZED_NAME_MIN_INTERVAL)
  private Integer minInterval;

  public static final String SERIALIZED_NAME_MIN_SAMPLES = "minSamples";
  @SerializedName(SERIALIZED_NAME_MIN_SAMPLES)
  private Integer minSamples;

  public static final String SERIALIZED_NAME_INCLUDE_SUCCEEDED = "includeSucceeded";
  @SerializedName(SERIALIZED_NAME_INCLUDE_SUCCEEDED)
  private Boolean includeSucceeded;


  public V1TruncationStoppingPolicy kind(String kind) {
    
    this.kind = kind;
    return this;
  }

   /**
   * Get kind
   * @return kind
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getKind() {
    return kind;
  }


  public void setKind(String kind) {
    this.kind = kind;
  }


  public V1TruncationStoppingPolicy percent(Integer percent) {
    
    this.percent = percent;
    return this;
  }

   /**
   * The percentage of runs to stop, at each evaluation interval. e.g. 1 - 99.
   * @return percent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The percentage of runs to stop, at each evaluation interval. e.g. 1 - 99.")

  public Integer getPercent() {
    return percent;
  }


  public void setPercent(Integer percent) {
    this.percent = percent;
  }


  public V1TruncationStoppingPolicy evaluationInterval(Integer evaluationInterval) {
    
    this.evaluationInterval = evaluationInterval;
    return this;
  }

   /**
   * Interval/Frequency for applying the policy.
   * @return evaluationInterval
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Interval/Frequency for applying the policy.")

  public Integer getEvaluationInterval() {
    return evaluationInterval;
  }


  public void setEvaluationInterval(Integer evaluationInterval) {
    this.evaluationInterval = evaluationInterval;
  }


  public V1TruncationStoppingPolicy minInterval(Integer minInterval) {
    
    this.minInterval = minInterval;
    return this;
  }

   /**
   * Get minInterval
   * @return minInterval
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getMinInterval() {
    return minInterval;
  }


  public void setMinInterval(Integer minInterval) {
    this.minInterval = minInterval;
  }


  public V1TruncationStoppingPolicy minSamples(Integer minSamples) {
    
    this.minSamples = minSamples;
    return this;
  }

   /**
   * Get minSamples
   * @return minSamples
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getMinSamples() {
    return minSamples;
  }


  public void setMinSamples(Integer minSamples) {
    this.minSamples = minSamples;
  }


  public V1TruncationStoppingPolicy includeSucceeded(Boolean includeSucceeded) {
    
    this.includeSucceeded = includeSucceeded;
    return this;
  }

   /**
   * Get includeSucceeded
   * @return includeSucceeded
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIncludeSucceeded() {
    return includeSucceeded;
  }


  public void setIncludeSucceeded(Boolean includeSucceeded) {
    this.includeSucceeded = includeSucceeded;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1TruncationStoppingPolicy v1TruncationStoppingPolicy = (V1TruncationStoppingPolicy) o;
    return Objects.equals(this.kind, v1TruncationStoppingPolicy.kind) &&
        Objects.equals(this.percent, v1TruncationStoppingPolicy.percent) &&
        Objects.equals(this.evaluationInterval, v1TruncationStoppingPolicy.evaluationInterval) &&
        Objects.equals(this.minInterval, v1TruncationStoppingPolicy.minInterval) &&
        Objects.equals(this.minSamples, v1TruncationStoppingPolicy.minSamples) &&
        Objects.equals(this.includeSucceeded, v1TruncationStoppingPolicy.includeSucceeded);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, percent, evaluationInterval, minInterval, minSamples, includeSucceeded);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1TruncationStoppingPolicy {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    percent: ").append(toIndentedString(percent)).append("\n");
    sb.append("    evaluationInterval: ").append(toIndentedString(evaluationInterval)).append("\n");
    sb.append("    minInterval: ").append(toIndentedString(minInterval)).append("\n");
    sb.append("    minSamples: ").append(toIndentedString(minSamples)).append("\n");
    sb.append("    includeSucceeded: ").append(toIndentedString(includeSucceeded)).append("\n");
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

