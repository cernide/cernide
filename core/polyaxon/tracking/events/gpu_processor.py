# !/usr/bin/python
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
from typing import Dict, List

from polyaxon import logger
from polyaxon.tracking.events.events_processors import metrics_dict_to_list

try:
    import psutil
except ImportError:
    psutil = None

try:
    import pynvml
except ImportError:
    pynvml = None


def can_log_gpu_resources():
    if pynvml is None:
        return False

    try:
        pynvml.nvmlInit()
        return True
    except pynvml.NVMLError:
        return False


def query_gpu(handle: int) -> Dict:
    memory = pynvml.nvmlDeviceGetMemoryInfo(handle)  # in Bytes
    utilization = pynvml.nvmlDeviceGetUtilizationRates(handle)

    return {
        "gpu_{}_memory_free".format(handle): int(memory.free),
        "gpu_{}_memory_used".format(handle): int(memory.used),
        "gpu_{}_utilization".format(handle): utilization.gpu,
    }


def get_gpu_metrics() -> List:
    try:
        pynvml.nvmlInit()
        device_count = pynvml.nvmlDeviceGetCount()
        results = []

        for i in range(device_count):
            handle = pynvml.nvmlDeviceGetHandleByIndex(i)
            results += metrics_dict_to_list(query_gpu(handle))
        return results
    except pynvml.NVMLError:
        logger.debug("Failed to collect gpu resources", exc_info=True)
        return []
