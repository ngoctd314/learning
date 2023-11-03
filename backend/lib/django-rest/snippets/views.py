from django.http import Http404
from rest_framework.response import Response
from rest_framework.request import Request
from rest_framework.views import APIView
from .models import Snippet

from .serializers import SnippetSerializer
from rest_framework.decorators import api_view
from rest_framework import generics, mixins, status

from django.views.decorators.csrf import csrf_exempt


# class-based view
class SnippetList(
    mixins.ListModelMixin, mixins.CreateModelMixin, generics.GenericAPIView
):
    """
    List app snippets, or create a new snippet
    """

    queryset = Snippet.objects.all()
    serializer_class = SnippetSerializer

    def get(self, request, *args, **kwargs):
        return self.list(request, *args, **kwargs)

    def post(self, request, *args, **kwargs):
        return self.create(request, *args, **kwargs)


class SnippetDetail(
    mixins.RetrieveModelMixin,
    mixins.UpdateModelMixin,
    mixins.DestroyModelMixin,
    generics.GenericAPIView,
):
    """
    Retrieve, update or delete a snippet instance
    """

    # queryset, serializer_class is extend from generics.GenericAPIView
    queryset = Snippet.objects.all()
    serializer_class = SnippetSerializer

    def get(self, request, *args, **kwargs):
        # retrieve is extend from mixins.RetrieveModelMixin
        return self.retrieve(request, *args, **kwargs)

    def put(self, request, *args, **kwargs):
        # update is extend from mixins.UpdateModelMixin
        return self.update(request, *args, **kwargs)

    def delete(self, request, *args, **kwargs):
        # destroy is extend from mixins.DestroyModelMixin
        return self.destroy(request, *args, **kwargs)
