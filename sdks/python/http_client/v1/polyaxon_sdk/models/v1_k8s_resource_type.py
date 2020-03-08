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


class V1K8sResourceType(object):
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
        "name": "str",
        "schema": "V1K8sResourceSchema",
        "is_requested": "bool",
    }

    attribute_map = {"name": "name", "schema": "schema", "is_requested": "is_requested"}

    def __init__(self, name=None, schema=None, is_requested=None):  # noqa: E501
        """V1K8sResourceType - a model defined in Swagger"""  # noqa: E501

        self._name = None
        self._schema = None
        self._is_requested = None
        self.discriminator = None

        if name is not None:
            self.name = name
        if schema is not None:
            self.schema = schema
        if is_requested is not None:
            self.is_requested = is_requested

    @property
    def name(self):
        """Gets the name of this V1K8sResourceType.  # noqa: E501


        :return: The name of this V1K8sResourceType.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this V1K8sResourceType.


        :param name: The name of this V1K8sResourceType.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def schema(self):
        """Gets the schema of this V1K8sResourceType.  # noqa: E501


        :return: The schema of this V1K8sResourceType.  # noqa: E501
        :rtype: V1K8sResourceSchema
        """
        return self._schema

    @schema.setter
    def schema(self, schema):
        """Sets the schema of this V1K8sResourceType.


        :param schema: The schema of this V1K8sResourceType.  # noqa: E501
        :type: V1K8sResourceSchema
        """

        self._schema = schema

    @property
    def is_requested(self):
        """Gets the is_requested of this V1K8sResourceType.  # noqa: E501


        :return: The is_requested of this V1K8sResourceType.  # noqa: E501
        :rtype: bool
        """
        return self._is_requested

    @is_requested.setter
    def is_requested(self, is_requested):
        """Sets the is_requested of this V1K8sResourceType.


        :param is_requested: The is_requested of this V1K8sResourceType.  # noqa: E501
        :type: bool
        """

        self._is_requested = is_requested

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
        if issubclass(V1K8sResourceType, dict):
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
        if not isinstance(other, V1K8sResourceType):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
