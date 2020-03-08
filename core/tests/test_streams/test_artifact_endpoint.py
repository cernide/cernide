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
import os

from tests.test_streams.base import create_tmp_files, get_streams_client, set_store
from tests.utils import BaseTestCase

from polyaxon import settings
from polyaxon.streams.app.main import STREAMS_URL
from polyaxon.utils.path_utils import create_path


class TestArtifactEndpoints(BaseTestCase):
    def setUp(self):
        super().setUp()
        self.store_root = set_store()
        self.run_path = os.path.join(self.store_root, "uuid")
        # Create run artifacts path and some files
        create_path(self.run_path)
        create_tmp_files(self.run_path)
        # Archive path
        self.archive_run_path = os.path.join(
            settings.CLIENT_CONFIG.archive_root, "uuid"
        )

        self.client = get_streams_client()
        self.base_url = STREAMS_URL + "/namespace/owner/project/runs/uuid/artifact"

    def test_download_artifact_not_passing_path(self):
        response = self.client.get(self.base_url)
        assert response.status_code == 400

    def test_stream_artifact_not_passing_path(self):
        response = self.client.get(self.base_url + "?=stream=true")
        assert response.status_code == 400

    def test_download_artifact_non_existing_path(self):
        response = self.client.get(self.base_url + "?path=foo")
        assert response.status_code == 404

    def test_stream_artifact_non_existing_path(self):
        response = self.client.get(self.base_url + "?stream=true&path=foo")
        assert response.status_code == 404

    def test_download_artifact_passing_path(self):
        filepath = os.path.join(self.archive_run_path, "1")
        assert os.path.exists(filepath) is False
        response = self.client.get(self.base_url + "?path=1")
        assert response.status_code == 200
        assert os.path.exists(filepath) is True
        assert response.headers["Content-Type"] == ""
        assert response.headers["X-Accel-Redirect"] == os.path.join(
            self.archive_run_path, "1"
        )
        assert response.headers[
            "Content-Disposition"
        ] == 'attachment; filename="{}"'.format("1")

        # Nested dirs
        nested_path = os.path.join(self.run_path, "foo")
        create_path(nested_path)
        create_tmp_files(nested_path)
        filepath = os.path.join(self.archive_run_path, "foo", "1")
        assert os.path.exists(filepath) is False
        response = self.client.get(self.base_url + "?path=foo/1")
        assert response.status_code == 200
        assert os.path.exists(filepath) is True
        assert response.headers["Content-Type"] == ""
        assert response.headers["X-Accel-Redirect"] == os.path.join(
            self.archive_run_path, "foo/1"
        )
        assert response.headers[
            "Content-Disposition"
        ] == 'attachment; filename="{}"'.format("1")

    def test_stream_artifact_passing_path(self):
        filepath = os.path.join(self.archive_run_path, "1")
        assert os.path.exists(filepath) is False
        response = self.client.get(self.base_url + "?stream=true&path=1")
        assert response.status_code == 200
        assert os.path.exists(filepath) is True
        assert response.headers["Content-Type"] == "text/plain; charset=utf-8"
        assert response.headers["content-length"] == "0"
        assert response.headers["last-modified"] is not None
        assert response.headers["etag"] is not None

        # Nested dirs
        nested_path = os.path.join(self.run_path, "foo")
        create_path(nested_path)
        create_tmp_files(nested_path)
        filepath = os.path.join(self.archive_run_path, "foo", "1")
        assert os.path.exists(filepath) is False
        response = self.client.get(self.base_url + "?stream=true&path=foo/1")
        assert response.status_code == 200
        assert os.path.exists(filepath) is True
        assert response.headers["Content-Type"] == "text/plain; charset=utf-8"
        assert response.headers["content-length"] == "0"
        assert response.headers["last-modified"] is not None
        assert response.headers["etag"] is not None
