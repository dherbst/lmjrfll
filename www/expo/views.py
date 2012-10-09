

from django.http import HttpResponse
from django.shortcuts import render_to_response, get_object_or_404
from django.template import loader, Context, RequestContext


def expo_home(request):
    
    return HttpResponse('expo home')

