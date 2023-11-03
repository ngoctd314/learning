from rest_framework import serializers
from .models import Snippet, LANGUAGE_CHOICES, STYLE_CHOICES


class SnippetSerializer(serializers.ModelSerializer):
    class Meta:
        model = Snippet
        fields = ["id", "title", "code", "linenos", "language", "style"]

    def create(self, validated_data):
        return Snippet.objects.create(**validated_data)

    def update(self, instance, validated_data):
        instance.title = validated_data.get("title", instance.title)
        instance.code = validated_data.get("code", instance.code)
        instance.lineons = validated_data.get("lineons", instance.lineons)
        instance.language = validated_data.get("language", instance.language)
        instance.style = validated_data.get("style", instance.style)
        instance.save()

        return instance
