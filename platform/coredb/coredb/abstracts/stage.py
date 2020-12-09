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

from django.core.serializers.json import DjangoJSONEncoder
from django.db import models

from polyaxon.lifecycle import V1Stages


class StageModel(models.Model):
    stage = models.CharField(
        max_length=16,
        blank=True,
        null=True,
        db_index=True,
        default=V1Stages.TESTING,
        choices=V1Stages.CHOICES,
    )
    stage_conditions = models.JSONField(
        encoder=DjangoJSONEncoder, blank=True, null=True, default=dict
    )

    class Meta:
        abstract = True
