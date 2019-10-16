# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

from polyaxon.cli.getters.project import get_project_or_local
from polyaxon.managers.run import RunManager


def get_run_or_local(experiment=None):
    return experiment or RunManager.get_config_or_raise().id


def get_project_run_or_local(project=None, experiment=None):
    user, project_name = get_project_or_local(project)
    experiment = get_run_or_local(experiment)
    return user, project_name, experiment
