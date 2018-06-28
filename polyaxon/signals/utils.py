from polyaxon_schemas.environments import PersistenceConfig

from libs.paths.data_paths import validate_persistence_data
from libs.paths.outputs_paths import validate_persistence_outputs


def set_tags(instance):
    if not instance.tags and instance.specification:
        instance.tags = instance.specification.tags


def set_persistence(instance):
    if instance.persistence:
        return

    persistence_data = None
    persistence_outputs = None
    if instance.specification and instance.specification.persistence:
        persistence_data = instance.specification.persistence.data
        persistence_outputs = instance.specification.persistence.outputs

    persistence_data = validate_persistence_data(persistence_data=persistence_data)
    persistence_outputs = validate_persistence_outputs(persistence_outputs=persistence_outputs)
    persistence_config = PersistenceConfig(data=persistence_data, outputs=persistence_outputs)
    instance.persistence = persistence_config.to_dict()
