# Django REST framework

Django REST framework is a powerful and flexible toolkit for building Web APIs

Some reasons you might want to use REST framework:

- The Web browable API is a huge usability win for your developers.
- Authentication policies including packages for OAuth1a and OAuth2.
- Serilization that supports both ORM and non-ORM data sources.
- Customizable all the way down - just use regular function-based views if you don't need the more powerful features.

**Installation**

```py
pip install djangorestframework
pip install markdown
pip install django-filter
```

**Serializers**

**Views**

**Urls**

Okay, now let's wire up the API URLs.

**Pagination**

**Settings**

**Testing our API**

## Serilization

Introduction

**Creating a Serializer class**

The first thing we need to get started on our Web API is to provide a way of serializing and deserializing the snippet instances into representation such as json. We can do this by declaring serializers that work very similar to Django's forms. Create a file in the snippets directory named serializers.py.

**Using ModelSerializers**

Our SnippetSerializer class is replicating a lot of information that's also contained in the Snippet model. It would be nice if we could keep our code a bit more concise.

Once nice property that serializers have it that you can inspect all the fields in a serializer instance, by printing its representation.

