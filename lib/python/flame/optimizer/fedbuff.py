# Copyright 2023 Cisco Systems, Inc. and its affiliates
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
#
# SPDX-License-Identifier: Apache-2.0
"""FedBuff optimizer.

The implementation is based on the following paper:
https://arxiv.org/pdf/2106.06639.pdf
https://arxiv.org/pdf/2111.04877.pdf

SecAgg algorithm is not the scope of this implementation.
"""
import logging
import math
from typing import Union

from diskcache import Cache

from ..common.util import (MLFramework, get_ml_framework_in_use,
                           valid_frameworks)
from .abstract import AbstractOptimizer

logger = logging.getLogger(__name__)


class FedBuff(AbstractOptimizer):
    """FedBuff class."""

    def __init__(self):
        """Initialize FedBuff instance."""
        self.agg_weights = None

        ml_framework_in_use = get_ml_framework_in_use()
        if ml_framework_in_use == MLFramework.PYTORCH:
            self.aggregate_fn = self._aggregate_pytorch
        elif ml_framework_in_use == MLFramework.TENSORFLOW:
            self.aggregate_fn = self._aggregate_tesnorflow
        else:
            raise NotImplementedError(
                "supported ml framework not found; "
                f"supported frameworks are: {valid_frameworks}")

    def do(self,
           cache: Cache,
           *,
           base_weights=None,
           total: int = 0,
           version: int = 0) -> Union[list, dict]:
        """Do aggregates models of trainers.

        Return: aggregated model
        """
        logger.debug("calling fedbuff")

        # reset global weights before aggregation
        self.agg_weights = base_weights

        if len(cache) == 0 or total == 0:
            return None

        for k in list(cache.iterkeys()):
            # after popping, the item is removed from the cache
            # hence, explicit cache cleanup is not needed
            tres = cache.pop(k)

            logger.debug(f"agg ver: {version}, trainer ver: {tres.version}")
            # rate determined based on the staleness of local model
            rate = 1 / math.sqrt(1 + version - tres.version)
            self.aggregate_fn(tres, rate)

        return self.agg_weights

    def _aggregate_pytorch(self, tres, rate):
        logger.debug("calling _aggregate_pytorch")

        if self.agg_weights is None:
            self.agg_weights = {k: v * rate for k, v in tres.weights.items()}
        else:
            for k, v in tres.weights.items():
                self.agg_weights[k] += v * rate

    def _aggregate_tesnorflow(self, tres, rate):
        logger.debug("calling _aggregate_tensorflow")

        if self.agg_weights is None:
            self.agg_weights = [weight * rate for weight in tres.weights]
        else:
            for idx in range(len(tres.weights)):
                self.agg_weights[idx] += tres.weights[idx] * rate