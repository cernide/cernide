from rest_framework import serializers

from db.models.activitylogs import ActivityLog
from event_manager import event_context


class ActivityLogsSerializer(serializers.ModelSerializer):
    object_name = serializers.SerializerMethodField()
    event_action = serializers.SerializerMethodField()
    event_subject = serializers.SerializerMethodField()
    actor = serializers.SerializerMethodField()

    class Meta:
        model = ActivityLog
        fields = [
            'id',
            'event_action',
            'event_subject',
            'actor',
            'created_at',
            'object_id',
            'object_name'
        ]

    def get_actor(self, obj):
        return obj.actor.username

    def get_event_action(self, obj):
        return event_context.get_event_action(event_type=obj.event_type)

    def get_event_subject(self, obj):
        return event_context.get_event_subject(event_type=obj.event_type)

    def get_object_name(self, obj):
        return event_context.get_event_object_context(
            event_content_object=obj.content_object,
            event_type=obj.event_type
        ).name
