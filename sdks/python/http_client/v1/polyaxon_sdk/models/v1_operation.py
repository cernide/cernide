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

    The version of the OpenAPI document: 1.8.0
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1Operation(object):
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
        'version': 'float',
        'kind': 'str',
        'name': 'str',
        'description': 'str',
        'tags': 'list[str]',
        'presets': 'list[str]',
        'queue': 'str',
        'cache': 'V1Cache',
        'termination': 'V1Termination',
        'plugins': 'V1Plugins',
        'schedule': 'object',
        'events': 'list[V1EventTrigger]',
        'hooks': 'list[V1Hook]',
        'dependencies': 'list[str]',
        'trigger': 'V1TriggerPolicy',
        'conditions': 'str',
        'skip_on_upstream_skip': 'bool',
        'matrix': 'object',
        'joins': 'dict(str, V1Join)',
        'params': 'dict(str, V1Param)',
        'run_patch': 'object',
        'patch_strategy': 'V1PatchStrategy',
        'is_preset': 'bool',
        'is_approved': 'bool',
        'template': 'V1Template',
        'url_ref': 'str',
        'path_ref': 'str',
        'hub_ref': 'str',
        'dag_ref': 'str',
        'component': 'V1Component'
    }

    attribute_map = {
        'version': 'version',
        'kind': 'kind',
        'name': 'name',
        'description': 'description',
        'tags': 'tags',
        'presets': 'presets',
        'queue': 'queue',
        'cache': 'cache',
        'termination': 'termination',
        'plugins': 'plugins',
        'schedule': 'schedule',
        'events': 'events',
        'hooks': 'hooks',
        'dependencies': 'dependencies',
        'trigger': 'trigger',
        'conditions': 'conditions',
        'skip_on_upstream_skip': 'skipOnUpstreamSkip',
        'matrix': 'matrix',
        'joins': 'joins',
        'params': 'params',
        'run_patch': 'runPatch',
        'patch_strategy': 'patchStrategy',
        'is_preset': 'isPreset',
        'is_approved': 'isApproved',
        'template': 'template',
        'url_ref': 'urlRef',
        'path_ref': 'pathRef',
        'hub_ref': 'hubRef',
        'dag_ref': 'dagRef',
        'component': 'component'
    }

    def __init__(self, version=None, kind=None, name=None, description=None, tags=None, presets=None, queue=None, cache=None, termination=None, plugins=None, schedule=None, events=None, hooks=None, dependencies=None, trigger=None, conditions=None, skip_on_upstream_skip=None, matrix=None, joins=None, params=None, run_patch=None, patch_strategy=None, is_preset=None, is_approved=None, template=None, url_ref=None, path_ref=None, hub_ref=None, dag_ref=None, component=None, local_vars_configuration=None):  # noqa: E501
        """V1Operation - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._version = None
        self._kind = None
        self._name = None
        self._description = None
        self._tags = None
        self._presets = None
        self._queue = None
        self._cache = None
        self._termination = None
        self._plugins = None
        self._schedule = None
        self._events = None
        self._hooks = None
        self._dependencies = None
        self._trigger = None
        self._conditions = None
        self._skip_on_upstream_skip = None
        self._matrix = None
        self._joins = None
        self._params = None
        self._run_patch = None
        self._patch_strategy = None
        self._is_preset = None
        self._is_approved = None
        self._template = None
        self._url_ref = None
        self._path_ref = None
        self._hub_ref = None
        self._dag_ref = None
        self._component = None
        self.discriminator = None

        if version is not None:
            self.version = version
        if kind is not None:
            self.kind = kind
        if name is not None:
            self.name = name
        if description is not None:
            self.description = description
        if tags is not None:
            self.tags = tags
        if presets is not None:
            self.presets = presets
        if queue is not None:
            self.queue = queue
        if cache is not None:
            self.cache = cache
        if termination is not None:
            self.termination = termination
        if plugins is not None:
            self.plugins = plugins
        if schedule is not None:
            self.schedule = schedule
        if events is not None:
            self.events = events
        if hooks is not None:
            self.hooks = hooks
        if dependencies is not None:
            self.dependencies = dependencies
        if trigger is not None:
            self.trigger = trigger
        if conditions is not None:
            self.conditions = conditions
        if skip_on_upstream_skip is not None:
            self.skip_on_upstream_skip = skip_on_upstream_skip
        if matrix is not None:
            self.matrix = matrix
        if joins is not None:
            self.joins = joins
        if params is not None:
            self.params = params
        if run_patch is not None:
            self.run_patch = run_patch
        if patch_strategy is not None:
            self.patch_strategy = patch_strategy
        if is_preset is not None:
            self.is_preset = is_preset
        if is_approved is not None:
            self.is_approved = is_approved
        if template is not None:
            self.template = template
        if url_ref is not None:
            self.url_ref = url_ref
        if path_ref is not None:
            self.path_ref = path_ref
        if hub_ref is not None:
            self.hub_ref = hub_ref
        if dag_ref is not None:
            self.dag_ref = dag_ref
        if component is not None:
            self.component = component

    @property
    def version(self):
        """Gets the version of this V1Operation.  # noqa: E501


        :return: The version of this V1Operation.  # noqa: E501
        :rtype: float
        """
        return self._version

    @version.setter
    def version(self, version):
        """Sets the version of this V1Operation.


        :param version: The version of this V1Operation.  # noqa: E501
        :type: float
        """

        self._version = version

    @property
    def kind(self):
        """Gets the kind of this V1Operation.  # noqa: E501


        :return: The kind of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this V1Operation.


        :param kind: The kind of this V1Operation.  # noqa: E501
        :type: str
        """

        self._kind = kind

    @property
    def name(self):
        """Gets the name of this V1Operation.  # noqa: E501


        :return: The name of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this V1Operation.


        :param name: The name of this V1Operation.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def description(self):
        """Gets the description of this V1Operation.  # noqa: E501


        :return: The description of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._description

    @description.setter
    def description(self, description):
        """Sets the description of this V1Operation.


        :param description: The description of this V1Operation.  # noqa: E501
        :type: str
        """

        self._description = description

    @property
    def tags(self):
        """Gets the tags of this V1Operation.  # noqa: E501


        :return: The tags of this V1Operation.  # noqa: E501
        :rtype: list[str]
        """
        return self._tags

    @tags.setter
    def tags(self, tags):
        """Sets the tags of this V1Operation.


        :param tags: The tags of this V1Operation.  # noqa: E501
        :type: list[str]
        """

        self._tags = tags

    @property
    def presets(self):
        """Gets the presets of this V1Operation.  # noqa: E501


        :return: The presets of this V1Operation.  # noqa: E501
        :rtype: list[str]
        """
        return self._presets

    @presets.setter
    def presets(self, presets):
        """Sets the presets of this V1Operation.


        :param presets: The presets of this V1Operation.  # noqa: E501
        :type: list[str]
        """

        self._presets = presets

    @property
    def queue(self):
        """Gets the queue of this V1Operation.  # noqa: E501


        :return: The queue of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._queue

    @queue.setter
    def queue(self, queue):
        """Sets the queue of this V1Operation.


        :param queue: The queue of this V1Operation.  # noqa: E501
        :type: str
        """

        self._queue = queue

    @property
    def cache(self):
        """Gets the cache of this V1Operation.  # noqa: E501


        :return: The cache of this V1Operation.  # noqa: E501
        :rtype: V1Cache
        """
        return self._cache

    @cache.setter
    def cache(self, cache):
        """Sets the cache of this V1Operation.


        :param cache: The cache of this V1Operation.  # noqa: E501
        :type: V1Cache
        """

        self._cache = cache

    @property
    def termination(self):
        """Gets the termination of this V1Operation.  # noqa: E501


        :return: The termination of this V1Operation.  # noqa: E501
        :rtype: V1Termination
        """
        return self._termination

    @termination.setter
    def termination(self, termination):
        """Sets the termination of this V1Operation.


        :param termination: The termination of this V1Operation.  # noqa: E501
        :type: V1Termination
        """

        self._termination = termination

    @property
    def plugins(self):
        """Gets the plugins of this V1Operation.  # noqa: E501


        :return: The plugins of this V1Operation.  # noqa: E501
        :rtype: V1Plugins
        """
        return self._plugins

    @plugins.setter
    def plugins(self, plugins):
        """Sets the plugins of this V1Operation.


        :param plugins: The plugins of this V1Operation.  # noqa: E501
        :type: V1Plugins
        """

        self._plugins = plugins

    @property
    def schedule(self):
        """Gets the schedule of this V1Operation.  # noqa: E501


        :return: The schedule of this V1Operation.  # noqa: E501
        :rtype: object
        """
        return self._schedule

    @schedule.setter
    def schedule(self, schedule):
        """Sets the schedule of this V1Operation.


        :param schedule: The schedule of this V1Operation.  # noqa: E501
        :type: object
        """

        self._schedule = schedule

    @property
    def events(self):
        """Gets the events of this V1Operation.  # noqa: E501


        :return: The events of this V1Operation.  # noqa: E501
        :rtype: list[V1EventTrigger]
        """
        return self._events

    @events.setter
    def events(self, events):
        """Sets the events of this V1Operation.


        :param events: The events of this V1Operation.  # noqa: E501
        :type: list[V1EventTrigger]
        """

        self._events = events

    @property
    def hooks(self):
        """Gets the hooks of this V1Operation.  # noqa: E501


        :return: The hooks of this V1Operation.  # noqa: E501
        :rtype: list[V1Hook]
        """
        return self._hooks

    @hooks.setter
    def hooks(self, hooks):
        """Sets the hooks of this V1Operation.


        :param hooks: The hooks of this V1Operation.  # noqa: E501
        :type: list[V1Hook]
        """

        self._hooks = hooks

    @property
    def dependencies(self):
        """Gets the dependencies of this V1Operation.  # noqa: E501


        :return: The dependencies of this V1Operation.  # noqa: E501
        :rtype: list[str]
        """
        return self._dependencies

    @dependencies.setter
    def dependencies(self, dependencies):
        """Sets the dependencies of this V1Operation.


        :param dependencies: The dependencies of this V1Operation.  # noqa: E501
        :type: list[str]
        """

        self._dependencies = dependencies

    @property
    def trigger(self):
        """Gets the trigger of this V1Operation.  # noqa: E501


        :return: The trigger of this V1Operation.  # noqa: E501
        :rtype: V1TriggerPolicy
        """
        return self._trigger

    @trigger.setter
    def trigger(self, trigger):
        """Sets the trigger of this V1Operation.


        :param trigger: The trigger of this V1Operation.  # noqa: E501
        :type: V1TriggerPolicy
        """

        self._trigger = trigger

    @property
    def conditions(self):
        """Gets the conditions of this V1Operation.  # noqa: E501


        :return: The conditions of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._conditions

    @conditions.setter
    def conditions(self, conditions):
        """Sets the conditions of this V1Operation.


        :param conditions: The conditions of this V1Operation.  # noqa: E501
        :type: str
        """

        self._conditions = conditions

    @property
    def skip_on_upstream_skip(self):
        """Gets the skip_on_upstream_skip of this V1Operation.  # noqa: E501


        :return: The skip_on_upstream_skip of this V1Operation.  # noqa: E501
        :rtype: bool
        """
        return self._skip_on_upstream_skip

    @skip_on_upstream_skip.setter
    def skip_on_upstream_skip(self, skip_on_upstream_skip):
        """Sets the skip_on_upstream_skip of this V1Operation.


        :param skip_on_upstream_skip: The skip_on_upstream_skip of this V1Operation.  # noqa: E501
        :type: bool
        """

        self._skip_on_upstream_skip = skip_on_upstream_skip

    @property
    def matrix(self):
        """Gets the matrix of this V1Operation.  # noqa: E501


        :return: The matrix of this V1Operation.  # noqa: E501
        :rtype: object
        """
        return self._matrix

    @matrix.setter
    def matrix(self, matrix):
        """Sets the matrix of this V1Operation.


        :param matrix: The matrix of this V1Operation.  # noqa: E501
        :type: object
        """

        self._matrix = matrix

    @property
    def joins(self):
        """Gets the joins of this V1Operation.  # noqa: E501


        :return: The joins of this V1Operation.  # noqa: E501
        :rtype: dict(str, V1Join)
        """
        return self._joins

    @joins.setter
    def joins(self, joins):
        """Sets the joins of this V1Operation.


        :param joins: The joins of this V1Operation.  # noqa: E501
        :type: dict(str, V1Join)
        """

        self._joins = joins

    @property
    def params(self):
        """Gets the params of this V1Operation.  # noqa: E501


        :return: The params of this V1Operation.  # noqa: E501
        :rtype: dict(str, V1Param)
        """
        return self._params

    @params.setter
    def params(self, params):
        """Sets the params of this V1Operation.


        :param params: The params of this V1Operation.  # noqa: E501
        :type: dict(str, V1Param)
        """

        self._params = params

    @property
    def run_patch(self):
        """Gets the run_patch of this V1Operation.  # noqa: E501


        :return: The run_patch of this V1Operation.  # noqa: E501
        :rtype: object
        """
        return self._run_patch

    @run_patch.setter
    def run_patch(self, run_patch):
        """Sets the run_patch of this V1Operation.


        :param run_patch: The run_patch of this V1Operation.  # noqa: E501
        :type: object
        """

        self._run_patch = run_patch

    @property
    def patch_strategy(self):
        """Gets the patch_strategy of this V1Operation.  # noqa: E501


        :return: The patch_strategy of this V1Operation.  # noqa: E501
        :rtype: V1PatchStrategy
        """
        return self._patch_strategy

    @patch_strategy.setter
    def patch_strategy(self, patch_strategy):
        """Sets the patch_strategy of this V1Operation.


        :param patch_strategy: The patch_strategy of this V1Operation.  # noqa: E501
        :type: V1PatchStrategy
        """

        self._patch_strategy = patch_strategy

    @property
    def is_preset(self):
        """Gets the is_preset of this V1Operation.  # noqa: E501


        :return: The is_preset of this V1Operation.  # noqa: E501
        :rtype: bool
        """
        return self._is_preset

    @is_preset.setter
    def is_preset(self, is_preset):
        """Sets the is_preset of this V1Operation.


        :param is_preset: The is_preset of this V1Operation.  # noqa: E501
        :type: bool
        """

        self._is_preset = is_preset

    @property
    def is_approved(self):
        """Gets the is_approved of this V1Operation.  # noqa: E501


        :return: The is_approved of this V1Operation.  # noqa: E501
        :rtype: bool
        """
        return self._is_approved

    @is_approved.setter
    def is_approved(self, is_approved):
        """Sets the is_approved of this V1Operation.


        :param is_approved: The is_approved of this V1Operation.  # noqa: E501
        :type: bool
        """

        self._is_approved = is_approved

    @property
    def template(self):
        """Gets the template of this V1Operation.  # noqa: E501


        :return: The template of this V1Operation.  # noqa: E501
        :rtype: V1Template
        """
        return self._template

    @template.setter
    def template(self, template):
        """Sets the template of this V1Operation.


        :param template: The template of this V1Operation.  # noqa: E501
        :type: V1Template
        """

        self._template = template

    @property
    def url_ref(self):
        """Gets the url_ref of this V1Operation.  # noqa: E501


        :return: The url_ref of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._url_ref

    @url_ref.setter
    def url_ref(self, url_ref):
        """Sets the url_ref of this V1Operation.


        :param url_ref: The url_ref of this V1Operation.  # noqa: E501
        :type: str
        """

        self._url_ref = url_ref

    @property
    def path_ref(self):
        """Gets the path_ref of this V1Operation.  # noqa: E501


        :return: The path_ref of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._path_ref

    @path_ref.setter
    def path_ref(self, path_ref):
        """Sets the path_ref of this V1Operation.


        :param path_ref: The path_ref of this V1Operation.  # noqa: E501
        :type: str
        """

        self._path_ref = path_ref

    @property
    def hub_ref(self):
        """Gets the hub_ref of this V1Operation.  # noqa: E501


        :return: The hub_ref of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._hub_ref

    @hub_ref.setter
    def hub_ref(self, hub_ref):
        """Sets the hub_ref of this V1Operation.


        :param hub_ref: The hub_ref of this V1Operation.  # noqa: E501
        :type: str
        """

        self._hub_ref = hub_ref

    @property
    def dag_ref(self):
        """Gets the dag_ref of this V1Operation.  # noqa: E501


        :return: The dag_ref of this V1Operation.  # noqa: E501
        :rtype: str
        """
        return self._dag_ref

    @dag_ref.setter
    def dag_ref(self, dag_ref):
        """Sets the dag_ref of this V1Operation.


        :param dag_ref: The dag_ref of this V1Operation.  # noqa: E501
        :type: str
        """

        self._dag_ref = dag_ref

    @property
    def component(self):
        """Gets the component of this V1Operation.  # noqa: E501


        :return: The component of this V1Operation.  # noqa: E501
        :rtype: V1Component
        """
        return self._component

    @component.setter
    def component(self, component):
        """Sets the component of this V1Operation.


        :param component: The component of this V1Operation.  # noqa: E501
        :type: V1Component
        """

        self._component = component

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
        if not isinstance(other, V1Operation):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1Operation):
            return True

        return self.to_dict() != other.to_dict()
