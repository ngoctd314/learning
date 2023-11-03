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

## Requests and Responses

**Request Objects**

REST framework introduces a Request object that extends the regular HttpRequest, and provides more flexible request passing. The core functionality of the Request object is the request.data attribute, which is similar to request.POST, but more useful for working with web APIS.

```txt
request.POST # Only handles form data. Only works for 'POST' method.
request.data # handles arbitrary data. Works for 'POST', 'PUT', and 'PATCH' methods.
```

**Response objects**

REST framework also introduces a Response object, which is a type of TemplateResponse that takes unrendered content and uses content negotiation to determine the correct type to return to the client.  

**Status codes**

Using numeric HTTP status codes in your views doesn't always make for obvious reading, and it's easy to not notice if you get an error code wrong. REST framework provides more explicit identifiers for each status code, such as HTTP_400_BAD_REQUEST in the status module.

**Wrapping API views**

Decorator @api_view

request.data can handle incoming json requests, but it can also handle other formats. Similar we're returning response objects with data, but allowing REST fw to render the response into the correct content type for us. 

## Class-based Views

We can also write our API views using class-based views, rather than function based views. As we'll see this is a powerful pattern that allows us to reuse common functionality, and helps us keep our code DRY.

We're building our view using GenericAPIView, and adding in ListModelMixin and CreateModelMixin

The base class provides the core functionality, and the mixin classes provide the .list() and .create() actions. We're then explicit binding the get and post methods to the appropriate actions.

## Authentication & Permissions

Currently our API doesn't have any restrictions on who can edit or delete code snippets. We'd like to have some more advanced behavior in order to make sure that: 

**Adding informatino to our model**

We're going to make a couple of changes to our Snippet model class. First, let's add a couple of fields. One of those fields will be used to represent the user
