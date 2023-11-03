from snippets import views
from django.urls import path


urlpatterns = [
    path("snippets/", views.SnippetList.as_view()),
    path("snippets/<int:pk>", views.SnippetDetail.as_view()),
]
