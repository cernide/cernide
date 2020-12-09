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

from typing import Dict, Optional, Set, Tuple, Union

from coredb.abstracts.getter import get_run_model
from coredb.abstracts.runs import BaseRun
from polyaxon.lifecycle import V1StatusCondition, V1Statuses
from polyaxon.metadata.keys import META_HAS_DAGS, META_HAS_JOBS, META_HAS_SERVICES
from polyaxon.polyaxonfile import OperationSpecification
from polyaxon.polyflow import (
    V1CloningKind,
    V1CompiledOperation,
    V1Init,
    V1MatrixKind,
    V1Operation,
    V1RunKind,
)
from polyaxon.schemas.types import V1ArtifactsType
from polycommon.service_interface import Service


class OperationsService(Service):
    DEFAULT_KINDS = V1RunKind.default_runtime_values
    __all__ = ("init_run",)

    @staticmethod
    def set_spec(spec: V1Operation, **kwargs) -> Tuple[V1Operation, Dict]:
        kwargs["raw_content"] = spec.to_dict(dump=True)
        return spec, kwargs

    @staticmethod
    def get_kind(compiled_operation: V1CompiledOperation) -> Tuple[str, Optional[str]]:
        if compiled_operation.is_tf_job_run:
            return V1RunKind.JOB, V1RunKind.TFJOB
        elif compiled_operation.is_pytorch_job_run:
            return V1RunKind.JOB, V1RunKind.PYTORCHJOB
        elif compiled_operation.is_mpi_job_run:
            return V1RunKind.JOB, V1RunKind.MPIJOB
        elif compiled_operation.is_dask_run:
            return V1RunKind.JOB, V1RunKind.DASK
        elif compiled_operation.is_spark_run:
            return V1RunKind.JOB, V1RunKind.SPARK
        # Default case
        kind = compiled_operation.run.kind
        return kind, kind

    @classmethod
    def supports_kind(
        cls, kind: str, runtime: str, supported_kinds: Set[str], is_managed: bool
    ) -> bool:
        supported_kinds = supported_kinds or cls.DEFAULT_KINDS
        error_message = (
            "You cannot create this operation, "
            "your account has reached the allowed quota or does not support operations of kind: {}"
        )
        if kind not in supported_kinds:
            if is_managed or kind not in V1RunKind.eager_values:
                raise ValueError(error_message.format(kind))
        if runtime and runtime not in supported_kinds:
            if is_managed or runtime not in V1MatrixKind.eager_values:
                raise ValueError(error_message.format(runtime))
        return True

    @classmethod
    def _finalize_meta_info(cls, meta_info: Dict, **kwargs):
        return meta_info

    @classmethod
    def get_meta_info(
        cls,
        compiled_operation: V1CompiledOperation,
        kind: str,
        runtime: str,
        meta_info: Dict = None,
        **kwargs,
    ) -> Tuple[str, str, Dict]:
        meta_info = meta_info or {}
        if compiled_operation.matrix or compiled_operation.schedule:
            if kind == V1RunKind.JOB:
                meta_info[META_HAS_JOBS] = True
            elif kind == V1RunKind.SERVICE:
                meta_info[META_HAS_SERVICES] = True
            elif kind == V1RunKind.DAG:
                meta_info[META_HAS_DAGS] = True
        if compiled_operation.schedule:
            kind = V1RunKind.SCHEDULE
            runtime = compiled_operation.schedule.kind
        elif compiled_operation.matrix:
            kind = V1RunKind.MATRIX
            runtime = compiled_operation.matrix.kind

        meta_info = cls._finalize_meta_info(meta_info=meta_info, **kwargs)
        return kind, runtime, meta_info

    @staticmethod
    def sanitize_kwargs(**kwargs):
        results = {}
        if kwargs.get("raw_content"):
            results["raw_content"] = kwargs["raw_content"]
        if kwargs.get("content"):
            results["content"] = kwargs["content"]

        return results

    def init_run(
        self,
        project_id: int,
        user_id: int,
        op_spec: V1Operation = None,
        compiled_operation: V1CompiledOperation = None,
        name: str = None,
        description: str = None,
        tags: str = None,
        override: Union[str, Dict] = None,
        params: Dict = None,
        readme: str = None,
        original_id: int = None,
        original_uuid: int = None,
        cloning_kind: str = None,
        is_managed: bool = True,
        is_approved: bool = True,
        meta_info: Dict = None,
        supported_kinds: Set[str] = None,
        **kwargs,
    ) -> Tuple[V1CompiledOperation, BaseRun]:
        if op_spec:
            op_spec, kwargs = self.set_spec(op_spec, **kwargs)
        if op_spec:
            if not compiled_operation or override:
                compiled_operation = OperationSpecification.compile_operation(
                    op_spec, override=override
                )
            params = op_spec.params

        params = params or {}
        inputs = {p: pv.value for p, pv in params.items() if pv.is_literal}
        params = {p: pv.to_dict() for p, pv in params.items()}
        kind = None
        meta_info = meta_info or {}
        if compiled_operation:
            if is_approved and compiled_operation.is_approved is not None:
                is_approved = compiled_operation.is_approved
            name = name or compiled_operation.name
            description = description or compiled_operation.description
            tags = tags or compiled_operation.tags
            kind, runtime = self.get_kind(compiled_operation)
            kind, runtime, meta_info = self.get_meta_info(
                compiled_operation, kind, runtime, meta_info, **kwargs
            )
            self.supports_kind(kind, runtime, supported_kinds, is_managed)
            if cloning_kind == V1CloningKind.COPY:
                if runtime not in {V1RunKind.JOB, V1RunKind.SERVICE}:
                    raise ValueError(
                        "Operation with kind `{}` does not support restart with copy mode."
                    )
                compiled_operation.run.add_init(
                    V1Init(artifacts=V1ArtifactsType(dirs=[original_uuid]))
                )
            kwargs["content"] = compiled_operation.to_dict(dump=True)
        instance = get_run_model()(
            project_id=project_id,
            user_id=user_id,
            name=name,
            description=description,
            tags=tags,
            readme=readme,
            params=params,
            inputs=inputs,
            kind=kind,
            runtime=runtime,
            meta_info=meta_info,
            original_id=original_id,
            cloning_kind=cloning_kind,
            is_managed=is_managed,
            is_approved=is_approved,
            status_conditions=[
                V1StatusCondition.get_condition(
                    type=V1Statuses.CREATED,
                    status="True",
                    reason=kwargs.pop("reason", "PolyaxonRunCreated"),
                    message=kwargs.pop("message", "Run is created"),
                ).to_dict()
            ],
            **self.sanitize_kwargs(**kwargs),
        )
        return compiled_operation, instance
