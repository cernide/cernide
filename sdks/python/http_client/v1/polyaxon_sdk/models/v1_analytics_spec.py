#!/usr/bin/python
#
# Copyright 2018-2021 Polyaxon, Inc.
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

    The version of the OpenAPI document: 1.12.0
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1AnalyticsSpec(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'view': 'str',
        'trunc': 'str',
        'groupby': 'str',
        'frequency': 'str'
    }

    attribute_map = {
        'view': 'view',
        'trunc': 'trunc',
        'groupby': 'groupby',
        'frequency': 'frequency'
    }

    def __init__(self, view=None, trunc=None, groupby=None, frequency=None, local_vars_configuration=None):  # noqa: E501
        """V1AnalyticsSpec - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._view = None
        self._trunc = None
        self._groupby = None
        self._frequency = None
        self.discriminator = None

        if view is not None:
            self.view = view
        if trunc is not None:
            self.trunc = trunc
        if groupby is not None:
            self.groupby = groupby
        if frequency is not None:
            self.frequency = frequency

    @property
    def view(self):
        """Gets the view of this V1AnalyticsSpec.  # noqa: E501


        :return: The view of this V1AnalyticsSpec.  # noqa: E501
        :rtype: str
        """
        return self._view

    @view.setter
    def view(self, view):
        """Sets the view of this V1AnalyticsSpec.


        :param view: The view of this V1AnalyticsSpec.  # noqa: E501
        :type: str
        """

        self._view = view

    @property
    def trunc(self):
        """Gets the trunc of this V1AnalyticsSpec.  # noqa: E501


        :return: The trunc of this V1AnalyticsSpec.  # noqa: E501
        :rtype: str
        """
        return self._trunc

    @trunc.setter
    def trunc(self, trunc):
        """Sets the trunc of this V1AnalyticsSpec.


        :param trunc: The trunc of this V1AnalyticsSpec.  # noqa: E501
        :type: str
        """

        self._trunc = trunc

    @property
    def groupby(self):
        """Gets the groupby of this V1AnalyticsSpec.  # noqa: E501


        :return: The groupby of this V1AnalyticsSpec.  # noqa: E501
        :rtype: str
        """
        return self._groupby

    @groupby.setter
    def groupby(self, groupby):
        """Sets the groupby of this V1AnalyticsSpec.


        :param groupby: The groupby of this V1AnalyticsSpec.  # noqa: E501
        :type: str
        """

        self._groupby = groupby

    @property
    def frequency(self):
        """Gets the frequency of this V1AnalyticsSpec.  # noqa: E501


        :return: The frequency of this V1AnalyticsSpec.  # noqa: E501
        :rtype: str
        """
        return self._frequency

    @frequency.setter
    def frequency(self, frequency):
        """Sets the frequency of this V1AnalyticsSpec.


        :param frequency: The frequency of this V1AnalyticsSpec.  # noqa: E501
        :type: str
        """

        self._frequency = frequency

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1AnalyticsSpec):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1AnalyticsSpec):
            return True

        return self.to_dict() != other.to_dict()
