#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    OpenAPI spec version: 1.0.5
    Contact: contact@polyaxon.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class V1TruncationStoppingPolicy(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        "kind": "str",
        "percent": "int",
        "evaluation_interval": "int",
        "min_interval": "int",
        "min_samples": "int",
        "include_succeeded": "bool",
    }

    attribute_map = {
        "kind": "kind",
        "percent": "percent",
        "evaluation_interval": "evaluation_interval",
        "min_interval": "min_interval",
        "min_samples": "min_samples",
        "include_succeeded": "include_succeeded",
    }

    def __init__(
        self,
        kind="truncation",
        percent=None,
        evaluation_interval=None,
        min_interval=None,
        min_samples=None,
        include_succeeded=None,
    ):  # noqa: E501
        """V1TruncationStoppingPolicy - a model defined in Swagger"""  # noqa: E501

        self._kind = None
        self._percent = None
        self._evaluation_interval = None
        self._min_interval = None
        self._min_samples = None
        self._include_succeeded = None
        self.discriminator = None

        if kind is not None:
            self.kind = kind
        if percent is not None:
            self.percent = percent
        if evaluation_interval is not None:
            self.evaluation_interval = evaluation_interval
        if min_interval is not None:
            self.min_interval = min_interval
        if min_samples is not None:
            self.min_samples = min_samples
        if include_succeeded is not None:
            self.include_succeeded = include_succeeded

    @property
    def kind(self):
        """Gets the kind of this V1TruncationStoppingPolicy.  # noqa: E501


        :return: The kind of this V1TruncationStoppingPolicy.  # noqa: E501
        :rtype: str
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this V1TruncationStoppingPolicy.


        :param kind: The kind of this V1TruncationStoppingPolicy.  # noqa: E501
        :type: str
        """

        self._kind = kind

    @property
    def percent(self):
        """Gets the percent of this V1TruncationStoppingPolicy.  # noqa: E501

        The percentage of runs to stop, at each evaluation interval. e.g. 1 - 99.  # noqa: E501

        :return: The percent of this V1TruncationStoppingPolicy.  # noqa: E501
        :rtype: int
        """
        return self._percent

    @percent.setter
    def percent(self, percent):
        """Sets the percent of this V1TruncationStoppingPolicy.

        The percentage of runs to stop, at each evaluation interval. e.g. 1 - 99.  # noqa: E501

        :param percent: The percent of this V1TruncationStoppingPolicy.  # noqa: E501
        :type: int
        """

        self._percent = percent

    @property
    def evaluation_interval(self):
        """Gets the evaluation_interval of this V1TruncationStoppingPolicy.  # noqa: E501

        Interval/Frequency for applying the policy.  # noqa: E501

        :return: The evaluation_interval of this V1TruncationStoppingPolicy.  # noqa: E501
        :rtype: int
        """
        return self._evaluation_interval

    @evaluation_interval.setter
    def evaluation_interval(self, evaluation_interval):
        """Sets the evaluation_interval of this V1TruncationStoppingPolicy.

        Interval/Frequency for applying the policy.  # noqa: E501

        :param evaluation_interval: The evaluation_interval of this V1TruncationStoppingPolicy.  # noqa: E501
        :type: int
        """

        self._evaluation_interval = evaluation_interval

    @property
    def min_interval(self):
        """Gets the min_interval of this V1TruncationStoppingPolicy.  # noqa: E501


        :return: The min_interval of this V1TruncationStoppingPolicy.  # noqa: E501
        :rtype: int
        """
        return self._min_interval

    @min_interval.setter
    def min_interval(self, min_interval):
        """Sets the min_interval of this V1TruncationStoppingPolicy.


        :param min_interval: The min_interval of this V1TruncationStoppingPolicy.  # noqa: E501
        :type: int
        """

        self._min_interval = min_interval

    @property
    def min_samples(self):
        """Gets the min_samples of this V1TruncationStoppingPolicy.  # noqa: E501


        :return: The min_samples of this V1TruncationStoppingPolicy.  # noqa: E501
        :rtype: int
        """
        return self._min_samples

    @min_samples.setter
    def min_samples(self, min_samples):
        """Sets the min_samples of this V1TruncationStoppingPolicy.


        :param min_samples: The min_samples of this V1TruncationStoppingPolicy.  # noqa: E501
        :type: int
        """

        self._min_samples = min_samples

    @property
    def include_succeeded(self):
        """Gets the include_succeeded of this V1TruncationStoppingPolicy.  # noqa: E501


        :return: The include_succeeded of this V1TruncationStoppingPolicy.  # noqa: E501
        :rtype: bool
        """
        return self._include_succeeded

    @include_succeeded.setter
    def include_succeeded(self, include_succeeded):
        """Sets the include_succeeded of this V1TruncationStoppingPolicy.


        :param include_succeeded: The include_succeeded of this V1TruncationStoppingPolicy.  # noqa: E501
        :type: bool
        """

        self._include_succeeded = include_succeeded

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(
                    map(lambda x: x.to_dict() if hasattr(x, "to_dict") else x, value)
                )
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(
                    map(
                        lambda item: (item[0], item[1].to_dict())
                        if hasattr(item[1], "to_dict")
                        else item,
                        value.items(),
                    )
                )
            else:
                result[attr] = value
        if issubclass(V1TruncationStoppingPolicy, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1TruncationStoppingPolicy):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
